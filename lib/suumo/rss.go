package suumo

import (
	"github.com/mmcdole/gofeed"
	"io"
	"net/http"
)

// BukkenRSS はRSSで取得した物件情報を表す
type BukkenRSS struct {
	Title       string
	Link        string
	Description string
}

// FetchBukkenRSS は指定されたURLから物件情報を取得する
func FetchBukkenRSS(url string) ([]BukkenRSS, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	bukkens, err := fetchBukkenRSS(resp.Body)
	if err != nil {
		return nil, err
	}
	return bukkens, nil
}

// fetchBukkenRSS は指定されたリーダーから物件情報を取得する
func fetchBukkenRSS(reader io.Reader) ([]BukkenRSS, error) {
	fp := gofeed.NewParser()
	feed, err := fp.Parse(reader)
	if err != nil {
		return nil, err
	}
	bukkens := make([]BukkenRSS, len(feed.Items))
	for i, item := range feed.Items {
		bukkens[i] = BukkenRSS{
			Title:       item.Title,
			Link:        item.Link,
			Description: item.Description,
		}
	}
	return bukkens, nil
}
