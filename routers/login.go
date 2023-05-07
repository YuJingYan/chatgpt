package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 登录数据的结构体（确保和数据库中的结构一致）
type Login struct {
	Id       int    `json:"id" form:"id"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

// 登录接口
func LoginHandler(c *gin.Context) {
	var loginForm, _ Login
	if err := c.ShouldBindJSON(&loginForm); err != nil {
		// 返回错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg":   "登录成功",
		"token": "123jkasdqwe1231a12r13",
	})
}
