package repositories

import (
	"database/sql"
	"task_api/src/models"
)

type TaskRepository struct{}

func (t *TaskRepository) FindAll() ([]models.Task, error) {
	db := createConnection()
	tasks := make([]models.Task, 0)

	rows, err := db.Query("SELECT * FROM task")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var task models.Task

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

func (t *TaskRepository) FindById(id int64) (*models.Task, error) {
	db := createConnection()

	row := db.QueryRow("SELECT * FROM task WHERE id = ?", id)

	var task models.Task

	if err := row.Scan(&task.Id, &task.Name, &task.Finished); err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *TaskRepository) Save(task models.Task) (*models.Task, error) {
	db := createConnection()

	var res sql.Result
	var err error

	if task.Id != 0 {
		res, err = db.Exec("UPDATE task SET name = ?, finished = ? WHERE id = ?", task.Name, task.Finished, task.Id)
	} else {
		res, err = db.Exec("INSERT INTO task (name, finished) VALUES (?, ?)", task.Name, task.Finished)
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
	db := createConnection()

	_, err := db.Exec("DELETE FROM task WHERE id = ?", id)

	return err
}
