package repositories

import (
	"service-golang/models"

	"gorm.io/gorm"
)

type InventoryRepository struct{
    DB *gorm.DB
}

func NewInventoryRepository (db *gorm.DB) *InventoryRepository{
    return &InventoryRepository{DB: db}
}

func (r *InventoryRepository) CreateInventory(inventory *models.Inventory)error {
    return r.DB.Create(inventory).Error
}

func (r *InventoryRepository) GetAllInventory()([]models.Inventory, error){
    var inventories []models.Inventory
    if err := r.DB.Find(&inventories).Error; err != nil{
        return nil, err
    }

    return inventories, nil
}

func (r *InventoryRepository) UpdateInventory(inventory *models.Inventory, id uint) error{
    return r.DB.Model(&models.Inventory{}).Where("id=?", id).Updates(inventory).Error
}

func (r *InventoryRepository) DeleteInventory(id uint) error{
    return r.DB.Delete(&models.Inventory{}, id).Error
}