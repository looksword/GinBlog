package model

import (
    "github.com/jinzhu/gorm"
    "GinBlog/utils/errmsg"
    "golang.org/x/crypto/scrypt"
    "log"
    "encoding/base64"
)

type User struct {
    gorm.Model
    Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"UserName"`
    Password string `gorm:"type:varchar(20);not null " json:"password" validate:"required,min=6,max=20" label:"Password"`
    Role int `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"Permission"`
}

// CheckUser
func CheckUser(name string)int {
    var users User
    db.Select("id").Where("username = ?",name).First(&users)
    if users.ID > 0 {
        return errmsg.ERROR_USERNAME_USED
    }
    return errmsg.SUCCESS
}

// AddUser
func CreateUser(data *User)int {
    //data.Password = ScryptPw(data.Password)
    err := db.Create(&data).Error
    if err != nil {
        return errmsg.ERROR
    }
    return errmsg.SUCCESS
}

// query users
func GetUsers(pageSize int,pageNum int) ([]User,int) {
    var users []User
    var total int
    err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total).Error
    if err != nil && err != gorm.ErrRecordNotFound {
        return nil,0
    }
    return users,total
}

// edit user
func EditUser(id int,data *User)int {
    //
    var user User
    var maps = make(map[string]interface{})
    maps["username"] = data.Username
    maps["password"] = data.Role
    err := db.Model(&user).Where("id = ?",id).Updates(maps).Error
    if err != nil {
        return errmsg.ERROR
    }
    return errmsg.SUCCESS
}

// delete user
func DeleteUser(id int)int {
    //
    var user User
    err := db.Where("id = ?",id).Delete(&user).Error
    if err != nil {
        return errmsg.ERROR
    }
    return errmsg.SUCCESS
}

// password encryption
func(u *User)BeforeSave() {
    u.Password = ScryptPw(u.Password)
}

func ScryptPw(password string)string {
    //
    const KeyLen = 10
    salt := make([]byte,8)
    salt = []byte{112,24,37,46,21,98,23,123}
    dk,err := scrypt.Key([]byte(password),salt,1<<14,8,1,KeyLen)
    if err != nil {
        log.Fatal(err)
    }
    FinalPw := base64.StdEncoding.EncodeToString(dk)
    return FinalPw
}

// login verify
func CheckLogin(username string,password string) int {
    var user User
    db.Where("username = ?",username).First(&user)
    if user.ID == 0 {
        return errmsg.ERROR_USER_NOT_EXIST
    }
    if ScryptPw(password) != user.Password {
        return errmsg.ERROR_PASSWORD_WRONG
    }
    if user.Role != 1 {
        return errmsg.ERROR_USER_NO_RIGHT
    }
    return errmsg.SUCCESS
}

