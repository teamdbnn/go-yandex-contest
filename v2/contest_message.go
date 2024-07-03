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
		return requiredError("contestID")
	}
	return nil
}

// Do send req
func (s *GetJuryClarifications) Do(ctx context.Context, opts ...RequestOption) (*Clarifications, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%d/clarifications", s.contestID),
	}
	if s.locale == "" {
		r.setParam("locale", "ru")
	} else {
		r.setParam("locale", s.locale)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(Clarifications)

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetContestMessagesService Get messages in contest by contest id
type GetContestMessagesService struct {
	c         *Client
	contestID int64
}

// ContestID Set contest id
func (s *GetContestMessagesService) ContestID(contestID int64) *GetContestMessagesService {
	s.contestID = contestID
	return s
}

func (s *GetContestMessagesService) validate() error {
	if s.contestID == 0 {
		return requiredError("contestID")
	}
	return nil
}

// Do send req
func (s *GetContestMessagesService) Do(ctx context.Context, opts ...RequestOption) (*Messages, error) {
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

	err = json.Unmarshal(data, &res)
	if err != nil {
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
		return requiredError("contestID")
	}
	if s.subject == "" {
		return requiredError("subject")
	}
	if s.text == "" {
		return requiredError("text")
	}
	return nil
}

// Do send req
func (s *SendQuestionToJury) Do(ctx context.Context, opts ...RequestOption) (error, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("/contests/%v/messages", s.contestID),
	}
	if s.problem != "" {
		r.setFormParam("problem", s.problem)
	}
	if s.subject != "" {
		r.setFormParam("subject", s.subject)
	}
	if s.text != "" {
		r.setFormParam("text", s.text)
	}

	_, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
