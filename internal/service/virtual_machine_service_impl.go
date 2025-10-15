package service

import (
	"context"
	"it-backend/internal/model/dto"
	"it-backend/internal/model/entity"
	"it-backend/internal/repository"

	"github.com/pkg/errors"
)

type VirtualMachineRepositoryImpl struct {
	VirtualMachineRepository repository.VirtualMachineRepository
}

func NewVirtualMachineService(virtualMachineRepository repository.VirtualMachineRepository) *VirtualMachineRepositoryImpl {
	return &VirtualMachineRepositoryImpl{
		VirtualMachineRepository: virtualMachineRepository,
	}
}

func (s *VirtualMachineRepositoryImpl) GetAll(ctx context.Context) ([]dto.VirtualMachineResponse, error) {
	var responses []dto.VirtualMachineResponse

	virtualMachines, err := s.VirtualMachineRepository.GetAll(ctx)
	if err != nil {
		return responses, errors.Wrap(err, "failed to find virtual machines")
	}

	for _, h := range virtualMachines {
		res := dto.VirtualMachineResponse{
			ID:                   h.ID,
			BareMetalID:          h.BareMetalID,
			BareMetal:            h.BareMetal,
			BareMetalNodeID:      h.BareMetalNodeID,
			BareMetalNode:        h.BareMetalNode,
			VmID:                 h.VmID,
			Code:                 h.Code,
			Name:                 h.Name,
			Cpu:                  h.Cpu,
			Memory:               h.Memory,
			Disk:                 h.Disk,
			VirtualMachineStatus: h.VirtualMachineStatus,
			CreatedAt:            h.CreatedAt,
			UpdatedAt:            h.UpdatedAt,
		}
		responses = append(responses, res)
	}

	return responses, nil
}

func (s *VirtualMachineRepositoryImpl) GetOne(ctx context.Context, id uint64) (dto.VirtualMachineResponse, error) {
	var response dto.VirtualMachineResponse

	virtualMachine, err := s.VirtualMachineRepository.GetOne(ctx, id)
	if err != nil {
		return response, errors.Wrap(err, "failed to find virtual machine")
	}

	response = dto.VirtualMachineResponse{
		ID:                   virtualMachine.ID,
		BareMetalID:          virtualMachine.BareMetalID,
		BareMetal:            virtualMachine.BareMetal,
		BareMetalNodeID:      virtualMachine.BareMetalNodeID,
		BareMetalNode:        virtualMachine.BareMetalNode,
		VmID:                 virtualMachine.VmID,
		Code:                 virtualMachine.Code,
		Name:                 virtualMachine.Name,
		Cpu:                  virtualMachine.Cpu,
		Memory:               virtualMachine.Memory,
		Disk:                 virtualMachine.Disk,
		VirtualMachineConfig: virtualMachine.VirtualMachineConfig,
		VirtualMachineStatus: virtualMachine.VirtualMachineStatus,
		CreatedAt:            virtualMachine.CreatedAt,
		UpdatedAt:            virtualMachine.UpdatedAt,
	}

	return response, nil
}

func (s *VirtualMachineRepositoryImpl) GetStatusHistory(ctx context.Context, id uint64) ([]dto.VirtualMachineStatusHistoryResponse, error) {
	var responses []dto.VirtualMachineStatusHistoryResponse

	virtualMachines, err := s.VirtualMachineRepository.GetStatusHistory(ctx, id)
	if err != nil {
		return responses, errors.Wrap(err, "failed to find virtual machine's status histories")
	}

	for _, h := range virtualMachines {
		res := dto.VirtualMachineStatusHistoryResponse{
			ID:               h.ID,
			VirtualMachineID: h.VirtualMachineID,
			Type:             h.Type,
			Status:           h.Status,
			CreatedAt:        h.CreatedAt,
			UpdatedAt:        h.UpdatedAt,
		}
		responses = append(responses, res)
	}

	return responses, nil
}

