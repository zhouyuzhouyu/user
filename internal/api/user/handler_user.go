package user

import (
    "net/http"
    "zy-apilabs/internal/models"
    "zy-apilabs/internal/views"

    "github.com/gin-gonic/gin"
)

func Logout(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, views.Success())
}

//func Set(ctx *gin.Context) {
//	payload := &KV{}
//
//	if err := ctx.BindJSON(payload); err != nil {
//		ctx.JSON(http.StatusInternalServerError, err)
//		return
//	}
//
//	redisClient := zyredis.RedisClient
//	cmd := redisClient.Set(ctx.Request.Context(), payload.Key, payload.Value, time.Minute*10)
//
//	if cmd.Err() != nil {
//		ctx.JSON(http.StatusInternalServerError, cmd.Err())
//		return
//	}
//
//	ctx.Status(http.StatusOK)
//}
//
//func Get(ctx *gin.Context) {
//	key := ctx.Query("key")
//	cmd := zyredis.RedisClient.Get(ctx.Request.Context(), key)
//
//	if cmd.Err() != nil {
//		if cmd.Err() == redis.Nil {
//			ctx.JSON(http.StatusNotFound, "Key not found!")
//		} else {
//			ctx.JSON(http.StatusInternalServerError, cmd.Err())
//		}
//		return
//	}
//
//	payload := &KV{
//		Key:   key,
//		Value: cmd.Val(),
//	}
//
//	ctx.JSON(http.StatusOK, payload)
//}

//func DeleteUser(ctx *gin.Context) {
//    uid, _ := strconv.Atoi(ctx.Param("id"))
//    user := models.User{
//        Model: gorm.Model{
//            ID: uint(uid),
//        },
//    }
//    res := db.GetMysqlDB().Delete(&user)
//    if res.Error != nil {
//        ctx.JSON(http.StatusInternalServerError, res.Error)
//        return
//    }
//    ctx.String(http.StatusOK, "success")
//}
//
//func UpdateUser(ctx *gin.Context) {
//    uid := ctx.Param("id")
//    user := &models.User{
//        Name: ctx.Query("name"),
//    }
//
//    res := db.GetMysqlDB().Where("id = ?", uid).Updates(user)
//
//    if res.Error != nil {
//        ctx.JSON(http.StatusInternalServerError, res.Error)
//        return
//    }
//
//    // get user
//    db.GetMysqlDB().Where("id = ?", uid).Find(user)
//
//    ctx.JSON(http.StatusOK, user)
//}
//
//func GetUser(ctx *gin.Context) {
//    uid := ctx.Param("id")
//    var userList []models.User
//    //userList := &(make([]*User, 0))
//    //userList := make([]*User, 0)
//    //pUserList := &userList
//    res := db.GetMysqlDB().Where("id = ?", uid).Find(&userList)
//
//    if res.Error != nil {
//        ctx.JSON(http.StatusInternalServerError, res.Error)
//        return
//    }
//    ctx.JSON(http.StatusOK, userList)
//
//}
//
//func CreateUser(ctx *gin.Context) {
//    user1 := &models.User{}
//    if err := ctx.BindJSON(user1); err != nil {
//        ctx.JSON(http.StatusBadRequest, "bad request")
//        return
//    }
//
//    name := user1.Name
//    if len(name) <= 0 {
//        ctx.JSON(http.StatusBadRequest, "bad request")
//        return
//    }
//
//    user := &models.User{
//        Name: name,
//    }
//    res := db.GetMysqlDB().Create(user)
//    if res.Error != nil {
//        ctx.JSON(http.StatusInternalServerError, res.Error)
//        return
//    }
//    ctx.JSON(http.StatusOK, user)
//}
//
//func ListUsers(ctx *gin.Context) {
//    var userList []*models.User
//    res := db.GetMysqlDB().Find(&userList)
//
//    if res.Error != nil {
//        ctx.JSON(http.StatusInternalServerError, res.Error)
//        return
//    }
//    ctx.JSON(http.StatusOK, userList)
//}
//
//
// 登录
//curl -X POST "localhost:8080/v1/user/login" -d '{"username":"zoyu1","password":"pw1"}' -i
func Login(c *gin.Context) {
    // 1、 接收参数
    var user models.User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusOK, gin.H{"code": 0, "err": err.Error()})
        return
    }
    // 2、 数据库验证OK
    users, _ := models.Login(user.Username, user.Password)

    if len(users) == 0 {
        c.JSON(http.StatusOK, gin.H{"code": 0, "err": "用户名或者密码错误"})
        return
    }
    // 3、返回token
    //fmt.Println(user[0])
    //token := util.GenerateToken(user[0])
    c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "登陆成功"})
    //"token": token})

}

// 注册
//curl -X POST "localhost:8080/v1/user/register" -d '{"username":"zoyu2","password":"pw2"}' -i
func Register(c *gin.Context) {
    var user models.User
    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusOK, gin.H{"code": 0, "err": err.Error()})
        return
    }

    if len(user.Username) <= 0 || len(user.Password) <= 0 {
        c.JSON(http.StatusOK, gin.H{"code": 0, "err": "用户名或者密码不能为空"})
        return
    }
    user.Password = models.EncryptPassword([]byte(user.Password))

    if err := models.CreateUser(&user); err != nil {
        c.JSON(http.StatusOK, gin.H{"code": 0, "err": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "注册成功"})
}
