package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode string
	HttpPort string

	DbHost string
	DbPort string
	DbUser string
	DbPassword string
	DbName string
)

func init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Print("配置文件读取错误，请检查文件路径：",err)
	}

	loadServer(file)
	loadData(file)
}

func loadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
}

func loadData(file *ini.File) {
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("zhuli")
	DbPassword = file.Section("database").Key("DbPassword").MustString("zhuli")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}