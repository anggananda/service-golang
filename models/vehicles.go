package models

import "gorm.io/gorm"

type Vehicles struct{
    gorm.Model

    LicensePlate string `form:"license_plate" json:"license_plate"`
    Brand string `form:"brand" json:"brand"`
    Status string `form:"status" json:"status" gorm:"default:'active'"`
}