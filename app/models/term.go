package models

import (
	"time"
)

type Term struct {
	ID           uint64     `gorm:"primary_key" json:"id"`
	Note         string     `json:"note"`
	TimeTime     time.Time  `json:"time_time"`
	NullableTime *time.Time `json:"nullable_time"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}
