package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdateAt  time.Time      `json:"updateAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
	Name      string         `json:"name"`
}
