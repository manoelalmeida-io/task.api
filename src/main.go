package main

import (
	"log"
	"task_api/src/handlers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const ROUTE = "/tasks"

func main() {
	e := echo.New()
	e.Use(middleware.Recover())

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	e.GET(ROUTE, handlers.GetTasksHandler)
	e.GET(ROUTE+"/:id", handlers.GetTaskByIdHandler)
	e.POST(ROUTE, handlers.CreateTaskHandler)
	e.PUT(ROUTE+"/:id", handlers.UpdateTaskHandler)
	e.PATCH(ROUTE+"/:id/toggle", handlers.ToggleTaskHandler)
	e.DELETE(ROUTE+"/:id", handlers.DeleteTaskByIdHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
