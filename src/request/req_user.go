package request

import (
	"fmt"
	"gdback/config"
	"gdback/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func posthandle_user_login(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	fmt.Println("id2, name:", id, name)
	if id == "" || name == "" {
		c.JSON(http.StatusOK, MSG102)
		return
	}
	// retmap := make(map[string]any)
	// retmap["id"] = id
	// retmap["name"] = name
	retdata := gin.H{
		"id":   id,
		"name": name,
	}
	c.JSON(http.StatusOK, retData(MSG100, retdata))
}

func posthandle_user_register(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	if account == "" || password == "" {
		c.JSON(http.StatusOK, MSG102)
		return
	}
	user := config.Config.GetString("user")
	fmt.Println("user:", user)
	logger.Info("posthandle_user_register, account:%v, password:%v", account, password)
	c.JSON(http.StatusOK, MSG100)
}
