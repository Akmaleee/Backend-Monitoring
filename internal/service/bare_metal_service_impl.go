package service

import (
	"context"
	"it-backend/internal/model/dto"
	"it-backend/internal/model/entity"
	"it-backend/internal/repository"

	"github.com/pkg/errors"
)

type BareMetalRepositoryImpl struct {
	BareMetalRepository repository.BareMetalRepository
}

func NewBareMetalService(bareMetalRepository repository.BareMetalRepository) *BareMetalRepositoryImpl {
	return &BareMetalRepositoryImpl{
		BareMetalRepository: bareMetalRepository,
	}
}

func (s *BareMetalRepositoryImpl) GetAll(ctx context.Context) ([]dto.BareMetalResponse, error) {
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

func (s *BareMetalRepositoryImpl) GetOne(ctx context.Context, id uint64) (dto.BareMetalResponse, error) {
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

func (s *BareMetalRepositoryImpl) GetNodeStatusHistory(ctx context.Context, id uint64) ([]dto.BareMetalNodeStatusHistoryResponse, error) {
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

func (s *BareMetalRepositoryImpl) Create(ctx context.Context, request dto.BareMetalRequest) (dto.BareMetalResponse, error) {
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

func (s *BareMetalRepositoryImpl) Update(ctx context.Context, request dto.BareMetalRequest, id uint64) (dto.BareMetalResponse, error) {
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

func (s *BareMetalRepositoryImpl) Delete(ctx context.Context, id uint64) error {
	return s.BareMetalRepository.Delete(ctx, id)
}
