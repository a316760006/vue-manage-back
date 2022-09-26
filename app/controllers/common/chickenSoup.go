package common

import (
	"vue-manage-back/app/common/request"
	"vue-manage-back/app/common/response"
	"vue-manage-back/app/services"

	"github.com/gin-gonic/gin"
)

//添加鸡汤
func AddChickenSoup(c *gin.Context) {
	var form request.ChickenSoup
	// ShouldBindJSON 解析json
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, chickenSoup := services.ChickenService.AddChickenSoup(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		// global.App.Redis.Set(chickenSoup.id, "chickenSoupLen")
		response.Success(c, chickenSoup)
	}
}

// 每日一语
func GetChickenSoup(c *gin.Context) {
	form := c.DefaultQuery("type", "0")
	if err, chickenSoup := services.ChickenService.GetChickenSoup(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		// global.App.Redis.Set(chickenSoup.id, "chickenSoupLen")
		response.Success(c, chickenSoup)
	}
}
