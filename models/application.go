package models

import (
    "time"
)

type Application struct {
    // Fields
    ID uint `gorm:"primaryKey" json:"id"`
    Name string `gorm:"uniqueIndex;not null" json:"name"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
