package entity

import "time"

type BareMetal struct {
	ID             uint64          `json:"id" gorm:"primary_key;auto_increment"`
	Type           string          `json:"type" gorm:"column:type;type:varchar(50);not null" validate:"required"`
	Name           string          `json:"name" gorm:"column:name;type:varchar(100);not null" validate:"required"`
	Url            string          `json:"url" gorm:"column:url;type:text" validate:"required"`
	ApiToken       string          `json:"api_token" gorm:"column:api_token;type:varchar(255);not null;" validate:"required"`
	BareMetalNodes []BareMetalNode `gorm:"foreignKey:BareMetalID"`
	CreatedAt      time.Time       `json:"-"`
	UpdatedAt      time.Time       `json:"-"`
}

func (*BareMetal) TableName() string {
	return "bare_metals"
}
