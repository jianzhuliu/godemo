package plugin

import "reflect"

var outputNames = make(map[string]reflect.Type)

func RegisterOutput(name string, rType reflect.Type) {
	outputNames[name] = rType
}

type OutputFactory struct{}

func (f *OutputFactory) Create(conf Config) Plugin {
	t, _ := outputNames[conf.Name]
	return reflect.New(t).Interface().(Plugin)
}
