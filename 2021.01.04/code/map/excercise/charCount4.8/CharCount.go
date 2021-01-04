//统计字母、数字等Unicode中不同的字符类别。
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main()  {
	typeUnicode := make(map[string]int)
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for	i:=0;i<10;i++ {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
		fmt.Println("还能输入的字数:",10-i)
		typeSwitch(typeUnicode,r)
		//if r=="end" {
		//
		//}
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
	//输出类型
	fmt.Println("输出练习所要的类型map：")
	fmt.Println(typeUnicode)

}

func typeSwitch(typeUnicode map[string]int, r rune) {
	switch ifType(r) {
	case "number":
		typeUnicode["number"]++
	case "control":
		typeUnicode["control"]++
	case "digit":
		typeUnicode["digit"]++
	case "graphic":
		typeUnicode["graphic"]++
	case "letter":
		typeUnicode["letter"]++
	case "lower":
		typeUnicode["lower"]++
	case "mark":
		typeUnicode["mark"]++
	case "other":
		typeUnicode["other"]++
	}
}

func ifType( r rune) string {
	if unicode.IsNumber(r) {
		return "number"
	}else if unicode.IsControl(r) {
		return "control"
	}else if unicode.IsDigit(r) {
		return "digit"
	}else if unicode.IsGraphic(r) {
		return "graphic"
	}else if unicode.IsLetter(r) {
		return "letter"
	}else if unicode.IsLower(r){
		return "lower"
	}else if unicode.IsMark(r) {
		return "mark"
	}else{
		return "other"
	}
}
