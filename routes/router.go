package routes

import (
    "GinBlog/utils"
    "github.com/gin-gonic/gin"
    "net/http"
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

        //Article Router

        //Login Router
    }

    r.Run(utils.HttpPort)
}
