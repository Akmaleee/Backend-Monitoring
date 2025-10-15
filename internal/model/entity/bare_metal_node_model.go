package entity

import "time"

type BareMetalNode struct {
	ID                  uint64                `json:"id" gorm:"primary_key;auto_increment"`
	BareMetalID         uint64                `json:"bare_metal_id" gorm:"column:bare_metal_id;type:bigint unsigned" validate:"required"`
	Node                string                `json:"node" gorm:"column:node;type:varchar(100);not null" validate:"required"`
	Cpu                 uint                  `json:"cpu" gorm:"column:cpu;type:int unsigned;not null"`
	Memory              uint64                `json:"memory" gorm:"column:memory;type:bigint unsigned;not null"`
	Disk                uint64                `json:"disk" gorm:"column:disk;type:bigint unsigned;not null"`
	BareMetalNodeStatus []BareMetalNodeStatus `gorm:"foreignKey:BareMetalNodeID"`
	CreatedAt           time.Time             `json:"-"`
	UpdatedAt           time.Time             `json:"-"`
}

func (*BareMetalNode) TableName() string {
	return "bare_metal_nodes"
}
