package v1

import (
    "GinBlog/model"
    "GinBlog/utils/errmsg"
    "github.com/gin-gonic/gin"
    "net/http"
    "log"
    "strconv"
)

// add article
func AddArt(c *gin.Context) {
    //
    var data model.Article
    err := c.ShouldBindJSON(&data)
    if err != nil {
         log.Fatal(err)
    }
    code = model.CreateArt(&data)

    c.JSON(http.StatusOK,gin.H{
        "status":code,
        "data":data,
        "message":errmsg.GetErrMsg(code),
    })
}

//query article from category
func GetCateArt(c *gin.Context){
    id,_ := strconv.Atoi(c.Param("id"))
    pageSize,_ := strconv.Atoi(c.Query("pagesize"))
    pageNum,_ := strconv.Atoi(c.Query("pagenum"))
    if pageSize == 0 {
        pageSize = -1
    }
    if pageNum == 0 {
        pageNum = -1
    }
    data,code := model.GetCateArt(id,pageSize,pageNum)
    c.JSON(http.StatusOK,gin.H{
        "status":code,
        "data":data,
        "message":errmsg.GetErrMsg(code),
    })
}

// query single article
func GetArtInfo(c *gin.Context){
    id,_ := strconv.Atoi(c.Param("id"))
    data,code := model.GetArtInfo(id)
    c.JSON(http.StatusOK,gin.H{
        "status":code,
        "data":data,
        "message":errmsg.GetErrMsg(code),
    })
}


// query articles
func GetArts(c *gin.Context) {
    //
    pageSize,_ := strconv.Atoi(c.Query("pagesize"))
    pageNum,_ := strconv.Atoi(c.Query("pagenum"))
    if pageSize == 0 {
        pageSize = -1
    }
    if pageNum == 0 {
        pageNum = -1
    }
    data,code := model.GetArts(pageSize,pageNum)
    c.JSON(http.StatusOK,gin.H{
        "status":code,
        "data":data,
        "message":errmsg.GetErrMsg(code),
    })
}

// edit article
func EditArt(c *gin.Context) {
    //
    var data model.Article
    id,_ := strconv.Atoi(c.Param("id"))
    err := c.ShouldBindJSON(&data)
    if err != nil {
         log.Fatal(err)
    }
    code = model.EditArt(id,&data)
    c.JSON(http.StatusOK,gin.H{
        "status":code,
        "message":errmsg.GetErrMsg(code),
    })
}

// delete article
func DeleteArt(c *gin.Context) {
    //
    id,_ := strconv.Atoi(c.Param("id"))
    code = model.DeleteArt(id)

    c.JSON(http.StatusOK,gin.H{
        "status":code,
        "message":errmsg.GetErrMsg(code),
    })
}

