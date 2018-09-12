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

func (r *taskRepository) Fetch() (result []*models.Task, err error) {
	err = r.DB.Table(table).Find(&result).Error
	return result, err
}

func (r *taskRepository) FetchUnDone() (result []*models.Task, err error) {
	err = r.DB.Table(table).Where("done = 0").Find(&result).Error
	return result, err
}

func (r *taskRepository) FetchDone() (result []*models.Task, err error) {
	err = r.DB.Table(table).Where("done = 1").Find(&result).Error
	return result, err
}

func (r *taskRepository) GetByID(id int) (result *models.Task, err error) {
	err = r.DB.Table(table).Where("id = ?", id).First(&result).Error
	return
}

func (r *taskRepository) Create(t *models.Task) (err error) {
	err = r.DB.Table(table).Create(&t).Error
	return
}

func (r *taskRepository) Update(t *models.Task) (err error) {
	err = r.DB.Table(table).Update(&t).Error
	return
}

func (r *taskRepository) Delete(id int) (err error) {
	err = r.DB.Table(table).Where("id = ?", id).Delete(models.Task{}).Error
	return
}
