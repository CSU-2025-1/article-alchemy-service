package service

import (
	"article-alchemy-service/internal/client"
	"article-alchemy-service/internal/parser"
	"article-alchemy-service/pkg/models"
	"encoding/json"
)

func GetSummary(articleURL string) ([]models.Chapter, error) {
	sharingURL, err := client.GetYandexSharingURL(articleURL)
	if err != nil {
		return nil, err
	}

	jsonData, err := parser.ParseSummary(sharingURL)
	if err != nil {
		return nil, err
	}
	var summary []models.Chapter
	if err := json.Unmarshal(jsonData, &summary); err != nil {
		return nil, err
	}
	return summary, nil
}
