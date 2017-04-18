package main

import (
	"github.com/anjunact/stock-monitor/crawler"
	"github.com/anjunact/stock-monitor/database"

)
func main() {
	//crawler.Scrape()
	crawler.Test()
	database.Test()
}

