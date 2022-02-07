package main

import (
    "GinBlog/routes"
    "GinBlog/model"
)

func main() {
    model.InitDb()
    routes.InitRouter()
}
