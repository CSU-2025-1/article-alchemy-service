package service

import (
	"article-alchemy-service/internal/client"
	"article-alchemy-service/internal/parser"
)

func GetSummary(articleURL string) (string, error) {
	sharingURL, err := client.GetYandexSharingURL(articleURL)
	if err != nil {
		return "", err
	}

	return parser.ParseSummary(sharingURL)
}
