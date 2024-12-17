package services

import (
	"service-golang/models"
	"service-golang/repositories"
)

type TaskService struct {
	TaskRepository *repositories.TaskRepository
}

func NewTaskService(repo *repositories.TaskRepository) *TaskService {
	return &TaskService{TaskRepository: repo}
}

func (s *TaskService) CreateTask(task *models.Task) error {
	return s.TaskRepository.CreateTask(task)
}

func (s *TaskService) GetAllTask() ([]models.Task, error) {
	return s.TaskRepository.GetAllTask()
}

func (s *TaskService) UpdateTask(task *models.Task, id uint) error {
	return s.TaskRepository.UpdateTask(task, id)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.TaskRepository.DeleteTask(id)
}
