package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"ch04_ex11/github"
)


// GitHubのWebAPIを用いる
// 未実施


//実行コマンドの引数 repo:golang/go is:open json decoder

func main() {
	os.Setenv("HTTP_PROXY", "http://p000480075:********@proxy.ricoh.co.jp:8080")
	os.Setenv("HTTPS_PROXY", "http://p000480075:********@proxy.ricoh.co.jp:8080")

	result, err := github.SearchIssue(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	var lessThanOneMonth, lessThanOneYear, moreThanOneYear []github.Issue
	now := time.Now()

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		duration := now.Sub(item.CreatedAt)
		// 一ヶ月未満
		if duration.Hours() < 24*30 {
			lessThanOneMonth = append(lessThanOneMonth, *item)
		// 一年未満
		}else if duration.Hours() < 24*365 {
			lessThanOneYear = append(lessThanOneYear, *item)
		// 一年以上
		} else {
			moreThanOneYear = append(moreThanOneYear, *item)
		}
	}

	fmt.Println("lessThanOneMonth")
	for _, item := range lessThanOneMonth {
		printIssue(item)
	}
	fmt.Println("lessThanOneYear")
	for _, item := range lessThanOneYear {
		printIssue(item)
	}
	fmt.Println("moreThanOneYear")
	for _, item := range moreThanOneYear {
		printIssue(item)
	}

}

func printIssue(item github.Issue) {
	fmt.Printf("%v #%-5d %9.9s %.55s\n ",
		item.CreatedAt, item.Number, item.User.Login, item.Title)
}
