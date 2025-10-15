package repository

import (
	"context"
	"it-backend/database"
	"it-backend/internal/model/entity"

	"gorm.io/gorm"
)

type VirtualMachineRepositoryImpl struct {
	DB database.DatabaseMySQL
}

func NewVirtualMachineRepository(db database.DatabaseMySQL) *VirtualMachineRepositoryImpl {
	return &VirtualMachineRepositoryImpl{
		DB: db,
	}
}

func (r *VirtualMachineRepositoryImpl) GetAll(ctx context.Context) ([]entity.VirtualMachine, error) {
	var (
		VirtualMachine []entity.VirtualMachine
		Total          int64
		err            error
	)

	err = r.DB.DBInfra.Model(&VirtualMachine).Order("created_at desc").Count(&Total).Error
	if err != nil {
		return VirtualMachine, err
	}

	err = r.DB.DBInfra.
		Model(&entity.VirtualMachine{}).
		Preload("BareMetal").
		Preload("BareMetalNode.BareMetalNodeStatus").
		Preload("VirtualMachineStatus").
		Order("created_at desc").
		Find(&VirtualMachine).Error
	if err != nil {
		return VirtualMachine, err
	}

	return VirtualMachine, nil
}

func (r *VirtualMachineRepositoryImpl) GetOne(ctx context.Context, id uint64) (entity.VirtualMachine, error) {
	var (
		VirtualMachine entity.VirtualMachine
		err            error
	)

	err = r.DB.DBInfra.Model(&entity.VirtualMachine{}).
		Preload("BareMetal").
		Preload("BareMetalNode.BareMetalNodeStatus").
		Preload("VirtualMachineConfig").
		Preload("VirtualMachineStatus").
		Where("id = ?", id).
		First(&VirtualMachine).Error

	if err != nil {
		return VirtualMachine, err
	}

	return VirtualMachine, nil
}

func (r *VirtualMachineRepositoryImpl) GetStatusHistory(ctx context.Context, id uint64) ([]entity.VirtualMachineStatusHistory, error) {
	var (
		VirtualMachineStatusHistory []entity.VirtualMachineStatusHistory
		err                         error
	)

	err = r.DB.DBInfra.Where("virtual_machine_id = ?", id).Find(&VirtualMachineStatusHistory).Error

	if err != nil {
		return VirtualMachineStatusHistory, err
	}

	return VirtualMachineStatusHistory, nil
}

func (r *VirtualMachineRepositoryImpl) Create(ctx context.Context, virtualMachine entity.VirtualMachine, virtualMachineConfig entity.VirtualMachineConfig) (entity.VirtualMachine, error) {
	err := r.DB.DBInfra.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&virtualMachine).Error; err != nil {
			return err
		}

		virtualMachineConfig.VirtualMachineID = virtualMachine.ID
		if err := tx.Model(&entity.VirtualMachineConfig{}).
			Create(&virtualMachineConfig).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return virtualMachine, err
	}

	err = r.DB.DBInfra.Model(&entity.VirtualMachine{}).Preload("VirtualMachineConfig").Where("id = ?", virtualMachine.ID).First(&virtualMachine).Error

	if err != nil {
		return virtualMachine, err
	}

	return virtualMachine, nil
}

func (r *VirtualMachineRepositoryImpl) Update(ctx context.Context, virtualMachine entity.VirtualMachine) (entity.VirtualMachine, error) {
	err := r.DB.DBInfra.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&entity.VirtualMachine{}).
			Where("id = ?", virtualMachine.ID).
			Updates(virtualMachine).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return virtualMachine, err
	}

	return virtualMachine, nil
}

func (r *VirtualMachineRepositoryImpl) UpdateConfig(ctx context.Context, virtualMachineConfig entity.VirtualMachineConfig) (entity.VirtualMachineConfig, error) {
	if err := r.DB.DBInfra.Select("virtual_machine_id", "is_alert_status", "is_alert_disk", "threshold_disk").Updates(virtualMachineConfig).Where("id = ?", virtualMachineConfig.ID).Error; err != nil {
		return virtualMachineConfig, err
	}

	return virtualMachineConfig, nil
}

func (r *VirtualMachineRepositoryImpl) Delete(ctx context.Context, id uint64) error {
	err := r.DB.DBInfra.Delete(&entity.VirtualMachine{}, id).Error

	if err != nil {
		return err
	}

	return nil
}
