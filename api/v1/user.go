package v1

import (
    "GinBlog/model"
    "github.com/gin-gonic/gin"
    "net/http"
    "GinBlog/utils/errmsg"
    "strconv"
)

var code int

// add user
func AddUser(c *gin.Context) {
    //
    var data model.User
    _ = c.ShouldBindJSON(&data)
    code = model.CheckUser(data.Username)
    if code == errmsg.SUCCESS {
        model.CreateUser(&data)
    }
    if code == errmsg.ERROR_USERNAME_USED {
        code = errmsg.ERROR_USERNAME_USED
    }

    c.JSON(http.StatusOK,gin.H{
        "status":code,
        "data":data,
        "message":errmsg.GetErrMsg(code),
    })
}

// query user

// query users
func GetUsers(c *gin.Context) {
    //
    pageSize,_ := strconv.Atoi(c.Query("pagesize"))
    pageNum,_ := strconv.Atoi(c.Query("pagenum"))
    if pageSize == 0 {
        pageSize = -1
    }
    if pageNum == 0 {
        pageNum = -1
    }
    data := model.GetUsers(pageSize,pageNum)
    code = errmsg.SUCCESS
    c.JSON(http.StatusOK,gin.H{
        "status":code,
        "data":data,
        "message":errmsg.GetErrMsg(code),
    })
}

// edit user
func EditUser(c *gin.Context) {
    //
}

// delete user
func DeleteUser(c *gin.Context) {
    //
}
