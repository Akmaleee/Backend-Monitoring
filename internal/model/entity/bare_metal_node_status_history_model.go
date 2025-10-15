package entity

import "time"

type BareMetalNodeStatusHistory struct {
	ID              uint64    `json:"id" gorm:"primary_key;auto_increment"`
	BareMetalNodeID uint64    `json:"bare_metal_node_id" gorm:"column:bare_metal_node_id;type:bigint unsigned"`
	Type            string    `json:"type" gorm:"column:type;type:varchar(50);not null" validate:"required"`
	Status          string    `json:"status" gorm:"column:status;type:varchar(50);null"`
	CreatedAt       time.Time `json:"-"`
	UpdatedAt       time.Time `json:"-"`
}

func (*BareMetalNodeStatusHistory) TableName() string {
	return "bare_metal_node_status_histories"
}
