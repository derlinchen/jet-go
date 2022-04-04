package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"jet/global"
	"strconv"
	"time"
)

var (
	Link *gorm.DB
)

func SetupDBLink() error {
	var err error
	//dsn := "root:password@tcp(127.0.0.1:3306)/business?charset=utf8&parseTime=True&loc=Local"
	Link, err = gorm.Open(mysql.Open(global.DatabaseSetting.DataSourceName), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	//DBLink.Logger.LogMode(true)
	if err != nil {
		return err
	}

	sqlDB, err := Link.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(30)
	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	return nil
}


func Paginate(page int,pageSize int) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func CalcPageCount(total int64, pageSize int) int64 {
	if total > 0 {
		pageSizeStr := strconv.Itoa(pageSize)
		pageSizeVal, _ := strconv.ParseInt(pageSizeStr, 10, 64)
		if total %  pageSizeVal == 0 {
			return total / pageSizeVal
		} else {
			return total / pageSizeVal + 1
		}
	}
	return 0
}