package domain

// TaskList is database table structure.
type TaskList struct {
	ID          uint   `json:"id" gorm:"column:id;unique_index"`
	TaskID      uint   `json:"task_id" gorm:"column:id;index:task_id"`
	Title       string `json:"title" gorm:"column:title;type:varchar(100)"`
	Description string `json:"description" gorm:"column:description;type:varchar(2000)"`
	CommonColumn
}
