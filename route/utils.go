package route

import (
	"errors"
	"reflect"
)

// Handler 功能分发
func Handler(args []string) interface{} {
	var target = args[1]
	var _, found = Setting.handlerMap[args[1]]
	if !found {
		return "[未知的命令]"
	}
	args = shift(args)
	ret, err := call(Setting.handlerMap, target, args)
	if err != nil {
		return err.Error()
	}
	return ret
}

// 函数反射
func call(m map[string]interface{}, name string, params []string) ([]reflect.Value, error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		return nil, errors.New("[参数错误]")
	}
	in := make([]reflect.Value, len(params))
	for k, v := range params {
		in[k] = reflect.ValueOf(v)
	}
	return f.Call(in), nil
}

// 删除数组开头元素
func shift(arr []string) []string {
	var res []string
	for index, one := range arr {
		if index != 0 && index != 1 {
			res = append(res, one)
		}
	}
	return res
}
