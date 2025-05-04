package converter

import (
	"github.com/CSU-2025-1/article-alchemy-service/content_service/internal/domain/model"
	contentv1 "github.com/tclutin/article-alchemy-service-protos/gen/go/content_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertContentModelsToProto(contents []model.Content) *contentv1.GetHistoryResponse {
	items := make([]*contentv1.GetHistoryItem, 0, len(contents))
	for _, item := range contents {
		items = append(items, &contentv1.GetHistoryItem{
			ContentId: item.ContentID,
			Status:    item.Status,
			Url:       item.URL,
			Data:      item.Data,
			Error:     item.Error,
			CreatedAt: timestamppb.New(item.CreatedAt),
		})
	}

	return &contentv1.GetHistoryResponse{
		Items: items,
	}
}
