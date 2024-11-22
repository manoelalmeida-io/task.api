package main

import (
	"log"
	"task_api/src/handlers"

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

	e.GET(TASK_ROUTE, handlers.GetTasksHandler)
	e.GET(TASK_ROUTE+"/:id", handlers.GetTaskByIdHandler)
	e.POST(TASK_ROUTE, handlers.CreateTaskHandler)
	e.PUT(TASK_ROUTE+"/:id", handlers.UpdateTaskHandler)
	e.PATCH(TASK_ROUTE+"/:id/toggle", handlers.ToggleTaskHandler)
	e.DELETE(TASK_ROUTE+"/:id", handlers.DeleteTaskByIdHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
