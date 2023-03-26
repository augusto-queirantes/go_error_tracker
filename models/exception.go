package models

import (
    "gorm.io/gorm"
)

type Exception struct {
    gorm.Model

    Name string `gorm:"uniqueIndex;not null"`
    StackTrace string `gorm:"uniqueIndex;not null"`
}
