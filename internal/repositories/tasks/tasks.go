package tasks

import (
	"task-api/internal/models"
)

type Tasks struct {
	tasks []models.Task
}

func New() *Tasks {
	return &Tasks{tasks: make([]models.Task, 0)}
}

/*
func  GetTask(id uuid.UUID) *models.Task {

}

func GetAllTask() []*models.Task {

}

func AddTask(newTask models.Task) {

}

func UpdateTask(id uuid.UUID) {

}
*/
