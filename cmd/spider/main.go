package main

import (
	"go-test/spider"
	"log"
)

func main() {
	//err := spider.GetCityData()
	err := spider.GetDataByJS()
	log.Println(err)
}
