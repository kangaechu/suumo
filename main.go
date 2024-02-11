package main

import (
	"encoding/json"
	"fmt"
	"github.com/kangaechu/sumoo/lib/sumoo"
	"os"
)

func main() {
	url := os.Getenv("SUUMO_URL")
	if url == "" {
		panic("環境変数 SUUMO_URL が設定されていません")
	}
	bukkenRss, err := sumoo.FetchBukkenRSS(url)
	if err != nil {
		panic(err)
	}
	bukkens := make([]sumoo.Bukken, len(bukkenRss))
	for _, b := range bukkenRss {
		bukken, err := sumoo.FetchBukken(b.Link)
		if err != nil {
			panic(err)
		}
		bukkens = append(bukkens, bukken)
	}
	// jsonとして出力
	out, err := json.Marshal(bukkens)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(out))
}
