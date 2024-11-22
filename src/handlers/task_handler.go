package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"task_api/src/models"
	"task_api/src/repositories"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	taskRepository *repositories.TaskRepository
}

func NewTaskHandler(taskRepository *repositories.TaskRepository) *TaskHandler {
	return &TaskHandler{taskRepository: taskRepository}
}

func (h *TaskHandler) GetTasksHandler(c echo.Context) error {
	tasks, err := h.taskRepository.FindAll()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) GetTaskByIdHandler(c echo.Context) error {
	strIndex := c.Param("id")

	id, err := strconv.Atoi(strIndex)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	task, err := h.taskRepository.FindById(int64(id))

	if err == sql.ErrNoRows {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) CreateTaskHandler(c echo.Context) error {
	task := new(models.Task)

	if err := c.Bind(task); err != nil {
		return err
	}

	task, err := h.taskRepository.Save(*task)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) UpdateTaskHandler(c echo.Context) error {
	strIndex := c.Param("id")

	id, err := strconv.Atoi(strIndex)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	task, err := h.taskRepository.FindById(int64(id))

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

	updatedTask, err := h.taskRepository.Save(*task)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updatedTask)
}

func (h *TaskHandler) ToggleTaskHandler(c echo.Context) error {
	strIndex := c.Param("id")

	id, err := strconv.Atoi(strIndex)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	task, err := h.taskRepository.FindById(int64(id))

	if err == sql.ErrNoRows {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	task.Finished = !task.Finished

	updatedTask, err := h.taskRepository.Save(*task)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updatedTask)
}

func (h *TaskHandler) DeleteTaskByIdHandler(c echo.Context) error {
	strIndex := c.Param("id")

	id, err := strconv.Atoi(strIndex)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Parameter id cannot be converted to integer")
	}

	_, err = h.taskRepository.FindById(int64(id))

	if err == sql.ErrNoRows {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err = h.taskRepository.DeleteById(int64(id)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusNoContent, nil)
}
