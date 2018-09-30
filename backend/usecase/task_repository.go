package usecase

import "github.com/anraku/tsk/backend/domain"

type TaskRepository interface {
	Fetch() ([]domain.Task, error)
	FetchUnDone() ([]domain.Task, error)
	FetchDone() ([]domain.Task, error)
	GetByID(int) (domain.Task, error)
	Create(domain.Task) error
	Update(domain.Task) error
	Delete(int) error
}
