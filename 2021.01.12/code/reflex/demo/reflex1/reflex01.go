package main

import "strconv"

func Sprint(x interface{}) string {
	//接口，接口方法是一个string()
	type stringer interface {
		String() string
	}
	//判断类型：根据类型返回对应的格式
	switch x := x.(type) {
	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		// array, chan, func, map, pointer, slice, struct 这些
		return "???"
	}
}
