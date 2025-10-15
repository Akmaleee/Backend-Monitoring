package entity

import "time"

type Role struct {
	ID        int       `json:"id" gorm:"primary_key;auto_increment"`
	Name      string    `json:"name" gorm:"column:name;type:varchar(100);not null" validate:"required"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func (m *Role) TableName() string {
	return "roles"
}
