package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer = os.Stdout
	f, ok := w.(*os.File)      // 成功:  ok, f == os.Stdout
	b, ok := w.(*bytes.Buffer) // 识别: !ok, b == nil
	fmt.Println(f,b,ok)
}
