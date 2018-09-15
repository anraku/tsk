package task

import (
	"strconv"

	"github.com/anraku/tsk/models"
	"github.com/anraku/tsk/presenter/response"
	"github.com/labstack/echo"
)

type TaskHandler struct {
	taskInteractor TaskInteractor
}

func NewTaskHandler(e *echo.Echo, tu TaskInteractor) {
	handler := &TaskHandler{
		taskInteractor: tu,
	}
	e.GET("/tasks", handler.Fetch)
	e.GET("/task/:id", handler.GetByID)
	e.POST("/task/create", handler.Create)
	e.PUT("/task/:id", handler.Update)
	e.DELETE("/task/:id", handler.Delete)
}

var testModel = &models.Task{
	Title:       "testtask",
	Description: "test_desc",
}

func (h *TaskHandler) Fetch(c echo.Context) error {
	result, err := h.taskInteractor.Fetch()
	if err != nil {
		return response.ErrInternalServerError(c, err)
	}
	return response.Success(c, result)
}

func (h *TaskHandler) FetchUnDone(c echo.Context) error {
	result, err := h.taskInteractor.FetchUnDone()
	if err != nil {
		return response.ErrInternalServerError(c, err)
	}
	return response.Success(c, result)
}

func (h *TaskHandler) FetchDone(c echo.Context) error {
	result, err := h.taskInteractor.FetchDone()
	if err != nil {
		return response.ErrInternalServerError(c, err)
	}
	return response.Success(c, result)
}

func (h *TaskHandler) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.ErrInternalServerError(c, err)
	}
	result, err := h.taskInteractor.GetByID(id)
	if err != nil {
		return response.ErrInternalServerError(c, err)
	}
	return response.Success(c, result)
}

func (h *TaskHandler) Create(c echo.Context) error {
	//TODO: create model
	err := h.taskInteractor.Create(testModel)
	if err != nil {
		return response.ErrInternalServerError(c, err)
	}
	return response.Success(c, nil)
}

func (h *TaskHandler) Update(c echo.Context) error {
	//TODO: create model
	err := h.taskInteractor.Update(testModel)
	if err != nil {
		return response.ErrInternalServerError(c, err)
	}
	return response.Success(c, nil)
}

func (h *TaskHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return response.ErrInternalServerError(c, err)
	}
	err = h.taskInteractor.Delete(id)
	if err != nil {
		return response.ErrInternalServerError(c, err)
	}
	return response.Success(c, nil)
}
