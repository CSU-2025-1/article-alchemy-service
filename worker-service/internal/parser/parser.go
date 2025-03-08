package parser

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ParseSummary(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	var summaryText string
	doc.Find("div.chapter-wrapper").Each(func(i int, chapter *goquery.Selection) {
		chapterTitle := chapter.Find("h2").Text()
		summaryText += chapterTitle

		chapter.Find("p.thesis-text").Each(func(j int, thesis *goquery.Selection) {
			summaryText += thesis.Text()
		})
	})

	return summaryText, nil
}
