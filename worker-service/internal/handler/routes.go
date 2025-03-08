package handler

import "net/http"

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/summarize", SummarizeHandler)
}
