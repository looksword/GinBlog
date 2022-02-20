package v1

import (
    "GinBlog/model"
    "GinBlog/utils/validator"
    "github.com/gin-gonic/gin"
    "net/http"
    "GinBlog/utils/errmsg"
    "strconv"
    "log"
)

var code int

// add user
func AddUser(c *gin.Context) {
    //
    var data model.User
    var msg string
    _ = c.ShouldBindJSON(&data)
    msg,code = validator.Validate(&data)
    if code != errmsg.SUCCESS {
        c.JSON(http.StatusOK,gin.H{
            "status":code,
            "message":msg,
        })
        return
    }
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
    data,total := model.GetUsers(pageSize,pageNum)
    code = errmsg.SUCCESS
    c.JSON(http.StatusOK,gin.H{
        "status":code,
        "data":data,
        "total":total,
        "message":errmsg.GetErrMsg(code),
    })
}

// edit user
func EditUser(c *gin.Context) {
    //
    var data model.User
    id,_ := strconv.Atoi(c.Param("id"))
    err := c.ShouldBindJSON(&data)
    if err != nil {
         log.Fatal(err)
    }
    code = model.CheckUser(data.Username)
    if code == errmsg.SUCCESS {
        model.EditUser(id,&data)
    }
    if code == errmsg.ERROR_USERNAME_USED {
        c.Abort()
    }
    c.JSON(http.StatusOK,gin.H{
        "status":code,
        "message":errmsg.GetErrMsg(code),
    })
}

// delete user
func DeleteUser(c *gin.Context) {
    //
    id,_ := strconv.Atoi(c.Param("id"))
    code = model.DeleteUser(id)

    c.JSON(http.StatusOK,gin.H{
        "status":code,
        "message":errmsg.GetErrMsg(code),
    })
}

