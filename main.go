package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Task struct {
	Name     string `json:"name"`
	Finished bool   `json:"finished"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()
	e.Use(middleware.Recover())

	tasks := make([]Task, 0)

	e.GET("/tasks", func(c echo.Context) error {
		return c.JSON(http.StatusOK, tasks)
	})

	e.POST("/tasks", func(c echo.Context) error {
		task := new(Task)

		if err := c.Bind(task); err != nil {
			return err
		}

		tasks = append(tasks, *task)

		return c.JSON(http.StatusCreated, task)
	})

	e.PUT("/tasks/:id", func(c echo.Context) error {
		strIndex := c.Param("id")

		id, err := strconv.Atoi(strIndex)
		index := id - 1

		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{"Parameter id cannot be converted to integer"})
		}

		if index < 0 || index >= len(tasks) {
			return c.JSON(http.StatusNotFound, ErrorResponse{fmt.Sprintf("Task with id %v not found", id)})
		}

		updatedTask := new(Task)

		if err := c.Bind(updatedTask); err != nil {
			return err
		}

		var task = tasks[index]
		task.Name = updatedTask.Name
		tasks[index] = task

		return c.JSON(http.StatusOK, task)
	})

	e.PATCH("/tasks/:id/toggle", func(c echo.Context) error {
		strIndex := c.Param("id")

		id, err := strconv.Atoi(strIndex)
		index := id - 1

		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{"Parameter id cannot be converted to integer"})
		}

		if index < 0 || index >= len(tasks) {
			return c.JSON(http.StatusNotFound, ErrorResponse{fmt.Sprintf("Task with id %v not found", id)})
		}

		var task = tasks[index]
		task.Finished = !task.Finished
		tasks[index] = task

		return c.JSON(http.StatusOK, task)
	})

	e.DELETE("/tasks/:id", func(c echo.Context) error {
		strIndex := c.Param("id")

		id, err := strconv.Atoi(strIndex)
		index := id - 1

		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{"Parameter id cannot be converted to integer"})
		}

		if index < 0 || index >= len(tasks) {
			return c.JSON(http.StatusNotFound, ErrorResponse{fmt.Sprintf("Task with id %v not found", id)})
		}

		tasks = append(tasks[:index], tasks[index+1:]...)

		return c.JSON(http.StatusNoContent, nil)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
