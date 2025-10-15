package entity

import "time"

type VirtualMachineStatusHistory struct {
	ID               uint64    `json:"id" gorm:"primary_key;auto_increment"`
	VirtualMachineID uint64    `json:"virtual_machine_id" gorm:"column:virtual_machine_id;type:bigint unsigned"`
	Type             string    `json:"type" gorm:"column:type;type:varchar(50);not null" validate:"required"`
	Status           string    `json:"status" gorm:"column:status;type:varchar(50);null"`
	CreatedAt        time.Time `json:"-"`
	UpdatedAt        time.Time `json:"-"`
}

func (*VirtualMachineStatusHistory) TableName() string {
	return "virtual_machine_status_histories"
}
