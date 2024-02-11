package suumo

import (
	"bytes"
	"github.com/google/go-cmp/cmp"
	"io"
	"os"
	"testing"
)

func Test_fetchBukkenRSS(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	test1, _ := os.ReadFile("testdata/rss1.xml")

	tests := []struct {
		name    string
		args    args
		want    []BukkenRSS
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				reader: bytes.NewReader(test1),
			},
			want: []BukkenRSS{{
				Title:       "物件名：パークホームズ豊洲ザ・レジデン…",
				Link:        "https://suumo.jp/jj/bukken/shosai/JJ010FJ100/?ar=030&bs=011&nc=73305798&ta=13&sc=13108",
				Description: "東京都江東区豊洲５1億800万円東京メトロ有楽町線豊洲徒歩5分74.04m&sup2;（壁芯）11.88m&sup2;3LDK2016年10月",
			}, {
				Title:       "物件名：ベイズタワー＆ガーデン",
				Link:        "https://suumo.jp/jj/bukken/shosai/JJ010FJ100/?ar=030&bs=011&nc=74260274&ta=13&sc=13108",
				Description: "東京都江東区豊洲６1億880万円新交通ゆりかもめ新豊洲徒歩6分69.22m&sup2;（20.93坪）（壁芯）16.52m&sup2;3LDK2016年6月",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fetchBukkenRSS(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("fetchBukkenRSS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("fetchBukkenRSS() is mismatch (-: got, +: want) : %s", diff)
			}
		})
	}
}
