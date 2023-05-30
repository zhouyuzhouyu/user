package models

import (
    "crypto/md5"
    "encoding/hex"
    "gorm.io/gorm"
    "zy-apilabs/internal/db"
)

//type User struct {
//    Id         int    `json:"id"`
//    Username   string `json:"username"`
//    Password   string `json:"password"`
//    Email      string `json:"email"`
//    RoleId     int    `json:"roleId"`
//    LoginCount int    `json:"login_count"`
//    LastTime   int64  `json:"last_time"`
//    LastIp     string `json:"last_ip"`
//    State      int    `json:"state"`
//    Created    int64  `json:"created"`
//    Updated    int64  `json:"updated"`
//}

type User struct {
    gorm.Model
    Username string `json:"username" binding:"required" gorm:"unique;not null"`
    Password string `json:"password" binding:"required" gorm:"not null"`
}

func (User) TableName() string {
    return "user"
}

const secret = "com.zoyu"

// 加密密码
func EncryptPassword(data []byte) (result string) {
    h := md5.New()
    h.Write([]byte(secret))
    return hex.EncodeToString(h.Sum(data))
}

// 登录
func Login(userName string, password string) (user []*User, err error) {
    // 生成加密密码
    dbClient := db.GetMysqlDB()
    dbClient = dbClient.Where("username = ?", userName).Where("password = ?", EncryptPassword([]byte(password)))
    if err = dbClient.Find(&user).Error; err != nil {
        return nil, err
    }
    return user, nil
}

// 注册
func CreateUser(user *User) (err error) {
    err = db.GetMysqlDB().Create(user).Error
    return
}
