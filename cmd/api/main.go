package main

import (
    "context"
    "fmt"
    "github.com/gin-gonic/gin"
    rkboot "github.com/rookie-ninja/rk-boot/v2"
    rkmysql "github.com/rookie-ninja/rk-db/mysql"
    rkgin "github.com/rookie-ninja/rk-gin/v2/boot"
    "net/http"
    "zy-apilabs/internal/db"
    "zy-apilabs/internal/models"
    "zy-apilabs/internal/router"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample rk-demo server.
// @termsOfService http://swagger.io/terms/

// @securityDefinitions.basic BasicAuth

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
    // Create a new boot instance.
    boot := rkboot.NewBoot()
    // Bootstrap
    boot.Bootstrap(context.TODO())

    // Auto migrate database and init global userDb variable
    mysqlEntry := rkmysql.GetMySqlEntry("user-db")
    mysqlDB := mysqlEntry.GetDB("sql_zoyu")
    if !mysqlDB.DryRun {
        mysqlDB.AutoMigrate(
            new(models.User),
            new(models.Student),
            new(models.Class),
            new(models.Grade),
        )
    }
    db.SetupMysqlDB(mysqlDB)

    //redisEntry := rkredis.GetRedisEntry("redis")
    //redisClient, _ := redisEntry.GetClient()
    //zyredis.Setup(redisClient)

    // Register handler
    entry := rkgin.GetGinEntry("zy-apilabs")
    entry.Router.GET("/v1/greeter", Greeter)
    entry.Router.GET("/v1/hello", Hello)
    router.SetApiRouter(entry.Router)

    boot.WaitForShutdownSig(context.TODO())
}

// Greeter handler
// @Summary Greeter
// @Id 1
// @Tags Hello
// @version 1.0
// @Param name query string true "name"
// @produce application/json
// @Success 200 {object} GreeterResponse
// @Router /v1/greeter [get]
func Greeter(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, &GreeterResponse{
        Message: fmt.Sprintf("Hello %s!", ctx.Query("name")),
    })
}

type GreeterResponse struct {
    Message string
}

func Hello(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, &HelloResponse{
        Message: fmt.Sprintf("Hello %s!", ctx.Query("name")),
    })
}

type HelloResponse struct {
    Message string
}
