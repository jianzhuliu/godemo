package plugin

import (
	"reflect"
	"strings"
)

type UpperFilter struct{}

func (u *UpperFilter) Process(msg string) string {
	return strings.ToUpper(msg)
}

func init() {
	RegisterFilter("upper", reflect.TypeOf(UpperFilter{}))
}
