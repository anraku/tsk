package infrastructure

import (
	"github.com/anraku/tsk/backend/interfaces"
	"github.com/anraku/tsk/backend/usecase"
	"github.com/labstack/echo"
)

func NewTaskHandler(e *echo.Echo, tu usecase.TaskInteractor, res interfaces.Response) {
	handler := interfaces.TaskHandler{
		Interactor: tu,
		Response:   res,
	}
	e.GET("/tasks", func(c echo.Context) error { return handler.Fetch(c) })
	e.GET("/task/:id", func(c echo.Context) error { return handler.GetByID(c) })
	e.POST("/task/create", func(c echo.Context) error { return handler.Create(c) })
	e.PUT("/task/:id", func(c echo.Context) error { return handler.Update(c) })
	e.DELETE("/task/:id", func(c echo.Context) error { return handler.Delete(c) })
}
