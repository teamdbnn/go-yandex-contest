package models

type ContestStandings struct {
	Rows   []*Row      `json:"rows"`
	Titles []*RowTitle `json:"titles"`
}