func (s *VirtualMachineRepositoryImpl) Create(ctx context.Context, request dto.VirtualMachineRequest) (dto.VirtualMachineResponse, error) {
	var response dto.VirtualMachineResponse

	if err := request.Validate(); err != nil {
		return response, err
	}

	virtualMachine, err := s.VirtualMachineRepository.Create(ctx, entity.VirtualMachine{
		BareMetalID:     request.BareMetalID,
		BareMetalNodeID: request.BareMetalNodeID,
		VmID:            request.VmID,
		Code:            request.Code,
		Name:            request.Name,
		Cpu:             request.Cpu,
		Memory:          request.Memory,
		Disk:            request.Disk,
	}, entity.VirtualMachineConfig{
		IsAlertStatus: request.VirtualMachineConfig.IsAlertStatus,
		IsAlertDisk:   request.VirtualMachineConfig.IsAlertDisk,
		ThresholdDisk: request.VirtualMachineConfig.ThresholdDisk,
	})

	if err != nil {
		return response, errors.Wrap(err, "failed to create virtual machine")
	}

	response = dto.VirtualMachineResponse{
		ID:                   virtualMachine.ID,
		BareMetalID:          virtualMachine.BareMetalID,
		BareMetalNodeID:      virtualMachine.BareMetalNodeID,
		VmID:                 virtualMachine.VmID,
		Code:                 virtualMachine.Code,
		Name:                 virtualMachine.Name,
		Cpu:                  virtualMachine.Cpu,
		Memory:               virtualMachine.Memory,
		Disk:                 virtualMachine.Disk,
		VirtualMachineConfig: virtualMachine.VirtualMachineConfig,
		CreatedAt:            virtualMachine.CreatedAt,
		UpdatedAt:            virtualMachine.UpdatedAt,
	}

	return response, nil
}

func (s *VirtualMachineRepositoryImpl) Update(ctx context.Context, request dto.VirtualMachineRequest, id uint64) (dto.VirtualMachineResponse, error) {
	var response dto.VirtualMachineResponse

	virtualMachine, err := s.VirtualMachineRepository.Update(ctx, entity.VirtualMachine{
		ID:              id,
		BareMetalID:     request.BareMetalID,
		BareMetalNodeID: request.BareMetalNodeID,
		VmID:            request.VmID,
		Code:            request.Code,
		Name:            request.Name,
		Cpu:             request.Cpu,
		Memory:          request.Memory,
		Disk:            request.Disk,
	})

	if err != nil {
		return response, errors.Wrap(err, "failed to update virtual machine")
	}

	response = dto.VirtualMachineResponse{
		ID:              virtualMachine.ID,
		BareMetalID:     virtualMachine.BareMetalID,
		BareMetalNodeID: virtualMachine.BareMetalNodeID,
		VmID:            virtualMachine.VmID,
		Code:            virtualMachine.Code,
		Name:            virtualMachine.Name,
		Cpu:             virtualMachine.Cpu,
		Memory:          virtualMachine.Memory,
		Disk:            virtualMachine.Disk,
		CreatedAt:       virtualMachine.CreatedAt,
		UpdatedAt:       virtualMachine.UpdatedAt,
	}

	return response, nil
}

func (s *VirtualMachineRepositoryImpl) UpdateConfig(ctx context.Context, request dto.VirtualMachineConfigRequest, id uint64) (dto.VirtualMachineConfigResponse, error) {
	var response dto.VirtualMachineConfigResponse

	virtualMachineConfig, err := s.VirtualMachineRepository.UpdateConfig(ctx, entity.VirtualMachineConfig{
		ID:               id,
		VirtualMachineID: request.VirtualMachineID,
		IsAlertStatus:    request.IsAlertStatus,
		IsAlertDisk:      request.IsAlertDisk,
		ThresholdDisk:    request.ThresholdDisk,
	})

	if err != nil {
		return response, errors.Wrap(err, "failed to update virtual machine's configs")
	}

	response = dto.VirtualMachineConfigResponse{
		ID:               virtualMachineConfig.ID,
		VirtualMachineID: virtualMachineConfig.VirtualMachineID,
		IsAlertStatus:    virtualMachineConfig.IsAlertStatus,
		IsAlertDisk:      virtualMachineConfig.IsAlertDisk,
		ThresholdDisk:    virtualMachineConfig.ThresholdDisk,
		CreatedAt:        virtualMachineConfig.CreatedAt,
		UpdatedAt:        virtualMachineConfig.UpdatedAt,
	}

	return response, nil
}

func (s *VirtualMachineRepositoryImpl) Delete(ctx context.Context, id uint64) error {
	return s.VirtualMachineRepository.Delete(ctx, id)
}
