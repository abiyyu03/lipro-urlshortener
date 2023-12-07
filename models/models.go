package models

import (
	"time"

	"gorm.io/gorm"
)

type Request struct {
	gorm.Model
	ID          int64         `gorm:"primary_key"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
	OriginalUrl string        `json:"original_url"`
}

type Response struct {
	gorm.Model
	ID          int64         `gorm:"primaryKey"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
	// Xra
}
