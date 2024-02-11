package sumoo

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
	"time"
)
import "github.com/gocolly/colly"

type Bukken struct {
	Name           string `json:"id"`              // 物件名
	Link           string `json:"link"`            // 物件のURL
	Price          string `json:"price"`           // 価格
	Madori         string `json:"madori"`          // 間取り
	Area           string `json:"area"`            // 占有面積
	SouKosuu       string `json:"soukosuu"`        // 総戸数
	Floor          string `json:"floor"`           // 所在階
	Kouzou         string `json:"kouzou"`          // 構造・階建て
	BuiltAt        string `json:"built_at"`        // 完成時期（築年月）
	Address        string `json:"address"`         // 住所
	Access         string `json:"access"`          // 交通
	KanriHi        string `json:"kanri_hi"`        // 管理費
	ShuzenHi       string `json:"shuzen_hi"`       // 修繕積立金
	HikiwatshiJiki string `json:"hikiwatshiJiki"`  // 引渡可能時期
	Muki           string `json:"muki"`            // 向き
	Kenri          string `json:"kenri"`           // 敷地の権利形態
	Youto          string `json:"youto"`           // 用途地域
	Parking        string `json:"parking"`         // 駐車場
	Sekou          string `json:"sekou"`           // 施工
	ProvidedAt     string `json:"provided_at"`     // 情報提供日
	NextUpdatedAt  string `json:"next_updated_at"` // 次回更新日
	CreatedAt      string `json:"created_at"`      // 登録日
	UpdatedAt      string `json:"updated_at"`      // 更新日
}

// FetchBukken は物件の情報を取得する
func FetchBukken(url string) (Bukken, error) {
	c := colly.NewCollector()
	valueRe := regexp.MustCompile(`\[.*?]`) // [ □支払シミュレーション ] などの文字列を削除するための正規表現

	bukkenInfo := make(map[string]string)
	var bukken Bukken
	now := time.Now()
	c.OnHTML("body", func(e *colly.HTMLElement) {
		e.DOM.Find("th").Each(func(i int, k *goquery.Selection) {
			var key, value string
			key = k.Text()
			key = strings.ReplaceAll(key, "ヒント", "")
			key = strings.TrimSpace(key)
			// thの次のtdを取得
			k.Next().Each(func(i int, v *goquery.Selection) {
				value = v.Text()
				value = valueRe.ReplaceAllString(value, "")
				value = strings.TrimSpace(value)
				// 交通などには複数の要素があるので、その間の改行やタブを削除
				value = strings.ReplaceAll(value, "\n", "")
				value = strings.ReplaceAll(value, "\t", "")
			})
			bukkenInfo[key] = value
		})
		bukken = Bukken{
			Name:           containsInMap(bukkenInfo, "物件名"),
			Link:           url,
			Price:          containsInMap(bukkenInfo, "価格"),
			Madori:         containsInMap(bukkenInfo, "間取り"),
			Area:           containsInMap(bukkenInfo, "専有面積"),
			SouKosuu:       containsInMap(bukkenInfo, "総戸数"),
			Floor:          containsInMap(bukkenInfo, "所在階"),
			Kouzou:         containsInMap(bukkenInfo, "構造・階建て"),
			BuiltAt:        containsInMap(bukkenInfo, "完成時期（築年月）"),
			Address:        containsInMap(bukkenInfo, "住所"),
			Access:         containsInMap(bukkenInfo, "交通"),
			KanriHi:        containsInMap(bukkenInfo, "管理費"),
			ShuzenHi:       containsInMap(bukkenInfo, "修繕積立金"),
			HikiwatshiJiki: containsInMap(bukkenInfo, "引渡可能時期"),
			Muki:           containsInMap(bukkenInfo, "向き"),
			Kenri:          containsInMap(bukkenInfo, "敷地の権利形態"),
			Youto:          containsInMap(bukkenInfo, "用途地域"),
			Parking:        containsInMap(bukkenInfo, "駐車場"),
			Sekou:          containsInMap(bukkenInfo, "施工"),
			ProvidedAt:     containsInMap(bukkenInfo, "情報提供日"),
			NextUpdatedAt:  containsInMap(bukkenInfo, "次回更新日"),
			CreatedAt:      now.Format("2006-01-02T15:04:05-07:00"),
			UpdatedAt:      now.Format("2006-01-02T15:04:05-07:00"),
		}
	})
	err := c.Visit(url)
	if err != nil {
		return Bukken{}, err
	}
	return bukken, nil
}

// containsInMap はmapにkeyが存在するかを確認し、存在する場合はvalueを返す
func containsInMap(m map[string]string, key string) string {
	_, ok := m[key]
	if ok {
		return m[key]
	} else {
		return ""
	}
}
