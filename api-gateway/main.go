package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	authpb "github.com/tclutin/article-alchemy-service-protos/gen/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
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

func main() {
	mux := runtime.NewServeMux(runtime.WithMetadata(AttachTokenMetadata))

	grpcOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		//grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	}

	err := authpb.RegisterAuthServiceHandlerFromEndpoint(
		context.Background(),
		mux,
		"localhost:9091",
		grpcOpts)

	if err != nil {
		panic(err)
	}

	err = http.ListenAndServe(":8080", mux)
	fmt.Println(err)
	if err != nil {
		panic(err)
	}

}
