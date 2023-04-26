package db

import (
	"database/sql"
	"fmt"

	"com.zoyu/user/model"
	_ "github.com/go-sql-driver/mysql"
)

// 定义 MySQL 数据库连接信息
const (
	dbHost     = "remote-host.com"
	dbPort     = "3306"
	dbUser     = "remote-user"
	dbPassword = "password"
	dbName     = "users"
)

/*
CREATE TABLE users (

	id INT(11) NOT NULL AUTO_INCREMENT,
	username VARCHAR(50) NOT NULL,
	password VARCHAR(255) NOT NULL,
	PRIMARY KEY (id),
	UNIQUE KEY username (username)

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
*/
var db *sql.DB

// 初始化数据库连接
func Init() {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	var err error
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

}

// 插入用户信息到数据库
func InsertUser(user model.User) error {
	_, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	if err != nil {
		return err
	}
	return nil
}

// 根据用户名查询用户信息
func FindUserByUsername(username string) (model.User, error) {
	var user model.User
	err := db.QueryRow("SELECT username, password FROM users WHERE username=?", username).Scan(&user.Username, &user.Password)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
