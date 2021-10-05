package models

import (
	"io"
)

type FilePackageAllOf1 struct {
	Data io.ReadCloser `json:"data,omitempty"`
	Name string        `json:"name,omitempty"`
}
