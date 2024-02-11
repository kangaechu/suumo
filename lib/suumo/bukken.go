package suumo

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strings"
)
import "github.com/gocolly/colly"

type Bukken struct {
	Name           string // 物件名
	Link           string // 物件のURL
	Price          string // 価格
	Madori         string // 間取り
	Area           string // 占有面積
	SouKosuu       string // 総戸数
	Floor          string // 所在階
	Kouzou         string // 構造・階建て
	BuiltAt        string // 完成時期（築年月）
	Address        string // 住所
	Access         string // 交通
	KanriHi        string // 管理費
	ShuzenHi       string // 修繕積立金
	HikiwatshiJiki string // 引渡可能時期
	Muki           string // 向き
	Kenri          string // 敷地の権利形態
	Youto          string // 用途地域
	Parking        string // 駐車場
	Sekou          string // 施工
	CreatedAt      string // 登録日
	UpdatedAt      string // 更新日
}

// FetchBukken は物件の情報を取得する
func FetchBukken(url string) (Bukken, error) {
	c := colly.NewCollector()
	valueRe := regexp.MustCompile(`\[.*?]`) // [ □支払シミュレーション ] などの文字列を削除するための正規表現

	bukkenInfo := make(map[string]string)
	var bukken Bukken
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
			CreatedAt:      containsInMap(bukkenInfo, "情報提供日"),
			UpdatedAt:      containsInMap(bukkenInfo, "次回更新日"),
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
