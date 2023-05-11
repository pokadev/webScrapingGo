package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

type Quote struct {
	Quote string `json:"quote"`
	Author string `json:"author"`
}


func main()  {
	quotes := []Quote{}
	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com"),
	)
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X" )
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response Code", r.StatusCode)
	})
	c.OnError(func(r *colly.Response, err error) {
    fmt.Println("Something went wrong:", err)
  })

	c.OnHTML(".quote", func(h *colly.HTMLElement) {
		div := h.DOM
		quote := div.Find(".text").Text()
		author := div.Find(".author").Text()
		q :=Quote{
			Quote: quote,
      Author: author,
		}
		quotes = append(quotes, q)
	})

	c.Visit("https://quotes.toscrape.com/random")
	fmt.Println(quotes)
}