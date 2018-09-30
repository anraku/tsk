package interfaces

import (
	"strconv"

	"github.com/anraku/tsk/backend/domain"
	"github.com/anraku/tsk/backend/usecase"
)

type TaskHandler struct {
	Interactor usecase.TaskInteractor
	Response   Response
}

var testModel = domain.Task{
	Title:       "testtask",
	Description: "test_desc",
}

func (h *TaskHandler) Fetch(c Context) error {
	result, err := h.Interactor.Fetch()
	if err != nil {
		return h.Response.ErrInternalServerError(c, err)
	}
	return h.Response.Success(c, result)
}

func (h *TaskHandler) FetchUnDone(c Context) error {
	result, err := h.Interactor.FetchUnDone()
	if err != nil {
		return h.Response.ErrInternalServerError(c, err)
	}
	return h.Response.Success(c, result)
}

func (h *TaskHandler) FetchDone(c Context) error {
	result, err := h.Interactor.FetchDone()
	if err != nil {
		return h.Response.ErrInternalServerError(c, err)
	}
	return h.Response.Success(c, result)
}

func (h *TaskHandler) GetByID(c Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return h.Response.ErrInternalServerError(c, err)
	}
	result, err := h.Interactor.GetByID(id)
	if err != nil {
		return h.Response.ErrInternalServerError(c, err)
	}
	return h.Response.Success(c, result)
}

func (h *TaskHandler) Create(c Context) error {
	//TODO: create model
	err := h.Interactor.Create(testModel)
	if err != nil {
		return h.Response.ErrInternalServerError(c, err)
	}
	return h.Response.Success(c, nil)
}

func (h *TaskHandler) Update(c Context) error {
	//TODO: create model
	err := h.Interactor.Update(testModel)
	if err != nil {
		return h.Response.ErrInternalServerError(c, err)
	}
	return h.Response.Success(c, nil)
}

func (h *TaskHandler) Delete(c Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return h.Response.ErrInternalServerError(c, err)
	}
	err = h.Interactor.Delete(id)
	if err != nil {
		return h.Response.ErrInternalServerError(c, err)
	}
	return h.Response.Success(c, nil)
}
