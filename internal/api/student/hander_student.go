package student

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "zy-apilabs/internal/models"
)

//curl -X POST "localhost:8080/v1/student/listStudent"
func ListStudent(c *gin.Context) {
    students, err := models.ListStudent()
    if err != nil {
        c.JSON(http.StatusOK, gin.H{"code": 0, "err": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "success", "data": students})
}

//curl -X POST "localhost:8080/v1/student/listStudentWithClass" -d '{"class_id":1}'
func ListStudentWithClass(c *gin.Context) {
    var student models.Student
    if err := c.BindJSON(&student); err != nil {
        c.JSON(http.StatusOK, gin.H{"code": 0, "err": err.Error()})
        return
    }
    if student.ClassId <= 0 {
        c.JSON(http.StatusOK, gin.H{"code": 0, "err": "参数错误"})
        return
    }
    students, err := models.GetOneClassStudents(student.ClassId)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{"code": 0, "err": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "success", "data": students})
}

//curl -X POST "localhost:8080/v1/student/saveStudent" -d '{"name":"zo1","sex":1,"age":10,"class_id":1,"grade_id":1}'
func SaveStudent(c *gin.Context) {
    var student models.Student
    if err := c.BindJSON(&student); err != nil {
        c.JSON(http.StatusOK, gin.H{"code": 0, "err": err.Error()})
        return
    }

    if err := models.CreateStudent(&student); err != nil {
        c.JSON(http.StatusOK, gin.H{"code": 0, "err": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "success"})
}

func DelStudent(c *gin.Context) {

}

//curl -X POST "localhost:8080/v1/student/listClass"
func ListClass(c *gin.Context) {
    classes, err := models.ListClass()
    if err != nil {
        c.JSON(http.StatusOK, gin.H{"code": 0, "err": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "success", "data": classes})
}

//curl -X POST "localhost:8080/v1/student/saveClass" -d '{"grade_id":1,"name":"class1"}'
func SaveClass(c *gin.Context) {
    var class models.Class
    if err := c.BindJSON(&class); err != nil {
        c.JSON(http.StatusOK, gin.H{"code": 0, "err": err.Error()})
        return
    }

    if err := models.CreateClass(&class); err != nil {
        c.JSON(http.StatusOK, gin.H{"code": 0, "err": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "success"})
}

//curl -X POST "localhost:8080/v1/student/listGrade"
func ListGrade(c *gin.Context) {
    grades, err := models.ListGrade()
    if err != nil {
        c.JSON(http.StatusOK, gin.H{"code": 0, "err": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "success", "data": grades})

}

//curl -X POST "localhost:8080/v1/student/saveGrade" -d '{"name":"grade1"}'
func SaveGrade(c *gin.Context) {
    var grade models.Grade
    if err := c.BindJSON(&grade); err != nil {
        c.JSON(http.StatusOK, gin.H{"code": 0, "err": err.Error()})
        return
    }

    if err := models.CreateGrade(&grade); err != nil {
        c.JSON(http.StatusOK, gin.H{"code": 0, "err": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "success"})
}
