package models

import (
	"time"

	"gorm.io/gorm"
)

type Request struct {
	*gorm.Model
	ID           int64         `json:"id" gorm:"primary_key"`
	CustomShort  string        `json:"short"`
	Expiry       time.Duration `json:"expiry"`
	OriginalUrl  string        `json:"original_url"`
	RequestCount RequestCount  `gorm:"foreignKey:RequestId"`
}

type RequestCount struct {
	*gorm.Model
	ID        int64  `json:"id" gorm:"primary_key"`
	IPAddress string `json:"ip_address"`
	RequestId int64  `json:"request_id"`
}
