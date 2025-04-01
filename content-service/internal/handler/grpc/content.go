package grpc

import (
	"context"
	"github.com/tclutin/article-alchemy-service-protos/gen/go/content_v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ContentHandler struct {
	content_v1.UnimplementedContentServiceServer
}

func NewContentHandler() *ContentHandler {
	return &ContentHandler{}
}

func (c *ContentHandler) Register(grpcServer *grpc.Server) {
	content_v1.RegisterContentServiceServer(grpcServer, c)
}

func (c *ContentHandler) ExtractContentPreview(ctx context.Context, request *content_v1.ExtractContentPreviewRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ContentHandler) ExtractContent(ctx context.Context, request *content_v1.ExtractContentRequest) (*content_v1.ExtractContentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ContentHandler) GetHistory(ctx context.Context, empty *emptypb.Empty) (*content_v1.GetHistoryResponse, error) {
	//TODO implement me
	panic("implement me")
}
