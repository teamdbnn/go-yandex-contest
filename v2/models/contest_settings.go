package models

type ContestSettings struct {

	// contest name
	ContestName string `json:"contestName,omitempty"`

	// contest type
	// Enum: [SIMPLE TCM]
	ContestType string `json:"contestType,omitempty"`

	// duration
	Duration string `json:"duration,omitempty"`

	// languages
	Languages []Language `json:"languages"`
}
