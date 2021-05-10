package utils

import (
	"strings"
)

//高效连接字符串
func JoinStrings(args ...string) string {
	var buf strings.Builder
	buf.Grow(len(args))
	for i := 0; i < len(args); i++ {
		buf.WriteString(args[i])
	}
	return buf.String()
}
