package repositories

import (
	"service-golang/models"

	"gorm.io/gorm"
)

type VehiclesRepository struct {
	DB *gorm.DB
}

func NewVehiclesRepository(db *gorm.DB) *VehiclesRepository {
	return &VehiclesRepository{DB: db}
}

func (r *VehiclesRepository) CreateVehicles(vehicles *models.Vehicles) error {
	return r.DB.Create(vehicles).Error
}

func (r *VehiclesRepository) GetAllVehicles() ([]models.Vehicles, error) {
	var vehicles []models.Vehicles
	if err := r.DB.Find(&vehicles).Error; err != nil {
		return nil, err
	}

	return vehicles, nil
}

func (r *VehiclesRepository) UpdateVehicles(vehicles *models.Vehicles, id uint) error {
	return r.DB.Model(&models.Vehicles{}).Where("id=?", id).Updates(vehicles).Error
}

func (r *VehiclesRepository) DeleteVehicles(id uint) error {
	return r.DB.Delete(&models.Vehicles{}, id).Error
}
