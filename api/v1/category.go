package v1

import (
    "GinBlog/model"
    "GinBlog/utils/errmsg"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    "log"
)

//check category exist

//query article from category


// add category
func AddCate(c *gin.Context) {
    //
    var data model.Category
    err := c.ShouldBindJSON(&data)
    if err != nil {
         log.Fatal(err)
    }
    code = model.CheckCategory(data.Name)
    if code == errmsg.SUCCESS {
        model.CreateCate(&data)
    }
    if code == errmsg.ERROR_CATENAME_USED {
        code = errmsg.ERROR_CATENAME_USED
    }

    c.JSON(http.StatusOK,gin.H{
        "status":code,
        "data":data,
        "message":errmsg.GetErrMsg(code),
    })
}

//query categorys
func GetCates(c *gin.Context) {
    //
    pageSize,_ := strconv.Atoi(c.Query("pagesize"))
    pageNum,_ := strconv.Atoi(c.Query("pagenum"))
    if pageSize == 0 {
        pageSize = -1
    }
    if pageNum == 0 {
        pageNum = -1
    }
    data,total := model.GetCates(pageSize,pageNum)
    code = errmsg.SUCCESS
    c.JSON(http.StatusOK,gin.H{
        "status":code,
        "data":data,
        "total":total,
        "message":errmsg.GetErrMsg(code),
    })
}

//edit category
func EditCate(c *gin.Context) {
    //
    var data model.Category
    id,_ := strconv.Atoi(c.Param("id"))
    err := c.ShouldBindJSON(&data)
    if err != nil {
         log.Fatal(err)
    }
    code = model.CheckCategory(data.Name)
    if code == errmsg.SUCCESS {
        model.EditCate(id,&data)
    }
    if code == errmsg.ERROR_CATENAME_USED {
        c.Abort()
    }
    c.JSON(http.StatusOK,gin.H{
        "status":code,
        "message":errmsg.GetErrMsg(code),
    })
}

//delete category
func DeleteCate(c *gin.Context) {
    //
    id,_ := strconv.Atoi(c.Param("id"))
    code = model.DeleteCate(id)

    c.JSON(http.StatusOK,gin.H{
        "status":code,
        "message":errmsg.GetErrMsg(code),
    })
}
