package routes

import (
    "GinBlog/utils"
    "GinBlog/middleware"
    "github.com/gin-gonic/gin"
    "GinBlog/api/v1"
)


func InitRouter() {
    gin.SetMode(utils.AppMode)
    r := gin.New()
    r.Use(middleware.Logger())
    r.Use(gin.Recovery())
    
    auth := r.Group("api/v1")
    auth.Use(middleware.JwtToken())
    {
        //User Router  
        auth.PUT("user/:id",v1.EditUser)
        auth.DELETE("user/:id",v1.DeleteUser)

        //Category Router
        auth.POST("category/add",v1.AddCate)
        auth.PUT("category/:id",v1.EditCate)
        auth.DELETE("category/:id",v1.DeleteCate)
        
        //Article Router
        auth.POST("article/add",v1.AddArt)
        auth.PUT("article/:id",v1.EditArt)
        auth.DELETE("article/:id",v1.DeleteArt)
        
        //upload img
        auth.POST("upload",v1.UpLoad)
    }
    router := r.Group("api/v1")
    {
        router.POST("user/add",v1.AddUser)
        router.GET("users",v1.GetUsers)
        router.GET("categorys",v1.GetCates)
        router.GET("articles",v1.GetArts)
        router.GET("articles-category/:id",v1.GetCateArt)
        router.GET("article/info/:id",v1.GetArtInfo)
        router.POST("login",v1.Login)
    }

    r.Run(utils.HttpPort)
}
