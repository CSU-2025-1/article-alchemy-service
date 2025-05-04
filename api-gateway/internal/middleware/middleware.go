package middleware

import (
	"context"
	"google.golang.org/grpc/metadata"
	"net"
	"net/http"
	"strings"
)

func AttachTokenMetadata(ctx context.Context, r *http.Request) metadata.MD {
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
		token := strings.TrimPrefix(authHeader, "Bearer ")
		return metadata.Pairs("token", token)
	}
	return nil
}

// AttachIpAddressMetadata TODO: x-forwarded-for, x-real-ip
func AttachIpAddressMetadata(ctx context.Context, r *http.Request) metadata.MD {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return nil
	}
	return metadata.Pairs("x-client-ip", ip)
}

func AllowCORSMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	})
}
