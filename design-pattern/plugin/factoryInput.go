package plugin

import "reflect"

var inputNames = make(map[string]reflect.Type)

func RegisterInput(name string, rType reflect.Type) {
	inputNames[name] = rType
}

type InputFactory struct{}

func (f *InputFactory) Create(conf Config) Plugin {
	t, _ := inputNames[conf.Name]
	return reflect.New(t).Interface().(Plugin)
}
