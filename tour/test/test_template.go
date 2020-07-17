package main

import (
	"os"
	"strings"
	"text/template"
)

const templateText = `
Output 0: {{title .Name1}}
OutPut 1: {{title .Name2}}
Output 2: {{.Name3 | title}}
`

func main() {
	funcMap := template.FuncMap{"title": strings.ToUpper}
	tpl, _ := template.New("go-programming-tour").Funcs(funcMap).Parse(templateText)
	data := map[string]string{
		"Name1": "go",
		"Name2": "programming",
		"Name3": "tour",
	}
	_ = tpl.Execute(os.Stdout, data)
}
