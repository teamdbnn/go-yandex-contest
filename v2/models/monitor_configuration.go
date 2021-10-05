package models

type MonitorConfiguration struct {
	ExternalLogs []*ExternalLog `json:"externalLogs"`
	ID           *int64         `json:"id"`
}
