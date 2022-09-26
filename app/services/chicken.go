package services

import (
	"errors"
	"fmt"
	"vue-manage-back/app/common/request"
	"vue-manage-back/app/models"
	"vue-manage-back/global"
	"vue-manage-back/utils"
)

type chicken_soupService struct {
}

var ChickenService = new(chicken_soupService)

// AddChickenSoup 添加鸡汤
func (chickenService *chicken_soupService) AddChickenSoup(params request.ChickenSoup) (err error, chicken models.Chicken_soup) {
	var result = global.App.DB.Where("text = ?", params.Text).Select("id").First(&models.Chicken_soup{})
	if result.RowsAffected != 0 {
		err = errors.New("该鸡汤已存在")
		return
	}
	chicken = models.Chicken_soup{Text: params.Text, Type: params.TypeId}
	err = global.App.DB.Create(&chicken).Error
	return
}

// GetChickenSoup每日一语
func (chickenService *chicken_soupService) GetChickenSoup(params string) (err error, chicken models.Chicken_soup) {
	fmt.Printf("params:%v", params)
	// 查询数据库数据总数
	var max int64
	_ = global.App.DB.Model(&models.Chicken_soup{}).Count(&max)
	var randomNum int
	if params == "0" {
		randomNum = utils.Random(int(max))
	} else {
		randomNum = utils.RandomNum(max)
	}
	result := new(models.Chicken_soup)
	var sqlResult = global.App.DB.First(result, randomNum)
	if sqlResult.RowsAffected >= 0 {
		chicken = models.Chicken_soup{ID: result.ID, Text: result.Text, Type: result.Type}
		return
	}
	err = errors.New("查询失败")
	return
}
