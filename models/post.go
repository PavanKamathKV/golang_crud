package models

import (
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	Title       string `gorm:"type:varchar(255);" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	AssignedTo  string `gorm:"type:text" json:"assignedto"`
}

func GetTask() Task {
	var task Task
	return task
}

func GetTasks() []Task {
	var tasks []Task
	return tasks
}
