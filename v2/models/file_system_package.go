package models

import (
	"github.com/go-openapi/runtime"
)

type FileSystemPackage interface {
	runtime.Validatable
	runtime.ContextValidatable
	Name() string
	SetName(string)
}
type fileSystemPackage struct {
	nameField string
}

func (m *fileSystemPackage) Name() string {
	return m.nameField
}
func (m *fileSystemPackage) SetName(val string) {
	m.nameField = val
}
