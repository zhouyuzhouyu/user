package models

import (
    "gorm.io/gorm"
    "zy-apilabs/internal/db"
)

type Student struct {
    gorm.Model
    Name    string `json:"name" gorm:"not null"`
    Sex     int    `json:"sex"`
    Age     int    `json:"age"`
    ClassId int    `json:"class_id"`
    GradeId int    `json:"grade_id"`
    Class   ClassInfo
    Grade   Grade
}

type ClassInfo struct {
    gorm.Model
    Name    string `json:"name"`
    GradeId int    `json:"grade_id"`
}

func (ClassInfo) TableName() string {
    return "class"
}
func (Student) TableName() string {
    return "student"
}

// 添加
func CreateStudent(student *Student) (err error) {
    err = db.GetMysqlDB().Create(student).Error
    return
}

// 删除
func DeleteStudent(id int) (err error) {
    err = db.GetMysqlDB().Where("id=?", id).Delete(&Student{}).Error
    return
}

// 修改
func UpdateStudent(student *Student) (err error) {
    err = db.GetMysqlDB().Save(student).Error
    return
}

// 获取全部学生
func ListStudent() (students []Student, err error) {
    err = db.GetMysqlDB().Preload("Grade").Preload("Class").Limit(200).Find(&students).Error
    return
}

// 获取某个班级的全部学生
func GetOneClassStudents(id int) (students []Student, err error) {
    err = db.GetMysqlDB().Where("class_id = ?", id).Preload("Grade").Preload("Class").Find(&students).Error
    return
}
