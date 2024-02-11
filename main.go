package main

import (
	"fmt"
	"github.com/kangaechu/sumoo/lib/sumoo"
	"os"
)

func main() {
	url := os.Getenv("SUUMO_URL")
	bukkenRss, err := sumoo.FetchBukkenRSS(url)
	if err != nil {
		return
	}
	bukkens := make([]sumoo.Bukken, len(bukkenRss))
	for _, b := range bukkenRss {
		bukken, err := sumoo.FetchBukken(b.Link)
		if err != nil {
			return
		}
		bukkens = append(bukkens, bukken)
	}
	fmt.Println(bukkens)
}
