package service

import (
	"context"
	"it-backend/internal/model/dto"
	"it-backend/internal/model/entity"
	"it-backend/internal/repository"

	"github.com/pkg/errors"
)

//	type BareMetalRepositoryImpl struct {
//		BareMetalRepository repository.BareMetalRepository
//	}
type BareMetalServiceImpl struct {
	BareMetalRepository repository.BareMetalRepository
}

func NewBareMetalService(bareMetalRepository repository.BareMetalRepository) *BareMetalServiceImpl {
	return &BareMetalServiceImpl{
		BareMetalRepository: bareMetalRepository,
	}
}

func (s *BareMetalServiceImpl) GetAll(ctx context.Context) ([]dto.BareMetalResponse, error) {
	var responses []dto.BareMetalResponse

	bareMetals, err := s.BareMetalRepository.GetAll(ctx)
	if err != nil {
		return responses, errors.Wrap(err, "failed to find bare metals")
	}

	for _, h := range bareMetals {
		res := dto.BareMetalResponse{
			ID:            h.ID,
			Type:          h.Type,
			Name:          h.Name,
			Url:           h.Url,
			ApiToken:      h.ApiToken,
			BareMetalNode: h.BareMetalNodes,
		}
		responses = append(responses, res)
	}

	return responses, nil
}

func (s *BareMetalServiceImpl) GetOne(ctx context.Context, id uint64) (dto.BareMetalResponse, error) {
	var response dto.BareMetalResponse

	bareMetal, err := s.BareMetalRepository.GetOne(ctx, id)
	if err != nil {
		return response, errors.Wrap(err, "failed to find bare metal")
	}

	response = dto.BareMetalResponse{
		ID:            bareMetal.ID,
		Type:          bareMetal.Type,
		Name:          bareMetal.Name,
		Url:           bareMetal.Url,
		ApiToken:      bareMetal.ApiToken,
		BareMetalNode: bareMetal.BareMetalNodes,
	}

	return response, nil
}

func (s *BareMetalServiceImpl) GetNodeStatusHistory(ctx context.Context, id uint64) ([]dto.BareMetalNodeStatusHistoryResponse, error) {
	var responses []dto.BareMetalNodeStatusHistoryResponse

	bareMetals, err := s.BareMetalRepository.GetNodeStatusHistory(ctx, id)
	if err != nil {
		return responses, errors.Wrap(err, "failed to find bare metals node status history")
	}

	for _, h := range bareMetals {
		res := dto.BareMetalNodeStatusHistoryResponse{
			ID:              h.ID,
			BareMetalNodeID: h.BareMetalNodeID,
			Type:            h.Type,
			Status:          h.Status,
			CreatedAt:       h.CreatedAt,
			UpdatedAt:       h.UpdatedAt,
		}
		responses = append(responses, res)
	}

	return responses, nil
}

func (s *BareMetalServiceImpl) Create(ctx context.Context, request dto.BareMetalRequest) (dto.BareMetalResponse, error) {
	var response dto.BareMetalResponse

	if err := request.Validate(); err != nil {
		return response, err
	}

	bareMetal, err := s.BareMetalRepository.Create(ctx, entity.BareMetal{
		Type:     request.Type,
		Name:     request.Name,
		Url:      request.Url,
		ApiToken: request.ApiToken,
	})

	if err != nil {
		return response, errors.Wrap(err, "failed to create bare metal")
	}

	response = dto.BareMetalResponse{
		ID:       bareMetal.ID,
		Type:     bareMetal.Type,
		Name:     bareMetal.Name,
		Url:      bareMetal.Url,
		ApiToken: bareMetal.ApiToken,
	}

	return response, nil
}

func (s *BareMetalServiceImpl) Update(ctx context.Context, request dto.BareMetalRequest, id uint64) (dto.BareMetalResponse, error) {
	var response dto.BareMetalResponse

	bareMetal, err := s.BareMetalRepository.Update(ctx, entity.BareMetal{
		ID:       id,
		Type:     request.Type,
		Name:     request.Name,
		Url:      request.Url,
		ApiToken: request.ApiToken,
	})

	if err != nil {
		return response, errors.Wrap(err, "failed to update bare metal")
	}

	response = dto.BareMetalResponse{
		ID:       bareMetal.ID,
		Type:     bareMetal.Type,
		Name:     bareMetal.Name,
		Url:      bareMetal.Url,
		ApiToken: bareMetal.ApiToken,
	}

	return response, nil
}

func (s *BareMetalServiceImpl) Delete(ctx context.Context, id uint64) error {
	return s.BareMetalRepository.Delete(ctx, id)
}

func (s *BareMetalServiceImpl) GetAllNodes(ctx context.Context) ([]dto.BareMetalNodeWithStatusResponse, error) {
	var responses []dto.BareMetalNodeWithStatusResponse

	nodes, err := s.BareMetalRepository.GetAllNodes(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find bare metal nodes")
	}

	for _, node := range nodes {
		res := dto.BareMetalNodeWithStatusResponse{
			ID:          node.ID,
			BareMetalID: node.BareMetalID,
			Node:        node.Node,
			Cpu:         node.Cpu,
			Memory:      node.Memory,
			Disk:        node.Disk,
			CreatedAt:   node.CreatedAt,
			UpdatedAt:   node.UpdatedAt,
		}
		// Ambil status terbaru (jika ada)
		if len(node.BareMetalNodeStatus) > 0 {
			latestStatus := node.BareMetalNodeStatus[0]
			res.Status = latestStatus.Status
			res.StatusType = latestStatus.Type
			res.StatusLastUpdatedAt = latestStatus.UpdatedAt
		}
		responses = append(responses, res)
	}

	return responses, nil
}
