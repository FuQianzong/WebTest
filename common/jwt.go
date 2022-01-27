package common

import (
	"github.com/dgrijalva/jwt-go"
	"mvctest/model"
	"time"
)

var jwtKey=[]byte("a_secret_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

//token发放
func ReleaseToken(user model.User)(string,error){
	//token有效时间
	expirationTime:=time.Now().Add(7*24*time.Hour)
	claims:=&Claims{
		UserId: user.ID,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt: time.Now().Unix(),
			Issuer: "oceanlearn.tech",
			Subject: "user token",
		},
	}
	//jwt加密包HS256加密算法
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString,err:=token.SignedString(jwtKey)
	if err!=nil{
		return "", err
	}
	return tokenString,nil
}