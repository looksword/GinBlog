package utils

import (
    "fmt"
    "gopkg.in/ini.v1"
)

var(
    AppMode string
    HttpPort string
    Db string
    DbHost string
    DbPort string
    DbUser string
    DbPassWord string
    DbName string
)

func init() {
    file, err := ini.Load("config/config.ini")
    if err != nil {
        fmt.Println("config.ini load false:",err)
    }
    LoadServer(file)
    LoadData(file)
}

func LoadServer(file *ini.File) {
    AppMode = file.Section("server").Key("AppMode").MustString("debug")
    HttpPort = file.Section("server").Key("HttpPort").MustString(":9103")
}

func LoadData(file *ini.File) {
    Db = file.Section("database").Key("Db").MustString("mysql")
    DbHost = file.Section("database").Key("DbHost").MustString("localhost")
    DbPort = file.Section("database").Key("DbPort").MustString("3306")
    DbUser = file.Section("database").Key("DbUser").MustString("ginblog")
    DbPassWord = file.Section("database").Key("DbPassWord").MustString("Admin537.")
    DbName = file.Section("database").Key("DbName").MustString("ginblog")
}
