package plugin

import "reflect"

var filterNames = make(map[string]reflect.Type)

func RegisterFilter(name string, rType reflect.Type) {
	filterNames[name] = rType
}

type FilterFactory struct{}

func (f *FilterFactory) Create(conf Config) Plugin {
	t, _ := filterNames[conf.Name]
	return reflect.New(t).Interface().(Plugin)
}
