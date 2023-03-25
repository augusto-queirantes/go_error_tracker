package models

import (
    "gorm.io/gorm"
)

type Error struct {
    gorm.Model

    Name string `gorm:"uniqueIndex;not null"`
    StackTrace string `gorm:"uniqueIndex;not null"`
}
