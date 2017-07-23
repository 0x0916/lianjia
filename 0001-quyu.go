package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://bj.lianjia.com/xiaoqu/"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div.position").Find("a").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		fmt.Printf("https://bj.lianjia.com/xiaoqu/%s\n", url)
	})
}
