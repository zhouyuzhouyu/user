package router

import (
    "zy-apilabs/internal/api/student"
    "zy-apilabs/internal/api/user"

    "github.com/gin-gonic/gin"
)

func SetApiRouter(serve *gin.Engine) {
    //v1
    v1Group := serve.Group("/v1")

    //user
    userGroup := v1Group.Group("/user")
    userGroup.POST("/logout", user.Logout)
    userGroup.POST("/register", user.Register)
    userGroup.POST("/login", user.Login)

    //student
    studentGroup := v1Group.Group("/student")
    studentGroup.POST("/listStudent", student.ListStudent)
    studentGroup.POST("/listStudentWithClass", student.ListStudentWithClass)
    studentGroup.POST("/saveStudent", student.SaveStudent)
    studentGroup.POST("/delStudent", student.DelStudent)
    studentGroup.POST("/listClass", student.ListClass)
    studentGroup.POST("/saveClass", student.SaveClass)
    studentGroup.POST("/listGrade", student.ListGrade)
    studentGroup.POST("/saveGrade", student.SaveGrade)

    //userGroup.POST("/get", user.Get)
    //userGroup.POST("/set", user.Set)

    //userGroup.POST("allusers", user.ListUsers)
    //userGroup.GET(":id", user.GetUser)
    //userGroup.POST("createuser", user.CreateUser)
    //userGroup.POST(":id", user.UpdateUser)
    //userGroup.DELETE(":id", user.DeleteUser)

}
