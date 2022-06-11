package model

import (
	"github.com/google/uuid"
	"time"
)

type Language struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid"`
	RequestStruct
	Date time.Time
}

type RequestStruct struct {
	Email string
	Hours float32
	Name  string
}
