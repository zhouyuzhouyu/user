package models

import (
    "gorm.io/gorm"
    "zy-apilabs/internal/db"
)

type Grade struct {
    gorm.Model
    Name string `json:"name" gorm:"unique;not null"`
}

func (Grade) TableName() string {
    return "grade"
}

func CreateGrade(grade *Grade) (err error) {
    err = db.GetMysqlDB().Create(grade).Error
    return
}

func ListGrade() (grades []Grade, err error) {
    err = db.GetMysqlDB().Find(&grades).Error
    return
}
