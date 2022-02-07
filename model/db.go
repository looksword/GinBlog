package model

import (
    "fmt"
    "GinBlog/utils"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "time"
)

var db *gorm.DB
var err error

func InitDb() {
    db,err = gorm.Open(utils.Db,fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
        utils.DbUser,
        utils.DbPassWord,
        utils.DbHost,
        utils.DbPort,
        utils.DbName,
        ))
    if err != nil {
        fmt.Printf("Connect Mysql Server failed! ",err)
    }

    db.SingularTable(true)
    db.AutoMigrate(&User{},&Article{},&Category{})
    db.DB().SetMaxIdleConns(10)
    db.DB().SetMaxOpenConns(100)
    db.DB().SetConnMaxLifetime(10 * time.Second)
    //db.Close()
}

