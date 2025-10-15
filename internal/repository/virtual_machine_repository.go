package repository

import (
	"context"
	"it-backend/internal/model/entity"
)

type VirtualMachineRepository interface {
	GetAll(ctx context.Context) ([]entity.VirtualMachine, error)
	GetOne(ctx context.Context, id uint64) (entity.VirtualMachine, error)
	GetStatusHistory(ctx context.Context, id uint64) ([]entity.VirtualMachineStatusHistory, error)
	Create(ctx context.Context, virtualMachine entity.VirtualMachine, virtualMachineConfig entity.VirtualMachineConfig) (entity.VirtualMachine, error)
	Update(ctx context.Context, virtualMachine entity.VirtualMachine) (entity.VirtualMachine, error)
	UpdateConfig(ctx context.Context, virtualMachineConfig entity.VirtualMachineConfig) (entity.VirtualMachineConfig, error)
	Delete(ctx context.Context, id uint64) error
}
