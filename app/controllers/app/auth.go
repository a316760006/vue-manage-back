package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"vue-manage-back/app/common/request"
	"vue-manage-back/app/common/response"
	"vue-manage-back/app/services"
	"vue-manage-back/global"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// Register 用户注册
func Register(c *gin.Context) {
	var form request.Register
	// ShouldBindJSON 解析json
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Register(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}

// 登录
func Login(c *gin.Context) {
	var form request.Login
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Login(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		tokenData, err, _ := services.JwtService.CreateToken(services.AppGuardName, user)
		if err != nil {
			fmt.Printf("err3: %v \n", err)
			response.BusinessFail(c, err.Error())
			return
		}
		response.Success(c, tokenData)
	}
}

// JWTAuth 中间件校验 Token 识别的用户 ID 来获取用户信息
func Info(c *gin.Context) {
	err, user := services.UserService.GetUserInfo(c.Keys["id"].(string))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, user)
}

// 登出接口
func Logout(c *gin.Context) {
	err := services.JwtService.JoinBlackList(c.Keys["token"].(*jwt.Token))
	if err != nil {
		response.BusinessFail(c, "登出失败")
		return
	}
	response.Success(c, nil)
}

var upgrader = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// websocket
func Ws(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		global.App.Log.Error("websocket connection failed", zap.Any("err", err))
		return
	}
	defer ws.Close()
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("read:", err)
			break
		}
		fmt.Printf("recv: %s", message)
		// 每5秒返回一个心跳响应
		ticker := time.NewTicker(5 * time.Second)
		for _ = range ticker.C {
			//写入ws数据

			data := make(map[string]interface{})
			data["code"] = 0
			data["data"] = "{aaa:1,bbb:'心跳响应'}"
			data["message"] = "ok"
			msg, _ := json.Marshal(data)
			err = ws.WriteMessage(mt, msg)
		}
		if err != nil {
			fmt.Println("write:", err)
			break
		}
	}
}
