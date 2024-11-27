package main

import (
	"fmt"
	"log"
	"os"
	"task_api/internal/handler"
	"task_api/internal/persistence"
	"task_api/internal/repository"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const TASK_ROUTE = "/tasks"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	m, err := migrate.New(
		"file://migrations",
		fmt.Sprintf("mysql://%v:%v@tcp(localhost:3306)/task_api", os.Getenv("DBUSER"), os.Getenv("DBPASS")),
	)

	if err != nil {
		log.Fatalf("Erro ao criar migrate: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Erro ao aplicar migrações: %v", err)
	}

	log.Println("Migrações aplicadas com sucesso!")

	e := echo.New()
	e.Use(middleware.Recover())

	db := persistence.CreateConnection()

	taskRepository := repository.NewTaskRepository(db)
	taskHandler := handler.NewTaskHandler(taskRepository)

	e.GET(TASK_ROUTE, taskHandler.GetTasksHandler)
	e.GET(TASK_ROUTE+"/:id", taskHandler.GetTaskByIdHandler)
	e.POST(TASK_ROUTE, taskHandler.CreateTaskHandler)
	e.PUT(TASK_ROUTE+"/:id", taskHandler.UpdateTaskHandler)
	e.PATCH(TASK_ROUTE+"/:id/toggle", taskHandler.ToggleTaskHandler)
	e.DELETE(TASK_ROUTE+"/:id", taskHandler.DeleteTaskByIdHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
