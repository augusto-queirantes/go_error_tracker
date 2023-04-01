package models

import (
    "time"
)

type Exception struct {
    // Fields
    ID uint `gorm:"primaryKey" json:"id"`
    Name string `gorm:"not null" json:"name"`
    StackTrace string `gorm:"not null" json:"stack_trace"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`

    // Relationships
    ApplicationID uint `gorm:"not null" json:"application_id"`
    Application Application
}
