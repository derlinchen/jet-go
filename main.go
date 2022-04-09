package main

import (
	"io"
	"jet/bean"
	"jet/config"
	"jet/db"
	"jet/global"
	"jet/routers"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	f, err := os.OpenFile("log/jet.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		_, err := os.Stat("log/")
		if os.IsNotExist(err) {
			os.MkdirAll("log/", os.ModePerm)
		}
		os.Create("log/jet.log")
	}

	gin.DefaultWriter = io.MultiWriter(f)
	r := routers.SetupRouter()
	if err := r.Run(bean.COLON + global.ServerSetting.Port); err != nil {
		log.Fatalf("startup service failed, err:%v\n", err)
	}
}

func init() {
	// 加载配置文件
	err := setupSetting()
	if err != nil {
		log.Fatal(err)
	}

	// 建立数据库连接
	err = db.SetupDBLink()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

func setupSetting() error {
	args := os.Args
	newSetting, err := config.NewSetting(args)
	if err != nil {
		return err
	}
	if err := newSetting.ReadSection("Database", &global.DatabaseSetting); err != nil {
		return err
	}
	global.DatabaseSetting.DataSourceName = global.DatabaseSetting.UserName +
		":" + global.DatabaseSetting.Password + "@tcp(" +
		global.DatabaseSetting.Domain + ":" +
		global.DatabaseSetting.Port + ")/" +
		global.DatabaseSetting.DBName + "?" +
		"charset=" + global.DatabaseSetting.Charset + "&" +
		"loc=" + global.DatabaseSetting.Local + "&" +
		"parseTime=" + global.DatabaseSetting.ParseTime
	if err := newSetting.ReadSection("Server", &global.ServerSetting); err != nil {
		return err
	}
	return nil
}
