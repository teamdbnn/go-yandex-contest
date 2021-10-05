package models

type AdditionalSolutionCheck struct {
	checkerFilesField FileSystemPackage
	CompilerID        string `json:"compilerId,omitempty"`
	MainFilename      string `json:"mainFilename,omitempty"`
}
