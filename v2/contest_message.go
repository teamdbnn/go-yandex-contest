package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetJuryClarifications Get jury clarifications
type GetJuryClarifications struct {
	c         *Client
	contestID int64
	locale    string
}

// ContestID Set contest id
func (s *GetJuryClarifications) ContestID(contestID int64) *GetJuryClarifications {
	s.contestID = contestID
	return s
}

// Locale Set locale
func (s *GetJuryClarifications) Locale(locale string) *GetJuryClarifications {
	s.locale = locale
	return s
}

func (s *GetJuryClarifications) validate() error {
	if s.contestID == 0 {
		return requiredFieldError("contestID")
	}
	return nil
}

// Do Get jury clarifications
//
// Docs: https://api.contest.yandex.net/api/public/swagger-ui.html#/message/getClarificationsUsingGET
// meta:operation GET /contests/{contestId}/clarifications
func (s *GetJuryClarifications) Do(ctx context.Context, opts ...RequestOption) (*Clarifications, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%d/clarifications", s.contestID),
	}
	if s.locale == "" {
		s.locale = defaultLocale
	}

	r.setParam("locale", s.locale)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(Clarifications)
	if err = json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetContestMessages Get messages in contest by contest id
type GetContestMessages struct {
	c         *Client
	contestID int64
}

// ContestID Set contest id
func (s *GetContestMessages) ContestID(contestID int64) *GetContestMessages {
	s.contestID = contestID
	return s
}

func (s *GetContestMessages) validate() error {
	if s.contestID == 0 {
		return requiredFieldError("contestID")
	}
	return nil
}

// Do Get your questions and jury answers
//
// Docs: https://api.contest.yandex.net/api/public/swagger-ui.html#/message/getMessagesUsingGET
// meta:operation GET /contests/{contestId}/messages
func (s *GetContestMessages) Do(ctx context.Context, opts ...RequestOption) (*Messages, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%d/messages", s.contestID),
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(Messages)
	if err = json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// SendQuestionToJury Get jury clarifications in contest by contest id
type SendQuestionToJury struct {
	c         *Client
	contestID int64
	problem   string
	subject   string
	text      string
}

// ContestID Set contest id
func (s *SendQuestionToJury) ContestID(contestID int64) *SendQuestionToJury {
	s.contestID = contestID
	return s
}

// Problem Set problem
func (s *SendQuestionToJury) Problem(problem string) *SendQuestionToJury {
	s.problem = problem
	return s
}

// Subject Set subject
func (s *SendQuestionToJury) Subject(subject string) *SendQuestionToJury {
	s.subject = subject
	return s
}

// Text Set text
func (s *SendQuestionToJury) Text(text string) *SendQuestionToJury {
	s.text = text
	return s
}

func (s *SendQuestionToJury) validate() error {
	if s.contestID == 0 {
		return requiredFieldError("contestID")
	}
	if s.subject == "" {
		return requiredFieldError("subject")
	}
	if s.text == "" {
		return requiredFieldError("text")
	}
	return nil
}

// Do Send question to jury
//
// Docs: https://api.contest.yandex.net/api/public/swagger-ui.html#/message/sendQuestionUsingPOST
// meta:operation POST /contests/{contestId}/messages
func (s *SendQuestionToJury) Do(ctx context.Context, opts ...RequestOption) error {
	if err := s.validate(); err != nil {
		return err
	}

	r := &request{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("/contests/%v/messages", s.contestID),
	}
	r.setFormParam("subject", s.subject)
	r.setFormParam("text", s.text)

	if s.problem != "" {
		r.setFormParam("problem", s.problem)
	}

	_, err := s.c.callAPI(ctx, r, opts...)

	return err
}
