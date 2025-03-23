package handler

import (
	"encoding/json"
	"net/http"

	"article-alchemy-service/internal/service"
	"article-alchemy-service/pkg/models"
)

func SummarizeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.YandexRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}

	if req.ArticleURL == "" {
		http.Error(w, "Missing 'article_url' field", http.StatusBadRequest)
		return
	}

	summary, err := service.GetSummary(req.ArticleURL)
	if err != nil {
		http.Error(w, "Failed to get summary: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(models.SummaryResponse{Summary: summary})
	if err != nil {
		http.Error(w, "Failed to encode JSON response: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
