package models

import (
    "gorm.io/gorm"
    "zy-apilabs/internal/db"
)

type Class struct {
    gorm.Model
    Name     string        `json:"name" gorm:"unique;not null"`
    GradeId  int           `json:"grade_id"`
    Students []StudentInfo `json:"students"`
    Grade    Grade
}

type StudentInfo struct {
    gorm.Model
    Name    string `json:"name"`
    ClassId int    `json:"class_id"`
}

func (StudentInfo) TableName() string {
    return "student"
}
func (Class) TableName() string {
    return "class"
}

func CreateClass(class *Class) (err error) {
    err = db.GetMysqlDB().Create(class).Error
    return
}

func ListClass() (classes []Class, err error) {
    err = db.GetMysqlDB().Model(new(Class)).Preload("Grade").Preload("Students").Find(&classes).Error
    return
}
