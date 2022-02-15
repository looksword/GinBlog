package v1

import (
    "GinBlog/model"
    "github.com/gin-gonic/gin"
    "net/http"
    "GinBlog/utils/errmsg"
    "log"
)

func UpLoad(c *gin.Context) {
    file,fileHeader,err := c.Request.FormFile("file")
    if err != nil {
        log.Fatal(err)
    }
    fileSize := fileHeader.Size
    url,code := model.UpLoadFile(file,fileSize)
    
    c.JSON(http.StatusOK,gin.H{
        "status":code,
        "message":errmsg.GetErrMsg(code),
        "url":url,
    })
}

