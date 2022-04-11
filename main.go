package main

import (
	"jet/bean"
	"jet/config"
	"jet/db"
	"jet/global"
	"jet/routers"
	"log"
	"os"
	"time"
)

func main() {
	r := routers.SetupRouter()
	if err := r.Run(bean.COLON + global.ServerSetting.Port); err != nil {
		log.Fatalf("startup service failed, err:%v\n", err)
	}
}

func init() {
	// 创建日志文件
	createLogFile(createFile)

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

func createLogFile(createFile func()) {
	go func() {
		for {
			createFile()
			now := time.Now()
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}

func createFile() {
	logdir := "log/" + time.Now().Format("20060102") + "/"
	f, err := os.OpenFile(logdir+"jet.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		_, err := os.Stat(logdir)
		if os.IsNotExist(err) {
			os.MkdirAll(logdir, os.ModePerm)
		}
		f, _ = os.Create(logdir + "jet.log")
	}
	log.SetOutput(f)
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}
