package models

import (
    "gorm.io/gorm"
)

type Application struct {
    gorm.Model

    Name string `gorm:"uniqueIndex;not null"`
}
