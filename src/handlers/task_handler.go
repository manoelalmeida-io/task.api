package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"task_api/src/models"
	"task_api/src/repositories"

	"github.com/labstack/echo/v4"
)

var taskRepository repositories.TaskRepository

func GetTasksHandler(c echo.Context) error {
	tasks, err := taskRepository.FindAll()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, tasks)
}

func GetTaskByIdHandler(c echo.Context) error {
	strIndex := c.Param("id")

	id, err := strconv.Atoi(strIndex)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	task, err := taskRepository.FindById(int64(id))

	if err == sql.ErrNoRows {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, task)
}

func CreateTaskHandler(c echo.Context) error {
	task := new(models.Task)

	if err := c.Bind(task); err != nil {
		return err
	}

	task, err := taskRepository.Save(*task)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, task)
}

func UpdateTaskHandler(c echo.Context) error {
	strIndex := c.Param("id")

	id, err := strconv.Atoi(strIndex)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	task, err := taskRepository.FindById(int64(id))

	if err == sql.ErrNoRows {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	requestTask := new(models.Task)

	if err := c.Bind(requestTask); err != nil {
		return err
	}

	task.Name = requestTask.Name

	updatedTask, err := taskRepository.Save(*task)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updatedTask)
}

func ToggleTaskHandler(c echo.Context) error {
	strIndex := c.Param("id")

	id, err := strconv.Atoi(strIndex)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	task, err := taskRepository.FindById(int64(id))

	if err == sql.ErrNoRows {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	task.Finished = !task.Finished

	updatedTask, err := taskRepository.Save(*task)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updatedTask)
}

func DeleteTaskByIdHandler(c echo.Context) error {
	strIndex := c.Param("id")

	id, err := strconv.Atoi(strIndex)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Parameter id cannot be converted to integer")
	}

	_, err = taskRepository.FindById(int64(id))

	if err == sql.ErrNoRows {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err = taskRepository.DeleteById(int64(id)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, nil)
}
