package entity

import "time"

type User struct {
	ID         int        `json:"id" gorm:"primary_key;auto_increment"`
	Name       string     `json:"name" gorm:"column:name;type:varchar(100);not null" validate:"required"`
	Username   string     `json:"username" gorm:"column:username;type:varchar(50);unique;not null" validate:"required"`
	RoleUserID int        `json:"role_user_id" gorm:"column:role_user_id;type:int; null" validate:"required"`
	RoleUser   []RoleUser `gorm:"foreignKey:UserID"`
	IsActive   bool       `json:"is_active" gorm:"column:is_active;type:boolean; null" validate:"required"`
	CreatedAt  time.Time  `json:"-"`
	UpdatedAt  time.Time  `json:"-"`
}

func (m *User) TableName() string {
	return "users"
}
