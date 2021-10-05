package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode    string
	HttpPort   string
	DB         string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
)

// init config file
func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Print("load config file error:", err)
	}
	LoadServer(file)
	LoadData(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":9093")

}

func LoadData(file *ini.File) {
	DB = file.Section("database").Key("DB").MustString("mysql")
	DBHost = file.Section("database").Key("DBHost").MustString("127.0.0.1")
	DBPort = file.Section("database").Key("DBPort ").MustString("3306")
	DBUser = file.Section("database").Key("DBUser").MustString("root")
	DBPassword = file.Section("database").Key("DBPassword").MustString("root123")
	DBName = file.Section("database").Key("DBName").MustString("ginblog")
}
