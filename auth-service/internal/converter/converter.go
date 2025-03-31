package converter

import (
	"errors"
	"github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ConvertErrorToGRPCStatus(err error) error {
	switch {
	case errors.Is(err, domain.ErrUserNotFound) || errors.Is(err, domain.ErrRefreshTokenNotFound):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, domain.ErrUserAlreadyExists):
		return status.Error(codes.AlreadyExists, err.Error())
	case errors.Is(err, domain.ErrWrongPassword):
		return status.Error(codes.Unauthenticated, err.Error())
	default:
		return status.Error(codes.Internal, "Internal server error")
	}
}
