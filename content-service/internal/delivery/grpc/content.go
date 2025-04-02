package grpc

import (
	"context"
	"fmt"
	"github.com/CSU-2025-1/article-alchemy-service/content_service/internal/domain"
	"github.com/CSU-2025-1/article-alchemy-service/content_service/internal/domain/model"
	contentv1 "github.com/tclutin/article-alchemy-service-protos/gen/go/content_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ContentHandler struct {
	contentv1.UnimplementedContentServiceServer
	contentService domain.ContentService
}

func NewContentHandler(service domain.ContentService) *ContentHandler {
	return &ContentHandler{
		contentService: service,
	}
}

func (c *ContentHandler) Register(grpcServer *grpc.Server) {
	contentv1.RegisterContentServiceServer(grpcServer, c)
}

func (c *ContentHandler) ExtractContentPreview(ctx context.Context, request *contentv1.ExtractContentPreviewRequest) (*emptypb.Empty, error) {
	err := c.contentService.ExtractContentPreview(ctx, model.ExtractContentPreviewDTO{
		Email: request.GetEmail(),
		Url:   request.GetUrl(),
	})

	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server Error")
	}

	return &emptypb.Empty{}, nil
}

func (c *ContentHandler) ExtractContent(ctx context.Context, request *contentv1.ExtractContentRequest) (*contentv1.ExtractContentResponse, error) {
	userId := ctx.Value("userId")
	if userId == nil {
		return nil, status.Error(codes.PermissionDenied, "User not found in context")
	}

	//need email

	contentId, err := c.contentService.ExtractContent(ctx, model.ExtractContentDTO{
		UserID: userId.(uint64),
		Email:  "tclutin01@gmail.com",
		Url:    request.GetUrl(),
	})

	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &contentv1.ExtractContentResponse{
		ContentId: contentId,
	}, nil
}

func (c *ContentHandler) GetHistory(ctx context.Context, empty *emptypb.Empty) (*contentv1.GetHistoryResponse, error) {
	userId := ctx.Value("userId")
	if userId == nil {
		return nil, status.Error(codes.PermissionDenied, "User not found in context")
	}

	contents, err := c.contentService.GetContentsByUserId(ctx, userId.(uint64))
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	fmt.Println(contents)

	return nil, err
}
