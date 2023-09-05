package request

import (
	"gdback/config"
	"gdback/pkg/jwt"
	"gdback/pkg/logger"
	"gdback/src/mysqldb"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	// return msg fileds
	STATE   = "state"
	MESSAGE = "message"

	//STATE
	S100     = 100
	S100_MSG = "success"
	S101     = 101
	S101_MSG = "failed"
	S102     = 102
	S102_MSG = "parameter error"

	// GET
	PING             = "/ping"
	GameDataRegister = "/gamedata/register"

	// POST
	POST          = "/post"
	EncryptAesCbc = "/encryptoAesCbc"
	DecryptAesCbc = "/decryptoAesCbc"
	Encodejwt     = "/encodejwt"
	Decodejwt     = "/decodejwt"

	UserLogin    = "/user/login"
	UserRegister = "/user/register"

	set_account = "account"
)

var (
	// limit
	LimitApi           = map[string]byte{EncryptAesCbc: 1, DecryptAesCbc: 1, Encodejwt: 1, Decodejwt: 1, UserRegister: 1}
	UnAuthorizationApi = map[string]byte{PING: 1, POST: 1, UserLogin: 1, UserRegister: 1, EncryptAesCbc: 1, DecryptAesCbc: 1, Encodejwt: 1, Decodejwt: 1}

	MSG100 = gin.H{
		"state":   S100,
		"message": S100_MSG,
	}

	MSG101 = gin.H{
		"state":   S101,
		"message": S101_MSG,
	}

	MSG102 = gin.H{
		"state":   S102,
		"message": S102_MSG,
	}
)

func request(req *gin.Engine) {
	// // 添加中间件函数来禁用某些路由
	// req.Use(func(c *gin.Context) {
	// 	// 在中间件函数中判断是否禁用路由
	// 	if shouldDisableRoute(c) {
	// 		logger.Info("request bin used, req:%v", c.FullPath())
	// 		// 如果需要禁用路由，中止请求处理
	// 		c.AbortWithStatus(http.StatusForbidden)
	// 		return
	// 	}
	// 	// 如果不需要禁用路由，继续处理请求
	// 	c.Next()
	// })
	req.Use(func(c *gin.Context) {
		is, err := shouldDisableRoute(c)
		if !is {
			logger.Info("request can't used, err:%v", err)
			c.JSON(http.StatusOK, retMsg(MSG102, err))
			c.AbortWithStatus(http.StatusForbidden)
		}
		c.Next()
	})

	//get
	req.GET(PING, gethandle_ping)
	req.GET(GameDataRegister, gethandle_gamedata_register)

	//post
	req.POST(POST, posthandle_post)
	req.POST(EncryptAesCbc, posthandle_encryptoAesCbc)
	req.POST(DecryptAesCbc, posthandle_decryptoAesCbc)
	req.POST(UserLogin, posthandle_user_login)
	req.POST(UserRegister, posthandle_user_register)
	req.POST(Encodejwt, posthandle_encodejwt)
	req.POST(Decodejwt, posthandle_decodejwt)
}

func shouldDisableRoute(c *gin.Context) (bool, string) {
	fullpath := c.FullPath()
	checkLimitApi, errapi := checkLimitApi(fullpath)
	if !checkLimitApi {
		return checkLimitApi, errapi
	}

	token := c.GetHeader("Authorization")
	checkAuth, errauth, account := checkAuth(fullpath, token)
	if !checkAuth {
		return checkAuth, errauth
	}
	c.Set(set_account, account)
	return true, ""
}

func checkLimitApi(fullpath string) (bool, string) {
	serverType := config.Config.GetString("server_type")
	if serverType == "dev" {
		return true, "isLimitApi"
	} else {
		return LimitApi[fullpath] != 1, "isLimitApi"
	}
}

func checkAuth(fullpath, token string) (bool, string, string) {
	retmsg := ""
	account := ""
	if UnAuthorizationApi[fullpath] == 1 {
		return true, retmsg, account
	}
	mapinfo, err := jwt.DecodeJwt(token)
	if err != nil {
		return false, err.Error(), ""
	}
	info := mapinfo["info"].(map[string]any)
	if info["account"] == nil {
		return false, "Authorization error", account
	}
	account = info["account"].(string)
	if !mysqldb.UserIsExist(account) {
		return false, "account not exist", account
	}
	return true, retmsg, account
}
