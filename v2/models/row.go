package models

type Row struct {
	ParticipantInfo *Participant     `json:"participantInfo,omitempty"`
	PlaceFrom       int32            `json:"placeFrom,omitempty"`
	PlaceTo         int32            `json:"placeTo,omitempty"`
	ProblemResults  []*ProblemResult `json:"problemResults"`
	Score           string           `json:"score,omitempty"`
}
