package request

import (
	"encoding/base64"
	"fmt"
	"gdback/config"
	"gdback/pkg/crypto"
	"gdback/pkg/jwt"
	"gdback/src/mysqldb"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Start() {
	req := gin.Default()
	request(req)
	ipport := config.Config.GetString("host.ip") + ":" + config.Config.GetString("host.port")
	req.Run(ipport) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func retMsg(ret gin.H, format string, a ...any) gin.H {
	msgstr := fmt.Sprintf(format, a...)
	ret[MESSAGE] = msgstr
	return ret
}

func retData(ret gin.H, data any) gin.H {
	ret[MESSAGE] = data
	return ret
}

// get
func gethandle_ping(c *gin.Context) {
	c.JSON(http.StatusOK, retMsg(MSG100, "pong"))
}

func gethandle_gamedata_register(c *gin.Context) {
	info, err := mysqldb.RegisterInfo()
	ret := MSG100
	if err != nil {
		ret = retMsg(MSG100, err.Error())
	}
	ret[MESSAGE] = string(info)
	c.JSON(http.StatusOK, ret)
}

// post
func posthandle_post(c *gin.Context) {
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

func posthandle_encryptoAesCbc(c *gin.Context) {
	info := c.PostForm("info")
	if info == "" {
		c.JSON(http.StatusOK, MSG102)
		return
	}
	infobyte := []byte(info)
	infostr, err := crypto.EncryptoAesCbc(infobyte)
	if err != nil {
		c.JSON(http.StatusOK, retMsg(MSG100, err.Error()))
		return
	}
	// infostr2 := base64.StdEncoding.EncodeToString(infostr)
	retdata := gin.H{
		"info": infostr,
	}
	c.JSON(http.StatusOK, retData(MSG100, retdata))
}

func posthandle_decryptoAesCbc(c *gin.Context) {
	info := c.PostForm("info")
	if info == "" {
		c.JSON(http.StatusOK, MSG102)
		return
	}
	infobyte, _ := base64.StdEncoding.DecodeString(info)
	infostr, err := crypto.DecryptoAesCbc(infobyte)
	if err != nil {
		c.JSON(http.StatusOK, retMsg(MSG100, err.Error()))
		return
	}
	// fmt.Println("解密后的数据 t1：", string(infostr))
	retdata := gin.H{
		"info": string(infostr),
	}
	c.JSON(http.StatusOK, retData(MSG100, retdata))
}

func posthandle_encodejwt(c *gin.Context) {
	info := c.PostForm("info")
	if info == "" {
		c.JSON(http.StatusOK, MSG102)
		return
	}
	mapinfo := map[string]any{"token": info}
	token, err := jwt.EncodeJwt(mapinfo)
	if err != nil {
		c.JSON(http.StatusOK, retMsg(MSG100, err.Error()))
		return
	}
	c.JSON(http.StatusOK, retData(MSG100, token))
}

func posthandle_decodejwt(c *gin.Context) {
	info := c.PostForm("info")
	if info == "" {
		c.JSON(http.StatusOK, MSG102)
		return
	}
	tokeninfo, err := jwt.DecodeJwt(info)
	if err != nil {
		c.JSON(http.StatusOK, retMsg(MSG100, err.Error()))
		return
	}
	c.JSON(http.StatusOK, retData(MSG100, tokeninfo))
}
