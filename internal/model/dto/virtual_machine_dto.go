package dto

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// Request
type VirtualMachineRequest struct {
	ID                   uint64                      `json:"id"`
	BareMetalID          *uint64                     `json:"bare_metal_id"`
	BareMetalNodeID      *uint64                     `json:"bare_metal_node_id"`
	VmID                 string                      `json:"vmid"`
	Code                 string                      `json:"code" validate:"required"`
	Name                 string                      `json:"name" validate:"required"`
	Cpu                  uint                        `json:"cpu"`
	Memory               uint64                      `json:"memory"`
	Disk                 uint64                      `json:"disk"`
	VirtualMachineConfig VirtualMachineConfigRequest `json:"virtual_machine_config,omitempty"`
	VirtualMachineStatus any                         `json:"virtual_machine_status,omitempty"`
	CreatedAt            time.Time                   `json:"-"`
	UpdatedAt            time.Time                   `json:"-"`
}

type VirtualMachineConfigRequest struct {
	VirtualMachineID uint64    `json:"virtual_machine_id"`
	IsAlertStatus    bool      `json:"is_alert_status"`
	IsAlertDisk      bool      `json:"is_alert_disk"`
	ThresholdDisk    *float32  `json:"threshold_disk"`
	CreatedAt        time.Time `json:"-"`
	UpdatedAt        time.Time `json:"-"`
}

type NewVirtualMachineRequest struct {
	VirtualMachine       VirtualMachineRequest       `json:"virtual_machine"`
	VirtualMachineConfig VirtualMachineConfigRequest `json:"virtual_machine_config"`
}

func (l VirtualMachineRequest) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

func (l VirtualMachineConfigRequest) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

func (l NewVirtualMachineRequest) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

// Response
type VirtualMachineResponse struct {
	ID                   uint64    `json:"id"`
	BareMetalID          *uint64   `json:"bare_metal_id"`
	BareMetal            any       `json:"bare_metal"`
	BareMetalNodeID      *uint64   `json:"bare_metal_node_id"`
	BareMetalNode        any       `json:"bare_metal_node"`
	VmID                 string    `json:"vmid"`
	Code                 string    `json:"code"`
	Name                 string    `json:"name"`
	Cpu                  uint      `json:"cpu"`
	Memory               uint64    `json:"memory"`
	Disk                 uint64    `json:"disk"`
	VirtualMachineConfig any       `json:"virtual_machine_config,omitempty"`
	VirtualMachineStatus any       `json:"virtual_machine_status,omitempty"`
	CreatedAt            time.Time `json:"-"`
	UpdatedAt            time.Time `json:"-"`
}

type VirtualMachineStatusHistoryResponse struct {
	ID               uint64    `json:"id"`
	VirtualMachineID uint64    `json:"virtual_machine_id"`
	Type             string    `json:"type"`
	Status           string    `json:"status"`
	CreatedAt        time.Time `json:"-"`
	UpdatedAt        time.Time `json:"-"`
}

type VirtualMachineConfigResponse struct {
	ID               uint64    `json:"id"`
	VirtualMachineID uint64    `json:"virtual_machine_id"`
	IsAlertStatus    bool      `json:"is_alert_status"`
	IsAlertDisk      bool      `json:"is_alert_disk"`
	ThresholdDisk    *float32  `json:"threshold_disk"`
	CreatedAt        time.Time `json:"-"`
	UpdatedAt        time.Time `json:"-"`
}
