package infrastructure

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
)

type ScrapeHandler struct{}

func NewScrapeHandler() *ScrapeHandler {
	return &ScrapeHandler{}
}

func (handler *ScrapeHandler) GetTitleFromURL(URL string) (title string, err error) {
	// Google検索の場合はクエリから検索ワードを抽出
	title, err = extractGoogleTitle(URL)
	if err != nil || title != "" {
		return
	}

	// Getリクエスト
	res, err := http.Get(URL)
	if err != nil {
		return
	}
	defer res.Body.Close()

	// 読み取り
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	// 文字コード判定
	det := chardet.NewTextDetector()
	detRslt, err := det.DetectBest(buf)
	if err != nil {
		return
	}

	// 文字コード変換
	bReader := bytes.NewReader(buf)
	reader, err := charset.NewReaderLabel(detRslt.Charset, bReader)
	if err != nil {
		return
	}

	// HTMLパース
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return
	}

	// titleを抜き出し
	title = doc.Find("title").Text()
	return
}

func extractGoogleTitle(URL string) (title string, err error) {
	u, err := url.Parse(URL)
	if err != nil {
		return
	}
	host := fmt.Sprintf(string(u.Host))
	if host == "www.google.com" || host == "www.google.co.jp" {
		for key, values := range u.Query() {
			if key == "q" {
				for _, v := range values {
					title = fmt.Sprintf("%s - Google 検索", v)
				}
			}
		}
	}
	return
}
