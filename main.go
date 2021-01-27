package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	c.OnHTML("h1", func(element *colly.HTMLElement) {

		fmt.Println("Vehicle : ", element.Text,"\n")

	})
	c.OnHTML(".amount--3NTpl", func(element *colly.HTMLElement) {
		fmt.Println("Price : ", element.Text,"\n")

	})
	c.OnHTML(".sub-title--37mkY", func(element *colly.HTMLElement) {
		fmt.Println("Posted On: ", element.Text,"\n")

	})
	c.OnHTML(".two-columns--19Hyo", func(element *colly.HTMLElement) {
		fmt.Println(element.Text,"\n")
	})



	desc := false
	c.OnHTML(".description--1nRbz > p", func(element *colly.HTMLElement) {
		if !desc {

			fmt.Println("Description : ", element.Text)
		} else {
			fmt.Println(element.Text ,"\n")
		}
		desc = true
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Website : ", r.URL.String(),"\n")
	})

	c.Visit("https://ikman.lk/en/ad/suzuki-wagon-r-stingray-2015-for-sale-colombo-1416")

	//c.OnHTML(".phone-numbers--2COKR", func(element *colly.HTMLElement) {
	//	fmt.Println("Mobile Number : ", element.Text,"\n")
	//
	//})
}
