package crawler

import (
	"fmt"
	"log"
	"github.com/PuerkitoBio/goquery"
)

func Scrape() {
	url := "https://gupiao.baidu.com/stock/sh601000.html"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".stock-bets").Each(func(i int, s *goquery.Selection) {
		name := s.Find(".bets-name").Find("span").Text()
		price := s.Find("._close").Text()
		fmt.Printf("name:%s,price:%s",name, price)

	})
}
func  Test()  {
	
}

