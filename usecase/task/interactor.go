package task

import "github.com/anraku/tsk/models"

type taskInteractor struct {
	taskRepo TaskRepository
}

func NewTaskInteractor(t TaskRepository) TaskInteractor {
	return &taskInteractor{
		taskRepo: t,
	}
}

func (r *taskInteractor) Fetch() ([]*models.Task, error) {
	result, err := r.taskRepo.Fetch()
	return result, err
}

func (r *taskInteractor) FetchUnDone() ([]*models.Task, error) {
	result, err := r.taskRepo.FetchUnDone()
	return result, err
}

func (r *taskInteractor) FetchDone() ([]*models.Task, error) {
	result, err := r.taskRepo.FetchDone()
	return result, err
}

func (r *taskInteractor) GetByID(id int) (*models.Task, error) {
	result, err := r.taskRepo.GetByID(id)
	return result, err
}

func (r *taskInteractor) Create(t *models.Task) error {
	err := r.taskRepo.Create(t)
	return err
}

func (r *taskInteractor) Update(t *models.Task) error {
	err := r.taskRepo.Update(t)
	return err
}

func (r *taskInteractor) Delete(id int) error {
	err := r.taskRepo.Delete(id)
	return err
}
