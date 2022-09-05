package app

import (
	"fmt"
	"vue-manage-back/app/common/request"
	"vue-manage-back/app/common/response"
	"vue-manage-back/app/services"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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
