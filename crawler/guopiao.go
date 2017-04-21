package crawler

import (
	"fmt"
	"log"
	"github.com/PuerkitoBio/goquery"
	"github.com/anjunact/stock-monitor/models"
	"strconv"
	"strings"
)

func Scrape(url string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".stock-bets").Each(func(i int, s *goquery.Selection) {
		name := s.Find(".bets-name").Text()
		ar := strings.Fields(name)
		name = ar[0]
		code := s.Find(".bets-name").Find("span").Text()
		price :=   s.Find("._close").Text()
		fmt.Printf("name:%s,code:%s,price:%s \n",name,code, price)

		stock ,err :=  models.GetStock(code)
		if err != nil{
			fmt.Printf("error:%v",err)
			f_price ,_ := strconv.ParseFloat(price,64)
			stock = models.Stock{ Name:name,Code:code,Price:f_price}
			stock.Create()
		}else{
			stock.Name=name
			stock.Update()
		}
	})
}
func  Test()  {

}

