package main

import (
	"fmt"
	"github.com/kangaechu/suumo/lib/suumo"
	"os"
)

func main() {
	url := os.Getenv("SUUMO_URL")
	bukkenRss, err := suumo.FetchBukkenRSS(url)
	if err != nil {
		return
	}
	bukkens := make([]suumo.Bukken, len(bukkenRss))
	for _, b := range bukkenRss {
		bukken, err := suumo.FetchBukken(b.Link)
		if err != nil {
			return
		}
		bukkens = append(bukkens, bukken)
	}
	fmt.Println(bukkens)
}
