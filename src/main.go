package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"task_api/src/data"
	"task_api/src/domain"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	var taskRepository data.TaskRepository

	e.GET("/tasks", func(c echo.Context) error {
		tasks, err := taskRepository.FindAll()

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, tasks)
	})

	e.GET("/tasks/:id", func(c echo.Context) error {
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
	})

	e.POST("/tasks", func(c echo.Context) error {
		task := new(domain.Task)

		if err := c.Bind(task); err != nil {
			return err
		}

		task, err := taskRepository.Save(*task)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusCreated, task)
	})

	e.PUT("/tasks/:id", func(c echo.Context) error {
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

		requestTask := new(domain.Task)

		if err := c.Bind(requestTask); err != nil {
			return err
		}

		task.Name = requestTask.Name

		updatedTask, err := taskRepository.Save(*task)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, updatedTask)
	})

	e.PATCH("/tasks/:id/toggle", func(c echo.Context) error {
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
	})

	e.DELETE("/tasks/:id", func(c echo.Context) error {
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
	})

	e.Logger.Fatal(e.Start(":1323"))
}
