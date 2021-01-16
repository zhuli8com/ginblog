package model

import (
	"fmt"
	. "ginblog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var db *gorm.DB
var err error

func InitDb()  {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DbUser,
		DbPassword,
		DbHost,
		DbPort,
		DbName)
	fmt.Println(dsn)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),// gorm日志模式：silent
		DisableForeignKeyConstraintWhenMigrating: true,// 外键约束
		SkipDefaultTransaction: true,// 禁用默认事务（提高运行速度）
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
		},
	})
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：",err)
	}

	db.AutoMigrate(&User{},&Article{},&Category{})

	//连接池
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetConnMaxLifetime(time.Hour)// SetConnMaxLifetime 设置了连接可复用的最大时间。
}
