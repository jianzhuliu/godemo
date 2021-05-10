package utils

import (
	"os"
)

//验证文件或者目录是否存在
func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
