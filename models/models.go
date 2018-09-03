package models

import "time"

// CommonColumn has common column in database table structure.
type CommonColumn struct {
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
}

// Tag is database table structure.
type Tag struct {
	ID     uint   `json:"id" gorm:"column:id;unique_index"`
	TaskID uint   `json:"task_id" gorm:"column:task_id;unique_index"`
	Name   string `json:"name" gorm:"column:name;type:varchar(30)"`
	CommonColumn
}

// Task is database table structure.
type Task struct {
	ID          uint       `json:"id" gorm:"column:id;unique_index"`
	Title       string     `json:"title" gorm:"column:title;type:varchar(100)"`
	Description string     `json:"description" gorm:"column:description;type:varchar(2000)"`
	Deadline    *time.Time `json:"deadline" gorm:"column:deadline;type:timestamp"`
	Priority    int        `json:"priority" gorm:"column:priority;type:integer"`
	Status      int        `json:"status" gorm:"column:status;type:integer"`
	ExecuteTime *time.Time `json"execute_time" gorm:"column:execute_time;type:timestamp"`
	CommonColumn
}

// TaskList is database table structure.
type TaskList struct {
	ID          uint   `json:"id" gorm:"column:id;unique_index"`
	TaskID      uint   `json:"task_id" gorm:"column:id;index:task_id"`
	Title       string `json:"title" gorm:"column:title;type:varchar(100)"`
	Description string `json:"description" gorm:"column:description;type:varchar(2000)"`
	CommonColumn
}
