package converter

import (
	"errors"
	domainErr "github.com/CSU-2025-1/article-alchemy-service/auth_service/internal/domain/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ConvertErrorToGRPCStatus(err error) error {
	switch {
	case errors.Is(err, domainErr.ErrUserNotFound) || errors.Is(err, domainErr.ErrRefreshTokenNotFound):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, domainErr.ErrUserAlreadyExists):
		return status.Error(codes.AlreadyExists, err.Error())
	case errors.Is(err, domainErr.ErrWrongPassword):
		return status.Error(codes.Unauthenticated, err.Error())
	default:
		return status.Error(codes.Internal, "Internal server error")
	}
}
