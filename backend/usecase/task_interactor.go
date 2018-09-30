package usecase

import "github.com/anraku/tsk/backend/domain"

type TaskInteractor struct {
	taskRepo TaskRepository
}

func NewTaskInteractor(t TaskRepository) TaskInteractor {
	return TaskInteractor{
		taskRepo: t,
	}
}

func (r *TaskInteractor) Fetch() ([]domain.Task, error) {
	result, err := r.taskRepo.Fetch()
	return result, err
}

func (r *TaskInteractor) FetchUnDone() ([]domain.Task, error) {
	result, err := r.taskRepo.FetchUnDone()
	return result, err
}

func (r *TaskInteractor) FetchDone() ([]domain.Task, error) {
	result, err := r.taskRepo.FetchDone()
	return result, err
}

func (r *TaskInteractor) GetByID(id int) (domain.Task, error) {
	result, err := r.taskRepo.GetByID(id)
	return result, err
}

func (r *TaskInteractor) Create(t domain.Task) error {
	err := r.taskRepo.Create(t)
	return err
}

func (r *TaskInteractor) Update(t domain.Task) error {
	err := r.taskRepo.Update(t)
	return err
}

func (r *TaskInteractor) Delete(id int) error {
	err := r.taskRepo.Delete(id)
	return err
}
