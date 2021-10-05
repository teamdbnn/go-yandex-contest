package models

type AccessLevel struct {
	ID                                 *int64  `json:"id"`
	CheckerSettingsModificationAllowed *bool   `json:"checkerSettingsModificationAllowed"`
	ContestLimit                       int32   `json:"contestLimit,omitempty"`
	FileSizeUploadLimit                int64   `json:"fileSizeUploadLimit,omitempty"`
	IsGroupCreationAllowed             *bool   `json:"isGroupCreationAllowed"`
	IsProblemsetModificationAllowed    *bool   `json:"isProblemsetModificationAllowed"`
	Name                               *string `json:"name"`
	ParticipantLimit                   int32   `json:"participantLimit,omitempty"`
	ProblemLimit                       int32   `json:"problemLimit,omitempty"`
	ProblemsetLimit                    int32   `json:"problemsetLimit,omitempty"`
	TestsetTemplateModificationAllowed *bool   `json:"testsetTemplateModificationAllowed"`
	TotalSizeUploadLimit               int64   `json:"totalSizeUploadLimit,omitempty"`
}
