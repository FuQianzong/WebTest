package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"mvctest/common"
	"mvctest/model"
	"mvctest/util"
)

//注册
func Register(c *gin.Context) {
	DB:=common.GetDB()
	//获取参数
	name:=c.PostForm("name")
	telephone:=c.PostForm("telephone")
	password:=c.PostForm("password")
	//log.Println(name,telephone,password)
	//数据认证
	if len(telephone)!=11{
		c.JSON(422,gin.H{"msg":"电话号码必须为11位"})
		return
	}
	if len(password)<6{
		c.JSON(422,gin.H{"msg":"密码不能小于6位"})
		return
	}
	if len(name)==0{
		name=util.RandString(10)
	}

	//判断是否存在用户
	if TelephoneExit(DB,telephone){
		c.JSON(422,gin.H{"msg":"该手机号已注册"})
		return
	}
	//加密密码
	hashedPassword,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err!=nil{
		c.JSON(500,gin.H{"msg":"加密错误"})
		return
	}
	//创建用户
	newUser:=model.User{
		Name: name,
		Telephone: telephone,
		Password: string(hashedPassword),
	}
	DB.Create(&newUser)
	//返回结果
	c.JSON(200, gin.H{
		"message": "注册成功",
	})
}
//登录
func Login(c *gin.Context){
	DB:=common.GetDB()
	//获取参数
	telephone:=c.PostForm("telephone")
	password:=c.PostForm("password")
	//数据验证
	if len(telephone)!=11{
		c.JSON(422,gin.H{"msg":"电话号码必须为11位"})
		return
	}
	if len(password)<6{
		c.JSON(422,gin.H{"msg":"密码不能小于6位"})
		return
	}
	//判断手机号是否存在
	var user model.User
	DB.Where("telephone=?",telephone).First(&user)
	if user.ID==0{
		c.JSON(422,gin.H{"msg":"手机号不存在"})
		return
	}
	//判断密码是否正确
	if err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password));err!=nil{
		c.JSON(400,gin.H{"msg":"密码错误"})
		return
	}
	//发放token
	token,err:=common.ReleaseToken(user)
	if err!=nil{
		c.JSON(500,gin.H{"msg":"系统错误"})
		log.Println("token generate error")
		return
	}
	//返回结果
	c.JSON(200, gin.H{
		"token":token,
		"message": "登录成功",
	})
}
//判断手机号是否存在
func TelephoneExit(db *gorm.DB,telephone string) bool{
	var user model.User
	db.Where("telephone=?",telephone).First(&user)
	//fmt.Println(telephone,user.ID)
	if user.ID!=0{
		return true
	}
	return false
}