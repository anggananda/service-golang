package models

import "gorm.io/gorm"

type Inventory struct{
    gorm.Model
    Name string `form:"name" json:"name"`
    Description string `form:"description" json:"description"`
    Quantity int `form:"quantity" json:"quantity"`
    Status string `form:"status" json:"status" gorm:"deafault:'availabel'"`
}