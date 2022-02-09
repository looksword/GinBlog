package model

import (
    "github.com/jinzhu/gorm"
    "GinBlog/utils/errmsg"
)

type Category struct {
    ID uint `gorm:"primary_key;auto_increment" json:"id"`
    Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// CheckCategory
func CheckCategory(name string)int {
    var cate Category
    db.Select("id").Where("name = ?",name).First(&cate)
    if cate.ID > 0 {
        return errmsg.ERROR_CATENAME_USED
    }
    return errmsg.SUCCESS
}

// AddCategory
func CreateCate(data *Category)int {
    //data.Password = ScryptPw(data.Password)
    err := db.Create(&data).Error
    if err != nil {
        return errmsg.ERROR
    }
    return errmsg.SUCCESS
}

//query article from category

// query categorys
func GetCates(pageSize int,pageNum int) []Category {
    var cates []Category
    err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates).Error
    if err != nil && err != gorm.ErrRecordNotFound {
        return nil
    }
    return cates
}

// edit category
func EditCate(id int,data *Category)int {
    //
    var cate Category
    var maps = make(map[string]interface{})
    maps["name"] = data.Name
    err := db.Model(&cate).Where("id = ?",id).Updates(maps).Error
    if err != nil {
        return errmsg.ERROR_CATENAME_USED
    }
    return errmsg.SUCCESS
}

// delete category
func DeleteCate(id int)int {
    //
    var cate Category
    err := db.Where("id = ?",id).Delete(&cate).Error
    if err != nil {
        return errmsg.ERROR
    }
    return errmsg.SUCCESS
}


