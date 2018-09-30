package domain

// Tag is database table structure.
type Tag struct {
	ID     uint   `json:"id" gorm:"column:id;unique_index"`
	TaskID uint   `json:"task_id" gorm:"column:task_id;unique_index"`
	Name   string `json:"name" gorm:"column:name;type:varchar(30)"`
	CommonColumn
}
