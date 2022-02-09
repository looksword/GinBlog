package model

import (
    "github.com/jinzhu/gorm"
    "GinBlog/utils/errmsg"
)


type Article struct {
    Category Category
    gorm.Model
    Title string `gorm:"type:varchar(100);not null" json:"title"`
    Cid int `gorm:"type:int;not null" json:"cid"`
    Desc string `gorm:"type:varchar(200)" json:"desc"`
    Context string `gorm:"type:longtext" json:"context"`
    Img string `gorm:"type:varchar(100)" json:"img"`
}

// Add article
func CreateArt(data *Article)int {
    //data.Password = ScryptPw(data.Password)
    err := db.Create(&data).Error
    if err != nil {
        return errmsg.ERROR
    }
    return errmsg.SUCCESS
}

//query article from category

//query single article

// query categorys
func GetArts(pageSize int,pageNum int) []Category {
    var cates []Category
    err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates).Error
    if err != nil && err != gorm.ErrRecordNotFound {
        return nil
    }
    return cates
}

// edit Article
func EditArt(id int,data *Article)int {
    //
    var art Article
    var maps = make(map[string]interface{})
    maps["title"] = data.Title
    maps["cid"] = data.Cid
    maps["desc"] = data.Desc
    maps["context"] = data.Context
    maps["img"] = data.Img
    err := db.Model(&art).Where("id = ?",id).Updates(maps).Error
    if err != nil {
        return errmsg.ERROR
    }
    return errmsg.SUCCESS
}

// delete category
func DeleteArt(id int)int {
    //
    var art Article
    err := db.Where("id = ?",id).Delete(&art).Error
    if err != nil {
        return errmsg.ERROR
    }
    return errmsg.SUCCESS
}

