package main

import (
	"github.com/anraku/tsk/database"
	"github.com/anraku/tsk/usecase/task"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	e := echo.New()
	e.Use(middleware.Logger())

	// Repository
	taskRepo := task.NewTaskRepository(db)

	// Interactor
	taskInteractor := task.NewTaskInteractor(taskRepo)

	// Handler
	task.NewTaskHandler(e, taskInteractor)

	e.Logger.Fatal(e.Start(":1323"))
}
