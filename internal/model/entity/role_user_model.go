package entity

import "time"

type RoleUser struct {
	ID        int       `json:"id" gorm:"primary_key;auto_increment"`
	UserID    int       `json:"user_id" gorm:"column:user_id;type:int; null" validate:"required"`
	RoleID    int       `json:"role_id" gorm:"column:role_id;type:int; null" validate:"required"`
	Role      Role      `gorm:"foreignKey:RoleID;references:ID"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (m *RoleUser) TableName() string {
	return "role_user"
}
