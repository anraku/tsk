package task

import "github.com/labstack/echo"

type TaskHandler struct {
	taskUsecase TaskUsecase
}

func NewTaskHandler(e *echo.Echo, tu TaskUsecase) {
	handler := &TaskHandler{
		taskUsecase: tu,
	}
	e.GET("/tasks", handler.FetchTasks)
	e.GET("/task/:id", handler.GetByID)
	e.POST("/task/create", handler.Create)
	e.PUT("/task/:id", handler.Update)
	e.DELETE("/task/:id", handler.Delete)

}

func (t *TaskHandler) FetchTasks(c echo.Context) error {
	return nil
}

func (t *TaskHandler) GetByID(c echo.Context) error {
	return nil
}

func (t *TaskHandler) Create(c echo.Context) error {
	return nil
}

func (t *TaskHandler) Update(c echo.Context) error {
	return nil
}

func (t *TaskHandler) Delete(c echo.Context) error {
	return nil
}
