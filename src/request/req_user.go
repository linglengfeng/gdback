package request

import (
	"gdback/config"
	"gdback/pkg/jwt"
	"gdback/pkg/logger"
	"gdback/pkg/myutil"
	"gdback/src/mysqldb"
	"net/http"

	"github.com/gin-gonic/gin"
)

func posthandle_user_login(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	if account == "" || password == "" {
		c.JSON(http.StatusOK, MSG102)
		return
	}
	user := config.Config.GetStringSlice("register_user")
	ismem := myutil.IsMember[string](account, user)
	if !ismem {
		c.JSON(http.StatusOK, retMsg(MSG101, "account error"))
		return
	}
	islogin := mysqldb.UserLogin(account, password)
	if !islogin {
		c.JSON(http.StatusOK, retMsg(MSG101, "password error"))
		return
	}
	mapinfo := map[string]any{"account": account}
	token, err := jwt.EncodeJwt(mapinfo)
	if err != nil {
		c.JSON(http.StatusOK, retMsg(MSG100, err.Error()))
		return
	}
	logger.Info("posthandle_user_login, account:%v, password:%v, canlogin:%v", account, password, islogin)
	c.JSON(http.StatusOK, retData(MSG100, token))
}

func posthandle_user_register(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	if account == "" || password == "" {
		c.JSON(http.StatusOK, MSG102)
		return
	}
	user := config.Config.GetStringSlice("register_user")
	ismem := myutil.IsMember[string](account, user)
	if !ismem {
		c.JSON(http.StatusOK, retMsg(MSG101, "account not register"))
		return
	}
	mysqldb.UserInsert(account, password)
	logger.Info("posthandle_user_register, account:%v, password:%v", account, password)
	c.JSON(http.StatusOK, MSG100)
}
