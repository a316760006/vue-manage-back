package utils

import (
	"math/rand"
	"time"
)

// 拿今天日期来当随机数种子,随机出当天的随机数
func Random(max int) int {
	now := time.Now()
	nowStr := now.Format("2006-01-02")
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation("2006-01-02", nowStr, loc)
	rand.Seed(theTime.Unix())
	num := rand.Intn(max)
	return int(num)
}

// 随机整数
func RandomNum(max int64) int {
	rand.Seed(time.Now().UnixNano()) // 纳秒时间戳
	data := rand.Int63n(max)
	return int(data)
}
