package task

import "github.com/anraku/tsk/models"

type taskUsecase struct {
	taskRepo TaskRepository
}

func NewTaskUsecase(t TaskRepository) TaskUsecase {
	return &taskUsecase{
		taskRepo: t,
	}
}

func (r *taskUsecase) Fetch() ([]*models.Task, error) {
	return nil, nil
}

func (r *taskUsecase) GetByID(id int) (*models.Task, error) {
	return nil, nil
}

func (r *taskUsecase) Create(t *models.Task) error {
	return nil
}

func (r *taskUsecase) Update(t *models.Task) error {
	return nil
}

func (r *taskUsecase) Delete(id int) error {
	return nil
}
