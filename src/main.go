package main

import (
	"log"
	"task_api/src/handlers"
	"task_api/src/persistence"
	"task_api/src/repositories"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const TASK_ROUTE = "/tasks"

func main() {
	e := echo.New()
	e.Use(middleware.Recover())

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db := persistence.CreateConnection()

	taskRepository := repositories.NewTaskRepository(db)
	taskHandler := handlers.NewTaskHandler(taskRepository)

	e.GET(TASK_ROUTE, taskHandler.GetTasksHandler)
	e.GET(TASK_ROUTE+"/:id", taskHandler.GetTaskByIdHandler)
	e.POST(TASK_ROUTE, taskHandler.CreateTaskHandler)
	e.PUT(TASK_ROUTE+"/:id", taskHandler.UpdateTaskHandler)
	e.PATCH(TASK_ROUTE+"/:id/toggle", taskHandler.ToggleTaskHandler)
	e.DELETE(TASK_ROUTE+"/:id", taskHandler.DeleteTaskByIdHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
