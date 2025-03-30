package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"article-alchemy-service/pkg/models"
)

func GetYandexSharingURL(articleURL string) (string, error) {
	yandexToken := os.Getenv("YANDEX_TOKEN")
	yandexAPI := os.Getenv("YANDEX_API")
	log.Println(yandexToken, yandexAPI)
	if yandexToken == "" {
		log.Println("Warning: YANDEX_TOKEN is not set!")
		return "", errors.New("missing YANDEX_TOKEN")
	}

	jsonData, _ := json.Marshal(models.YandexRequest{ArticleURL: articleURL})
	req, _ := http.NewRequest("POST", yandexAPI, bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", "OAuth "+yandexToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var yandexResp models.YandexResponse
	if err := json.NewDecoder(resp.Body).Decode(&yandexResp); err != nil {
		return "", err
	}

	if yandexResp.Status != "success" {
		return "", errors.New("error API Yandex")
	}

	return yandexResp.SharingURL, nil
}
