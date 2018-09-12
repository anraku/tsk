package task

import "github.com/labstack/echo"

type TaskHandler struct {
	taskInteractor TaskInteractor
}

func NewTaskHandler(e *echo.Echo, tu TaskInteractor) {
	handler := &TaskHandler{
		taskInteractor: tu,
	}
	e.GET("/tasks", handler.FetchTasks)
	e.GET("/task/:id", handler.GetByID)
	e.POST("/task/create", handler.Create)
	e.PUT("/task/:id", handler.Update)
	e.DELETE("/task/:id", handler.Delete)
}

func (h *TaskHandler) FetchTasks(c echo.Context) error {
	result, err := h.taskInteractor.Fetch()
	if err != nil {
		return err
	}
	return c.JSON(result)
}

func (h *TaskHandler) FetchUnDoneTasks(c echo.Context) error {
	result, err := h.taskInteractor.FetchUnDone()
	if err != nil {
		return err
	}
	return c.JSON(result)
}

func (h *TaskHandler) FetchDoneTasks(c echo.Context) error {
	result, err := h.taskInteractor.FetchDone()
	if err != nil {
		return err
	}
	return echo.c.JSON(result)
}

func (h *TaskHandler) GetByID(c echo.Context) error {
	result, err := h.taskInteractor.GetByID()
	if err != nil {
		return err
	}
	return c.JSON(result)
}

func (h *TaskHandler) Create(c echo.Context) error {
	result, err := h.taskInteractor.Create()
	if err != nil {
		return err
	}
	return c.JSON(result)
}

func (h *TaskHandler) Update(c echo.Context) error {
	result, err := h.taskInteractor.Update()
	if err != nil {
		return err
	}
	return c.JSON(result)
}

func (h *TaskHandler) Delete(c echo.Context) error {
	result, err := h.taskInteractor.Delete()
	if err != nil {
		return err
	}
	return c.JSON(result)
}
