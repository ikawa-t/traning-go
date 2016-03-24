package main

import (
	//"fmt"
	"html/template"
	"log"
	"os"
	"time"

	"ch04_ex14/github"
)

//実行コマンドの引数 repo:golang/go is:open json decoder

const templ = `{{.TotalCount}} issues:
{{range .Items}}-----------------------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end]}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {
	os.Setenv("HTTP_PROXY", "http://p000480075:********@proxy.ricoh.co.jp:8080")
	os.Setenv("HTTPS_PROXY", "http://p000480075:********@proxy.ricoh.co.jp:8080")

	result, err := github.SearchIssue(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
