//我们这里有一个未完成的工作。
//其实map的key的类型并不局限于formatAtom能完美处理的类型；
//数组、结构体和接口都可以作为map的key。
//针对这种类型，完善key的显示信息是练习12.1的任务。
//练习 12.1： 扩展Displayhans，使它可以显示包含以结构体或数组作为map的key类型的值。
package main

import (
"fmt"
"reflect"
"strconv"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}


//应该尽量避免在一个包的API中暴露涉及反射的接口。
//定义一个未导出的display函数用于递归处理工作，导出的是Display函数，
//它只是display函数简单的包装以接受interface{}类型的参数：
func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}


func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			t:=key.Kind()
			if t!=reflect.Struct&&t!=reflect.Array {
				display(fmt.Sprintf("%s[%s]", path,
					formatAtom(key)), v.MapIndex(key))
			}
			if t==reflect.Array {
				for i := 0; i < key.Len(); i++ {
					display(fmt.Sprintf("%s[%d]", key, i), key.Index(i))
				}
				display(fmt.Sprintf("%s[%s]", path,
					formatAtom(key)), v.MapIndex(key))
			}
			if t==reflect.Struct {
				for i := 0; i < key.NumField(); i++ {
					fieldPath := fmt.Sprintf("%s.%s", path, key.Type().Field(i).Name)
					display(fieldPath, v.Field(i))
				}
				display(fmt.Sprintf("%s[%s]", path,
					formatAtom(key)), v.MapIndex(key))
			}
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

func main() {
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},

		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	Display("strangelove",strangelove)
	//Display("os.Stderr", os.Stderr)
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}
