package main

import (
	"github.com/anjunact/stock-monitor/crawler"

)
func main() {
	 urls := []string{
		 "https://gupiao.baidu.com/stock/sh601000.html",
		 "https://gupiao.baidu.com/stock/sh600048.html",
		 "https://gupiao.baidu.com/stock/sh601228.html",
		 "https://gupiao.baidu.com/stock/sh600717.html",
	 }
	for _, url := range urls{
		 crawler.Scrape(url)
	}
}

