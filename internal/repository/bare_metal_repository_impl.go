package repository

import (
	"context"
	"it-backend/database"
	"it-backend/internal/model/entity"

	"gorm.io/gorm"
)

type BareMetalRepositoryImpl struct {
	DB database.DatabaseMySQL
}

func NewBareMetalRepository(db database.DatabaseMySQL) *BareMetalRepositoryImpl {
	return &BareMetalRepositoryImpl{
		DB: db,
	}
}

func (r *BareMetalRepositoryImpl) GetAll(ctx context.Context) ([]entity.BareMetal, error) {
	var (
		BareMetal []entity.BareMetal
		Total     int64
		err       error
	)

	err = r.DB.DBInfra.Model(&BareMetal).Order("created_at desc").Count(&Total).Error
	if err != nil {
		return BareMetal, err
	}

	err = r.DB.DBInfra.
		Model(&entity.BareMetal{}).
		Preload("BareMetalNodes.BareMetalNodeStatus").
		Order("created_at desc").
		Find(&BareMetal).Error
	if err != nil {
		return BareMetal, err
	}

	return BareMetal, nil
}

func (r *BareMetalRepositoryImpl) GetOne(ctx context.Context, id uint64) (entity.BareMetal, error) {
	var (
		BareMetal entity.BareMetal
		err       error
	)

	err = r.DB.DBInfra.Model(&entity.BareMetal{}).Preload("BareMetalNodes.BareMetalNodeStatus").Where("id = ?", id).First(&BareMetal).Error

	if err != nil {
		return BareMetal, err
	}

	return BareMetal, nil
}

func (r *BareMetalRepositoryImpl) GetNodeStatusHistory(ctx context.Context, id uint64) ([]entity.BareMetalNodeStatusHistory, error) {
	var (
		BareMetalNodeStatusHistory []entity.BareMetalNodeStatusHistory
		err                        error
	)

	err = r.DB.DBInfra.Where("bare_metal_node_id = ?", id).Find(&BareMetalNodeStatusHistory).Error

	if err != nil {
		return BareMetalNodeStatusHistory, err
	}

	return BareMetalNodeStatusHistory, nil
}

func (r *BareMetalRepositoryImpl) Create(ctx context.Context, bareMetal entity.BareMetal) (entity.BareMetal, error) {
	err := r.DB.DBInfra.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&bareMetal).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return bareMetal, err
	}

	return bareMetal, nil
}

func (r *BareMetalRepositoryImpl) Update(ctx context.Context, bareMetal entity.BareMetal) (entity.BareMetal, error) {
	err := r.DB.DBInfra.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&entity.BareMetal{}).
			Where("id = ?", bareMetal.ID).
			Updates(bareMetal).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return bareMetal, err
	}

	return bareMetal, nil
}

func (r *BareMetalRepositoryImpl) Delete(ctx context.Context, id uint64) error {
	err := r.DB.DBInfra.Delete(&entity.BareMetal{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
