package repository

import (
	"context"
	"it-backend/internal/model/entity"
)

type BareMetalRepository interface {
	GetAll(ctx context.Context) ([]entity.BareMetal, error)
	GetAllNodes(ctx context.Context) ([]entity.BareMetalNode, error)
	GetOne(ctx context.Context, id uint64) (entity.BareMetal, error)
	GetNodeStatusHistory(ctx context.Context, id uint64) ([]entity.BareMetalNodeStatusHistory, error)
	Create(ctx context.Context, bareMetal entity.BareMetal) (entity.BareMetal, error)
	Update(ctx context.Context, bareMetal entity.BareMetal) (entity.BareMetal, error)
	Delete(ctx context.Context, id uint64) error
}
