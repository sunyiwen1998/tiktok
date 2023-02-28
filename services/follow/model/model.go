package model

import (
	"log"
	"time"

	"github.com/ozline/tiktok/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

// Setup 初始化数据库
func Setup() {
	var (
		err         error
		tablePrefix string
	)

	// open the database and buffer the config
	db, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix, // set the prefix name of table
			SingularTable: true,        // use singular table by default
		},
		// gorm日志模式：silent
		Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	mysqlDB, err := db.DB()
	if err != nil {
		log.Panicln("db.DB() err: ", err)
	}

	// some init set of database
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	mysqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	mysqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifeTime 设置连接的最大可复用时间。
	mysqlDB.SetConnMaxLifetime(10 * time.Second)
}
