package task

import (
	"github.com/anraku/tsk/models"
	"github.com/jinzhu/gorm"
)

const table = "task"

type taskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{
		DB: db,
	}
}

func (r *taskRepository) Fetch() ([]*models.Task, error) {
	return nil, nil
}

func (r *taskRepository) GetByID(id int) (*models.Task, error) {
	return nil, nil
}

func (r *taskRepository) Create(t *models.Task) error {
	return nil
}

func (r *taskRepository) Update(t *models.Task) error {
	return nil
}

func (r *taskRepository) Delete(id int) error {
	return nil
}
