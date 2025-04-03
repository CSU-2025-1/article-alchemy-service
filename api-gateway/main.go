package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tclutin/article-alchemy-service-protos/gen/go/auth_v1"
	"github.com/tclutin/article-alchemy-service-protos/gen/go/content_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log/slog"
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
	slog.Info("Чепенькус чекай, должен быть ip тут", slog.String("ip", ip), slog.String("remote", r.RemoteAddr), slog.String("host", r.Host))
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

// BLAZING
func main() {
	mux := runtime.NewServeMux(
		runtime.WithMetadata(AttachTokenMetadata),
		runtime.WithMetadata(AttachIpAddressMetadata),
	)

	grpcOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	err := auth_v1.RegisterAuthServiceHandlerFromEndpoint(
		context.Background(),
		mux,
		"localhost:9091",
		grpcOpts)

	if err != nil {
		panic(err)
	}

	err = content_v1.RegisterContentServiceHandlerFromEndpoint(
		context.Background(),
		mux,
		"localhost:9092",
		grpcOpts)

	if err != nil {
		panic(err)
	}

	corsMux := AllowCORSMiddleware(mux)

	err = http.ListenAndServe(":8091", corsMux)
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
}
