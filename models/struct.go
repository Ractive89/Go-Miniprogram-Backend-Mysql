package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	OpenID string `gorm:"type:varchar(32) ; not null" json:"open_id" binding:"required"`
	Name   string `gorm:"varchar(20) ; not null" json:"name" binding:"required"`
	Phone  string `gorm:"varchar(10) ; not null" json:"phone" binding:"required"`
}
