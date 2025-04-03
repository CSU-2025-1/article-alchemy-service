package interceptor

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	contentv1 "github.com/tclutin/article-alchemy-service-protos/gen/go/content_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"time"
)

type LimiterInterceptor struct {
	methods map[string]bool
	client  *redis.Client
	limit   int64
	ttl     time.Duration
}

func NewLimiterInterceptor(client *redis.Client) *LimiterInterceptor {
	methods := map[string]bool{
		contentv1.ContentService_ExtractContentPreview_FullMethodName: true,
	}

	return &LimiterInterceptor{
		methods: methods,
		client:  client,
		limit:   5,
		ttl:     24 * time.Hour,
	}
}

func (l *LimiterInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if !l.methods[info.FullMethod] {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "metadata is not provided")
		}

		ip := md["x-client-ip"]
		if len(ip) == 0 {
			return nil, status.Error(codes.Unauthenticated, "x-client-ip header is missing")
		}

		key := fmt.Sprintf("ip_limit:%s", ip[0])

		count, err := l.client.Incr(ctx, key).Result()
		if err != nil {
			return nil, status.Error(codes.Internal, "failed to check ip limit")
		}

		if count == 1 {
			if _, err = l.client.Expire(ctx, key, l.ttl).Result(); err != nil {
				return nil, status.Error(codes.Internal, "failed to set ip limit")
			}
		}

		if count > l.limit {
			return nil, status.Error(codes.ResourceExhausted, fmt.Sprintf("ip limit exceeded: %d requests per day allowed", l.limit))
		}

		return handler(ctx, req)
	}
}
