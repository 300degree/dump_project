package domain

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	Id       uuid.UUID `gorm:"type:uuid; primaryKey"`
	Email    string    `gorm:"type:varchar(120); unqiue; not null"`
	Username string    `gorm:"type:varchar(120); unqiue; not null"`
	Password string    `gorm:"type:varchar(60)"`
	CreateAt time.Time `gorm:"autoCreateTime"`
	UpdateAt time.Time `gorm:"autoUpdateTime"`
}
