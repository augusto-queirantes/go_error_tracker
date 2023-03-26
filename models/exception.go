package models

import (
    "gorm.io/gorm"
)

type Exception struct {
    gorm.Model

    // Fields
    Name string `gorm:"uniqueIndex;not null"`
    StackTrace string `gorm:"uniqueIndex;not null"`

    // Relationships
    ApplicationID int `gorm:"not null"`
    Application Application
}
