package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := os.Args[1]
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("ul.listContent").Find("li").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Find("div.title").Find("a").Attr("href")
		fmt.Println(url)
	})
}
