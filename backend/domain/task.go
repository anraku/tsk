package domain

import "time"

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
