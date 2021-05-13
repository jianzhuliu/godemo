package gosort

import (
	"math/rand"
	"time"
)

//获取随机长度随机数值int切片
func GetRandIntSlice(maxCnt int) (res, copyData []int) {
	rand.Seed(time.Now().UnixNano())
	randLen := rand.Intn(maxCnt) + 1
	res = make([]int, randLen)
	copyData = make([]int, randLen)

	for i := 0; i < randLen; i++ {
		res[i] = rand.Intn(maxCnt)
		copyData[i] = res[i]
	}

	return
}

//切片复制
func CopySlice(data []int) (copyData []int) {
	copyData = make([]int, len(data))
	copy(copyData, data)
	return
}
