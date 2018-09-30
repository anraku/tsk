package infrastructure

import (
	"github.com/anraku/tsk/backend/domain"
)

type TaskRepository struct {
	SqlHandler
}

func NewTaskRepository(handler SqlHandler) *TaskRepository {
	return &TaskRepository{
		SqlHandler: handler,
	}
}

const table = "tasks"

func (r *TaskRepository) Fetch() (result []domain.Task, err error) {
	err = r.DB.Table(table).Find(&result).Error
	return result, err
}

func (r *TaskRepository) FetchUnDone() (result []domain.Task, err error) {
	err = r.DB.Table(table).Where("done = 0").Find(&result).Error
	return result, err
}

func (r *TaskRepository) FetchDone() (result []domain.Task, err error) {
	err = r.DB.Table(table).Where("done = 1").Find(&result).Error
	return result, err
}

func (r *TaskRepository) GetByID(id int) (result domain.Task, err error) {
	err = r.DB.Table(table).Where("id = ?", id).First(&result).Error
	return
}

func (r *TaskRepository) Create(t domain.Task) (err error) {
	err = r.DB.Table(table).Create(t).Error
	return
}

func (r *TaskRepository) Update(t domain.Task) (err error) {
	err = r.DB.Table(table).Update(t).Error
	return
}

func (r *TaskRepository) Delete(id int) (err error) {
	err = r.DB.Table(table).Where("id = ?", id).Delete(&domain.Task{}).Error
	return
}
