package dto

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// Request
type BareMetalRequest struct {
	Type      string    `json:"type" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	Url       string    `json:"url" validate:"required"`
	ApiToken  string    `json:"api_token" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BareMetalNodeRequest struct {
	BareMetalID uint64    `json:"bare_metal_id"`
	Node        string    `json:"node"`
	Cpu         uint      `json:"cpu"`
	Memory      uint64    `json:"memory"`
	Disk        uint64    `json:"disk"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BareMetalNodeStatusRequest struct {
	BareMetalNodeID uint64    `json:"bare_metal_node_id"`
	Type            string    `json:"type"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (l BareMetalRequest) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type NewBareMetalRequest struct {
	BareMetal BareMetalRequest           `json:"bare_metal"`
	Node      BareMetalNodeRequest       `json:"node"`
	Status    BareMetalNodeStatusRequest `json:"status"`
}

// Response
type NewBareMetalResponse struct {
	BareMetal any `json:"bare_metal"`
	Node      any `json:"bare_metal_node,omitempty"`
	Status    any `json:"bare_metal_node_status,omitempty"`
}

type BareMetalResponse struct {
	ID            uint64 `json:"id"`
	Type          string `json:"type"`
	Name          string `json:"name"`
	Url           string `json:"url"`
	ApiToken      string `json:"api_token"`
	BareMetalNode any    `json:"bare_metal_node,omitempty"`
}

type BareMetalNodeResponse struct {
	ID          uint64    `json:"id"`
	BareMetalID uint64    `json:"bare_metal_id"`
	Node        string    `json:"node"`
	Cpu         uint      `json:"cpu"`
	Memory      uint64    `json:"memory"`
	Disk        uint64    `json:"disk"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

type BareMetalNodeStatusResponse struct {
	ID              uint64    `json:"id"`
	BareMetalNodeID uint64    `json:"bare_metal_node_id"`
	Type            string    `json:"type"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"-"`
	UpdatedAt       time.Time `json:"-"`
}

type BareMetalNodeStatusHistoryResponse struct {
	ID              uint64    `json:"id"`
	BareMetalNodeID uint64    `json:"bare_metal_node_id"`
	Type            string    `json:"type"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"-"`
}

type BareMetalNodeWithStatusResponse struct {
	ID                  uint64    `json:"id"`
	BareMetalID         uint64    `json:"bare_metal_id"`
	Node                string    `json:"node"`
	Cpu                 uint      `json:"cpu"`
	Memory              uint64    `json:"memory"`
	Disk                uint64    `json:"disk"`
	Status              string    `json:"status"`
	StatusType          string    `json:"status_type"`
	StatusLastUpdatedAt time.Time `json:"status_last_updated_at"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
