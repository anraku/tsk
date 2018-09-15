package task

import "github.com/anraku/tsk/models"

type TaskRepository interface {
	Fetch() ([]*models.Task, error)
	FetchUnDone() ([]*models.Task, error)
	FetchDone() ([]*models.Task, error)
	GetByID(id int) (models.Task, error)
	Create(t *models.Task) error
	Update(t *models.Task) error
	Delete(id int) error
}

type TaskInteractor interface {
	Fetch() ([]*models.Task, error)
	FetchUnDone() ([]*models.Task, error)
	FetchDone() ([]*models.Task, error)
	GetByID(id int) (*models.Task, error)
	Create(t *models.Task) error
	Update(t *models.Task) error
	Delete(id int) error
}
