package models

import (
    "time"
)

type Exception struct {
    // Fields
    ID uint `gorm:"primaryKey" json:"id"`
    Name string `gorm:"uniqueIndex;not null" json:"name"`
    StackTrace string `gorm:"uniqueIndex;not null" json:"stack_trace"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`

    // Relationships
    ApplicationID int `gorm:"not null"`
    Application Application
}
