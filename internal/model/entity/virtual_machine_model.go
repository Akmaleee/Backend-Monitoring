package entity

import "time"

type VirtualMachine struct {
	ID                   uint64               `json:"id" gorm:"primary_key;auto_increment"`
	BareMetalID          *uint64              `json:"bare_metal_id" gorm:"column:bare_metal_id;type:bigint unsigned" validate:"required"`
	BareMetal            BareMetal            `gorm:"foreignKey:BareMetalID;reference:ID"`
	BareMetalNodeID      *uint64              `json:"bare_metal_node_id" gorm:"column:bare_metal_node_id;type:bigint unsigned"`
	BareMetalNode        BareMetalNode        `gorm:"foreignKey:BareMetalNodeID;reference:ID"`
	VmID                 string               `json:"vmid" gorm:"column:vmid;type:varchar(50);not null" validate:"required"`
	Code                 string               `json:"code" gorm:"column:code;type:varchar(100);null"`
	Name                 string               `json:"name" gorm:"column:name;type:varchar(255);not null" validate:"required"`
	Cpu                  uint                 `json:"cpu" gorm:"column:cpu;type:int unsigned;not null"`
	Memory               uint64               `json:"memory" gorm:"column:memory;type:bigint unsigned;not null"`
	Disk                 uint64               `json:"disk" gorm:"column:disk;type:bigint unsigned;not null"`
	VirtualMachineConfig VirtualMachineConfig `gorm:"foreignKey:VirtualMachineID"`
	VirtualMachineStatus VirtualMachineStatus `gorm:"foreignKey:VirtualMachineID"`
	CreatedAt            time.Time            `json:"-"`
	UpdatedAt            time.Time            `json:"-"`
}

func (*VirtualMachine) TableName() string {
	return "virtual_machines"
}
