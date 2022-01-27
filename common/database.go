package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"mvctest/model"
)

var DB *gorm.DB

//数据库连接
func InitDB() *gorm.DB{
	driverName:="mysql"
	host:="localhost"
	port:="3306"
	database:="mvctest"
	username:="root"
	password:="root"
	charset:="utf8"
	//root:root@tcp(localhost:3306)/mvctest?charset=utf8&parseTime=true
	args:=fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db,err:=gorm.Open(driverName,args)
	if err!=nil{
		panic("数据库连接失败，错误原因:"+err.Error())
	}
	db.AutoMigrate(&model.User{})
	DB=db
	return db
}

func GetDB() *gorm.DB{
	return DB
}