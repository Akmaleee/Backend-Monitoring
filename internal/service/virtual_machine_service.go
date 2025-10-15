package service

import (
	"context"
	"it-backend/internal/model/dto"
)

type VirtualMachineService interface {
	GetAll(ctx context.Context) ([]dto.VirtualMachineResponse, error)
	GetOne(ctx context.Context, id uint64) (dto.VirtualMachineResponse, error)
	GetStatusHistory(ctx context.Context, id uint64) ([]dto.VirtualMachineStatusHistoryResponse, error)
	Create(ctx context.Context, request dto.VirtualMachineRequest) (dto.VirtualMachineResponse, error)
	Update(ctx context.Context, request dto.VirtualMachineRequest, id uint64) (dto.VirtualMachineResponse, error)
	UpdateConfig(ctx context.Context, request dto.VirtualMachineConfigRequest, id uint64) (dto.VirtualMachineConfigResponse, error)
	Delete(ctx context.Context, id uint64) error
}
