package handler

import (
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetSessionByKey(c *gin.Context, key string) interface{} {
	session := sessions.Default(c)
	return session.Get(key)
}

func SetSessionKey(c *gin.Context, key string, value interface{}) {
	session := sessions.Default(c)

	session.Set(key, value)

	err := session.Save()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func DelSessionKey(c *gin.Context, key string) {
	session := sessions.Default(c)

	session.Delete(key)

	err := session.Save()
	if err != nil {
		fmt.Println(err.Error())
	}
}
