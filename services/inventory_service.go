package services

import (
	"service-golang/models"
	"service-golang/repositories"
)

type InventoryService struct{
    InventoryRepository *repositories.InventoryRepository
}

func NewInventoryService(repo *repositories.InventoryRepository) *InventoryService{
    return &InventoryService{InventoryRepository: repo}
}

func (s *InventoryService) CreateInventory(inventory *models.Inventory) error{
    return s.InventoryRepository.CreateInventory(inventory)
}

func (s *InventoryService) GetAllInventory()([]models.Inventory, error){
    return s.InventoryRepository.GetAllInventory()
}

func (s *InventoryService) UpdateInventory(inventory *models.Inventory, id uint) error{
    return s.InventoryRepository.UpdateInventory(inventory, id)
}

func (s *InventoryService) DeleteInventory(id uint) error{
    return s.InventoryRepository.DeleteInventory(id)
}