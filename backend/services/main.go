package main

import (
	"github.com/anraku/tsk/backend/infrastructure"
	"github.com/anraku/tsk/backend/interfaces"
	"github.com/anraku/tsk/backend/usecase"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	handler := infrastructure.NewSqlHandler()
	defer handler.DB.Close()

	e := echo.New()
	e.Use(middleware.Logger())

	// Repository
	taskRepo := infrastructure.NewTaskRepository(handler)

	// Interactor
	taskInteractor := usecase.NewTaskInteractor(taskRepo)

	// Response
	res := interfaces.NewResponse()
	// Handler
	infrastructure.NewTaskHandler(e, taskInteractor, res)

	e.Logger.Fatal(e.Start(":1323"))
}
