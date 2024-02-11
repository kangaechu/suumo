package sumoo

import (
	"github.com/google/go-cmp/cmp"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

func TestFetchBukken(t *testing.T) {
	srv := newTestServer()
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    Bukken
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				url: srv.URL + "/test1",
			},
			want: Bukken{
				Name:           "パークホームズ豊洲ザ・レジデンス",
				Link:           "http://127.0.0.1/test1",
				Price:          "1億800万円",
				Madori:         "3LDK",
				Area:           "74.04m2（壁芯）",
				SouKosuu:       "693戸",
				Floor:          "4階",
				Kouzou:         "RC22階建",
				BuiltAt:        "2016年10月",
				Address:        "東京都江東区豊洲５",
				Access:         "東京メトロ有楽町線「豊洲」歩5分新交通ゆりかもめ「豊洲」歩4分",
				KanriHi:        "1万9620円／月（委託(通勤)）",
				ShuzenHi:       "9230円／月",
				HikiwatshiJiki: "相談",
				Muki:           "北西",
				Kenri:          "所有権",
				Youto:          "準工業",
				Parking:        "空無",
				Sekou:          "株式会社大林組 東京本店",
				ProvidedAt:     "2024年2月10日",
				NextUpdatedAt:  "情報提供より８日以内に更新",
				CreatedAt:      "2024/02/01T00:00:00+09:00",
				UpdatedAt:      "2024/02/01T00:00:00+09:00",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FetchBukken(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchBukken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// モックサーバのportが実行ごとに変わるため、Linkのポート番号を削除して比較
			got.Link = trimPort(got.Link)
			// CreatedAt, UpdatedAtは実行ごとに変わるため、固定値に変更
			got.CreatedAt = "2024/02/01T00:00:00+09:00"
			got.UpdatedAt = "2024/02/01T00:00:00+09:00"

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("FetchBukken() is mismatch (-: got, +: want) : %s", diff)
			}
		})
	}
}

func newUnstartedTestServer() *httptest.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/test1", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		testdata, _ := os.ReadFile("testdata/test1.html")
		_, _ = w.Write(testdata)
	})
	return httptest.NewUnstartedServer(mux)
}

func newTestServer() *httptest.Server {
	srv := newUnstartedTestServer()
	srv.Start()
	return srv
}

func trimPort(s string) string {
	u, err := url.Parse(s)
	if err != nil {
		return s
	}
	return u.Scheme + "://" + u.Hostname() + u.Path
}
