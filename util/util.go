package util

import (
	"math/rand"
	"time"
)

//生成随机字符串
func RandString(n int) string{
	letters:=[]byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result:=make([]byte,n)
	rand.Seed(time.Now().Unix())
	for i:=range result{
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
