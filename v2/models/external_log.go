package models

import (
	"github.com/go-openapi/strfmt"
)

type ExternalLog struct {
	ContestLog        *ContestLog     `json:"contestLog,omitempty"`
	ContestLogKey     []string        `json:"contestLogKey"`
	DownloadStartTime strfmt.DateTime `json:"downloadStartTime,omitempty"`
	Dynamic           bool            `json:"dynamic,omitempty"`
	DynamicURL        string          `json:"dynamicUrl,omitempty"`
	EventsCount       int32           `json:"eventsCount,omitempty"`
	ID                *int64          `json:"id"`
	ImportAsInternal  bool            `json:"importAsInternal,omitempty"`
	LastUpdate        strfmt.DateTime `json:"lastUpdate,omitempty"`
	LastUpdateAttempt strfmt.DateTime `json:"lastUpdateAttempt,omitempty"`
	LastUpdateError   []string        `json:"lastUpdateError"`
	OwnerID           []int64         `json:"ownerId"`
	RawXMLLog         strfmt.Base64   `json:"rawXmlLog,omitempty"`
	Title             string          `json:"title,omitempty"`
	XMLLogKey         []string        `json:"xmlLogKey"`
}
