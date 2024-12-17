package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	UserID      int    `form:"user_id" json:"user_id"`
	VehicleID   int    `form:"vehicle_id" json:"vehicle_id"`
	InventoryID int    `form:"inventory_id" json:"inventory_id"`
	Destination string `form:"destination" json:"destination"`
	Status      string `form:"status" json:"status"`
}
