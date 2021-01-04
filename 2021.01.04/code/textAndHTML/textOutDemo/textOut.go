//例子分析 将上一个json中查询结果按templ模板输出
package main

import (
	"com/swrici/dataType/complexType/json/excercise/excercise4.10/github"
	"log"
	"os"
	"text/template"
	"time"
)

//printf是内置函数
//daysAgo是自写的函数
const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

//用于输出时间格式
func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))



func main()  {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
