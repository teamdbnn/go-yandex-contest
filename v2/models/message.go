package models

type Message struct {
	Answers  []string `json:"answers"`
	Problem  string   `json:"problem,omitempty"`
	Question string   `json:"question,omitempty"`
	Subject  string   `json:"subject,omitempty"`
}
