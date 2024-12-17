package services

import (
	"service-golang/models"
	"service-golang/repositories"
)

type VehiclesService struct {
	VehiclesRepository *repositories.VehiclesRepository
}

func NewVehiclesService(repo *repositories.VehiclesRepository) *VehiclesService {
	return &VehiclesService{VehiclesRepository: repo}
}

func (s *VehiclesService) CreateVehicles(vehicles *models.Vehicles) error {
	return s.VehiclesRepository.CreateVehicles(vehicles)
}

func (s *VehiclesService) GetAllVehicles() ([]models.Vehicles, error) {
	return s.VehiclesRepository.GetAllVehicles()
}

func (s *VehiclesService) UpdateVehicles(vehicles *models.Vehicles, id uint) error {
	return s.VehiclesRepository.UpdateVehicles(vehicles, id)
}

func (s *VehiclesService) DeleteVehicles(id uint) error {
	return s.VehiclesRepository.DeleteVehicles(id)
}
