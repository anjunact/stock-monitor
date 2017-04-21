package main

import (
	"github.com/anjunact/stock-monitor/crawler"

	"time"
	"fmt"
	"log"
)
func main() {
	 urls := []string{
		 "https://gupiao.baidu.com/stock/sh601000.html",
		 "https://gupiao.baidu.com/stock/sh600048.html",
		 "https://gupiao.baidu.com/stock/sh601228.html",
		 "https://gupiao.baidu.com/stock/sh600717.html",
	 }
	for _, url := range urls{
		 //crawler.Scrape(url)
		crawler.Test()
		fmt.Println(url)
	}
	ticker()

}
func  ticker()  {
	timer := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-timer.C:
			go func() {
				log.Println(time.Now())
			}()
		}
	}
}

