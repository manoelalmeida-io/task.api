package repository

import (
	"database/sql"
	"task_api/internal/model"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (t *TaskRepository) FindAll() ([]model.Task, error) {
	tasks := make([]model.Task, 0)

	rows, err := t.db.Query("SELECT * FROM task")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var task model.Task

		if err := rows.Scan(&task.Id, &task.Name, &task.Finished); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *TaskRepository) FindById(id int64) (*model.Task, error) {
	row := t.db.QueryRow("SELECT * FROM task WHERE id = ?", id)

	var task model.Task

	if err := row.Scan(&task.Id, &task.Name, &task.Finished); err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *TaskRepository) Save(task model.Task) (*model.Task, error) {
	var res sql.Result
	var err error

	if task.Id != 0 {
		res, err = t.db.Exec("UPDATE task SET name = ?, finished = ? WHERE id = ?", task.Name, task.Finished, task.Id)
	} else {
		res, err = t.db.Exec("INSERT INTO task (name, finished) VALUES (?, ?)", task.Name, task.Finished)
	}

	if err != nil {
		return nil, err
	}

	if task.Id != 0 {
		return t.FindById(task.Id)
	}

	lastInsertedId, err := res.LastInsertId()

	if err != nil {
		return nil, err
	}

	return t.FindById(lastInsertedId)
}

func (t *TaskRepository) DeleteById(id int64) error {
	_, err := t.db.Exec("DELETE FROM task WHERE id = ?", id)

	return err
}
