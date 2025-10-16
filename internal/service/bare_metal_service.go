package service

import (
	"context"
	"it-backend/internal/model/dto"
)

type BareMetalService interface {
	GetAll(ctx context.Context) ([]dto.BareMetalResponse, error)
	GetAllNodes(ctx context.Context) ([]dto.BareMetalNodeWithStatusResponse, error)
	GetOne(ctx context.Context, id uint64) (dto.BareMetalResponse, error)
	GetNodeStatusHistory(ctx context.Context, id uint64) ([]dto.BareMetalNodeStatusHistoryResponse, error)
	Create(ctx context.Context, request dto.BareMetalRequest) (dto.BareMetalResponse, error)
	Update(ctx context.Context, request dto.BareMetalRequest, id uint64) (dto.BareMetalResponse, error)
	Delete(ctx context.Context, id uint64) error
}
