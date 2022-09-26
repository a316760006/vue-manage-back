package models

type Chicken_soup struct {
	ID
	Text string `json:"text" gorm:"not null;comment:毒鸡汤文案"`
	Type string `json:"type" gorm:"not null;index;comment:毒鸡汤类型"`
}

// // 后续其他的用户模型都可以通过实现 JwtUser 接口
// func (chicken Chicken_soup) GetChicken() string {
// 	return strconv.Itoa(int(chicken.ID.ID))
// }
