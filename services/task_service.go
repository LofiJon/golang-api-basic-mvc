package services

import (
	"todo-api/models"
	"todo-api/repositories"
)

type TaskService interface {
	CreateTask(task *models.Task) error
	GetTasks() ([]models.Task, error)
	GetTask(id uint) (models.Task, error)
	UpdateTask(task *models.Task) error
	DeleteTask(id uint) error
}

type taskService struct {
	repository repositories.TaskRepository
}

func NewTaskService(repo repositories.TaskRepository) TaskService {
	return &taskService{repo}
}

func (s *taskService) CreateTask(task *models.Task) error {
	return s.repository.CreateTask(task)
}

func (s *taskService) GetTasks() ([]models.Task, error) {
	return s.repository.GetTasks()
}

func (s *taskService) GetTask(id uint) (models.Task, error) {
	return s.repository.GetTask(id)
}

func (s *taskService) UpdateTask(task *models.Task) error {
	return s.repository.UpdateTask(task)
}

func (s *taskService) DeleteTask(id uint) error {
	return s.repository.DeleteTask(id)
}
