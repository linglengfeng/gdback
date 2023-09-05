package request

import (
	"gdback/pkg/logger"
	"gdback/src/mysqldb"
	"net/http"

	"github.com/gin-gonic/gin"
)

func gethandle_gamedata_register(c *gin.Context) {
	info := mysqldb.RegisterInfo()
	a, b := c.Get(set_account)
	logger.Debug("gethandle_gamedata_register,account:%v, b:%v", a, b)
	c.JSON(http.StatusOK, retData(MSG100, info))
}
