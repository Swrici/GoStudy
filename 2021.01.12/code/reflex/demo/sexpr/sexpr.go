package main

import (
	"bytes"
	"fmt"
	"reflect"
)

//展示不同数据类型
func encode(buf *bytes.Buffer, v reflect.Value) error {
	//获取底层类型
	switch v.Kind() {
		//找不到的
		case reflect.Invalid:
			buf.WriteString("nil")
		//int
		case reflect.Int, reflect.Int8, reflect.Int16,
			reflect.Int32, reflect.Int64:
			fmt.Fprintf(buf, "%d", v.Int())
		//uint
		case reflect.Uint, reflect.Uint8, reflect.Uint16,
			reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			fmt.Fprintf(buf, "%d", v.Uint())
		//string
		case reflect.String:
			fmt.Fprintf(buf, "%q", v.String())
		//指针
		case reflect.Ptr:
			return encode(buf, v.Elem())
		//数组或者是切片
		case reflect.Array, reflect.Slice: // (value ...)
			buf.WriteByte('(')
			for i := 0; i < v.Len(); i++ {
				if i > 0 {
					buf.WriteByte(' ')
				}
				if err := encode(buf, v.Index(i)); err != nil {
					return err
				}
			}
			buf.WriteByte(')')
		//结构体
		case reflect.Struct: // ((name value) ...)
			buf.WriteByte('(')
			for i := 0; i < v.NumField(); i++ {
				if i > 0 {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
				if err := encode(buf, v.Field(i)); err != nil {
					return err
				}
				buf.WriteByte(')')
			}
			buf.WriteByte(')')
		//map
		case reflect.Map: // ((key value) ...)
			buf.WriteByte('(')
			for i, key := range v.MapKeys() {
				if i > 0 {
					buf.WriteByte(' ')
				}
				buf.WriteByte('(')
				if err := encode(buf, key); err != nil {
					return err
				}
				buf.WriteByte(' ')
				if err := encode(buf, v.MapIndex(key)); err != nil {
					return err
				}
				buf.WriteByte(')')
			}
			buf.WriteByte(')')
		//更复杂的
		default: // float, complex, bool, chan, func, interface
			return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		fmt.Println("nil")
		return nil, err
	}
	fmt.Println("noNil")
	return buf.Bytes(), nil
}

type Movie struct {
	Title, Subtitle string
	Year            int
	//Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}
func main() {

	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		//Color:    false,
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
	msg,err :=Marshal(strangelove)
	//fmt.Println(Marshal(strangelove))
	if err!=nil {
		for i,b :=range msg {
			fmt.Printf("%c",b)
			fmt.Println(i)
		}
		fmt.Println(msg)
	}
	if msg == nil {
		fmt.Println("msg是nil")
	}
}
