package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"pledge-backend-test/config"
	"pledge-backend-test/log"
	"time"
)

func InitMysql() {
	mysqlConfig := config.Config.Mysql
	//打印日志
	log.Logger.Info("init mysql")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlConfig.UserName,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.DbName)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   //datasource
		DefaultStringSize:         256,   //string类型字段默认长度
		DisableDatetimePrecision:  true,  //禁用datatime精度，mysql5.6之前的数据库不支持
		DontSupportRenameIndex:    true,  //不支持重命名索引，重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  //不支持重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, //根据mysql当前版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //关闭表名的复数形式（默认为false，GORM 会将结构体名称后面加个s作为表名）
		},
		SkipDefaultTransaction: true, //禁用 GORM 的默认事务
	})
	if err != nil {
		//打印日志
		log.Logger.Panic(fmt.Sprintf("mysql connection err ==> %+v", err))
	}

	//在数据库操作后，通过回调函数，打印sql的执行计划
	_ = db.Callback().Create().After("gorm:after_create").Register("after_create", After)
	_ = db.Callback().Update().After("gorm:after_update").Register("after_update", After)
	_ = db.Callback().Query().After("gorm:after_query").Register("after_query", After)
	_ = db.Callback().Delete().After("gorm:after_delete").Register("after_delete", After)
	_ = db.Callback().Raw().After("gorm:raw").Register("raw", After)
	_ = db.Callback().Row().After("gorm:row").Register("row", After)

	sqlDb, err := db.DB()
	if err != nil {
		//todo 打印日志
		log.Logger.Error("db.DB() err:" + err.Error())
	}
	sqlDb.SetConnMaxLifetime(time.Duration(mysqlConfig.MaxLifeTime) * time.Second) //最大连接时间，否则连接超时
	sqlDb.SetMaxIdleConns(mysqlConfig.MaxIdelConns)                                //最大空闲连接数
	sqlDb.SetMaxOpenConns(mysqlConfig.MaxOpenConns)                                //最大连接数

	//全局变量赋值
	Mysql = db
}

func After(db *gorm.DB) {
	db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
}
