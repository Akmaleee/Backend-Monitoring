package entity

import "time"

type VirtualMachineConfig struct {
	ID               uint64    `json:"id" gorm:"primary_key;auto_increment"`
	VirtualMachineID uint64    `json:"virtual_machine_id" gorm:"column:virtual_machine_id;type:bigint unsigned"`
	IsAlertStatus    bool      `json:"is_alert_status" gorm:"column:is_alert_status;type:tinyint;not null" validate:"required"`
	IsAlertDisk      bool      `json:"is_alert_disk" gorm:"column:is_alert_disk;type:tinyint;not null" validate:"required"`
	ThresholdDisk    *float32  `json:"threshold_disk" gorm:"column:threshold_disk;type:float;null"`
	CreatedAt        time.Time `json:"-"`
	UpdatedAt        time.Time `json:"-"`
}

func (*VirtualMachineConfig) TableName() string {
	return "virtual_machine_configs"
}
