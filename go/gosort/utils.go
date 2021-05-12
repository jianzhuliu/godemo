package gosort

import (
	"math/rand"
	"time"
)

//获取随机长度随机数值int切片
func GetRandIntSlice(maxCnt int) []int {
	rand.Seed(time.Now().UnixNano())
	randLen := rand.Intn(maxCnt) + 1
	res := make([]int, randLen)

	for i := 0; i < randLen; i++ {
		res[i] = rand.Intn(maxCnt)
	}

	return res
}
