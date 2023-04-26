package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"com.zoyu/user/db"
	"com.zoyu/user/model"
)

// 注册接口
func registe(w http.ResponseWriter, r *http.Request) {
	// 读取请求体中的数据
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// 判断用户名是否已存在
	if _, err := db.FindUserByUsername(user.Username); err == nil {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}

	// 将用户信息存入数据库中
	if err := db.InsertUser(user); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 返回注册成功的消息
	fmt.Fprint(w, "Register successfully")
}

// 登录接口
func login(w http.ResponseWriter, r *http.Request) {
	// 读取请求体中的数据
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	// 判断用户名和密码是否正确
	if userindb, err := db.FindUserByUsername(user.Username); err == nil {
		// 登录成功
		if userindb.Password == user.Password {
			fmt.Fprint(w, "Login successfully")
		} else {
			// 登录失败
			http.Error(w, "Invalid password", http.StatusUnauthorized)
		}
	} else {
		// 登录失败
		http.Error(w, "Invalid username, please registe first", http.StatusUnauthorized)
	}
}

// 退出登录接口
func logout(w http.ResponseWriter, r *http.Request) {
	// 退出登录的逻辑
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user.Username); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, "Logout successfully")
}
