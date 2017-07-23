package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := os.Args[1]
	if url == "I" {
		fmt.Fprintf(os.Stderr, "skip url: %s\n", os.Args[2])
		os.Exit(0)
	}
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		log.Fatal(err)
	}
	name := doc.Find("h1").Text()
	hushuString := doc.Find("div.xiaoquInfo>div:nth-child(7) .xiaoquInfoContent").Text()
	hushu := strings.Replace(hushuString, "户", "", -1)

	fmt.Printf("%s %s\n", name, hushu)
}
