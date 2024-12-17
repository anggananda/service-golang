package repositories

import (
	"service-golang/models"

	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) CreateTask(task *models.Task) error {
	return r.DB.Create(task).Error
}

func (r *TaskRepository) GetAllTask() ([]models.Task, error) {
	var tasks []models.Task
	if err := r.DB.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepository) UpdateTask(task *models.Task, id uint) error {
	return r.DB.Model(&models.Task{}).Where("id=?", id).Updates(task).Error
}

func (r *TaskRepository) DeleteTask(id uint) error {
	return r.DB.Delete(&models.Task{}, id).Error
}
