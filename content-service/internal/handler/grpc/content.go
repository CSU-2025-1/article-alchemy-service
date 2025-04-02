package grpc

import (
	"context"
	contentv1 "github.com/tclutin/article-alchemy-service-protos/gen/go/content_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ContentHandler struct {
	contentv1.UnimplementedContentServiceServer
}

func NewContentHandler() *ContentHandler {
	return &ContentHandler{}
}

func (c *ContentHandler) Register(grpcServer *grpc.Server) {
	contentv1.RegisterContentServiceServer(grpcServer, c)
}

func (c *ContentHandler) ExtractContentPreview(ctx context.Context, request *contentv1.ExtractContentPreviewRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ContentHandler) ExtractContent(ctx context.Context, request *contentv1.ExtractContentRequest) (*contentv1.ExtractContentResponse, error) {
	userId := ctx.Value("userId")
	if userId == nil {
		return nil, status.Error(codes.PermissionDenied, "User not found in context")
	}

	//need email
	panic("implement me")
}

func (c *ContentHandler) GetHistory(ctx context.Context, empty *emptypb.Empty) (*contentv1.GetHistoryResponse, error) {
	userId := ctx.Value("userId")
	if userId == nil {
		return nil, status.Error(codes.PermissionDenied, "User not found in context")
	}
	panic("implement me")
}
