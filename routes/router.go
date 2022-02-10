package routes

import (
    "GinBlog/utils"
    "github.com/gin-gonic/gin"
    "GinBlog/api/v1"
)


func InitRouter() {
    gin.SetMode(utils.AppMode)
    r := gin.Default()
    routerv1 := r.Group("api/v1")
    {
        //User Router
        routerv1.POST("user/add",v1.AddUser)
        routerv1.GET("users",v1.GetUsers)
        routerv1.PUT("user/:id",v1.EditUser)
        routerv1.DELETE("user/:id",v1.DeleteUser)

        //Category Router
        routerv1.POST("category/add",v1.AddCate)
        routerv1.GET("categorys",v1.GetCates)
        routerv1.PUT("category/:id",v1.EditCate)
        routerv1.DELETE("category/:id",v1.DeleteCate)
        
        //Article Router
        routerv1.POST("article/add",v1.AddArt)
        routerv1.GET("articles",v1.GetArts)
        routerv1.GET("articles-category/:id",v1.GetCateArt)
        routerv1.GET("article/info/:id",v1.GetArtInfo)
        routerv1.PUT("article/:id",v1.EditArt)
        routerv1.DELETE("article/:id",v1.DeleteArt)
        
        //Login Router
    }

    r.Run(utils.HttpPort)
}
