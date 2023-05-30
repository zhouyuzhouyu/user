package db

import (
    "gorm.io/gorm"
)

var mysqlDB *gorm.DB

func GetMysqlDB() *gorm.DB {
    return mysqlDB
}

func SetupMysqlDB(dbClient *gorm.DB) {
    mysqlDB = dbClient
}
