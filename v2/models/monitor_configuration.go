package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// MonitorConfiguration MonitorConfiguration
//
// swagger:model MonitorConfiguration
type MonitorConfiguration struct {

	// external logs
	ExternalLogs []*ExternalLog `json:"externalLogs"`

	// id
	// Required: true
	ID *int64 `json:"id"`
}