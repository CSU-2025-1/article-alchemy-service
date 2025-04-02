package parser

import (
	"article-alchemy-service/pkg/models"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"unicode"

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

	var chapters []models.Chapter

	doc.Find("div.chapter-wrapper").Each(func(i int, chapter *goquery.Selection) {
		chapterTitle := cleanText(chapter.Find("h2").Text())
		var points []string

		chapter.Find("p.thesis-text").Each(func(j int, thesis *goquery.Selection) {
			cleanedText := cleanText(thesis.Text())
			points = append(points, cleanedText)
		})

		chapters = append(chapters, models.Chapter{
			Title:  chapterTitle,
			Points: points,
		})
	})

	jsonData, err := json.Marshal(chapters)
	if err != nil {
		log.Printf("Error converting to JSON: %v", err)
		return "", err
	}

	return string(jsonData), nil
}

func cleanText(text string) string {
	text = strings.ReplaceAll(text, "\u2009", "")
	text = strings.ReplaceAll(text, "\u200b", "")

	text = strings.TrimLeft(text, "• ")
	text = strings.TrimSpace(text)

	var cleanedText strings.Builder
	for _, r := range text {
		if unicode.IsGraphic(r) && !unicode.IsControl(r) {
			cleanedText.WriteRune(r)
		}
	}

	return cleanedText.String()
}
