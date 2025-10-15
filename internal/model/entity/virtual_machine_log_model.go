package entity

import "time"

type VirtualMachineLog struct {
	ID               uint64    `json:"id" gorm:"primary_key;auto_increment"`
	VirtualMachineID uint64    `json:"virtual_machine_id" gorm:"column:virtual_machine_id;type:bigint unsigned"`
	Log              string    `json:"log" gorm:"column:log;type:longtext;null"`
	CreatedAt        time.Time `json:"-"`
	UpdatedAt        time.Time `json:"-"`
}

func (*VirtualMachineLog) TableName() string {
	return "virtual_machine_logs"
}
