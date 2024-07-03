package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// GetContestProblems Get contest problems
type GetContestProblems struct {
	c         *Client
	contestID int64
	locale    string
}

// ContestID Set contest id
func (s *GetContestProblems) ContestID(contestID int64) *GetContestProblems {
	s.contestID = contestID
	return s
}

// Locale Set locale
func (s *GetContestProblems) Locale(locale string) *GetContestProblems {
	s.locale = locale
	return s
}

func (s *GetContestProblems) validate() error {
	if s.contestID == 0 {
		return requiredError("contestID")
	}
	return nil
}

// Do send req
func (s *GetContestProblems) Do(ctx context.Context, opts ...RequestOption) (*ContestProblems, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/problems", s.contestID),
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
	res := new(ContestProblems)
	err = json.Unmarshal(data, &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetProblemStatement Get problem statement
type GetProblemStatement struct {
	c         *Client
	contestID int64
	alias     string
	locale    string
	types     string
}

// ContestID Set contest id
func (s *GetProblemStatement) ContestID(contestID int64) *GetProblemStatement {
	s.contestID = contestID
	return s
}

// Alias Set alias
func (s *GetProblemStatement) Alias(alias string) *GetProblemStatement {
	s.alias = alias
	return s
}

// Locale Set locale
func (s *GetProblemStatement) Locale(locale string) *GetProblemStatement {
	s.locale = locale
	return s
}

// Type Set type
func (s *GetProblemStatement) Type(types string) *GetProblemStatement {
	s.types = types
	return s
}

func (s *GetProblemStatement) validate() error {
	if s.contestID == 0 {
		return requiredError("contestID")
	}
	if s.alias == "" {
		return requiredError("alias")
	}
	return nil
}

// Do Send GET request
func (s *GetProblemStatement) Do(ctx context.Context, opts ...RequestOption) error {
	if err := s.validate(); err != nil {
		return err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/problems/%v/statement", s.contestID, s.alias),
	}
	if s.locale == "" {
		r.setParam("locale", "ru")
	} else {
		r.setParam("locale", s.locale)
	}
	r.setParam("type", s.types)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}

	f, err := os.Create("problemStatement")
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}

// GetProblemFile Get Problem file
type GetProblemFile struct {
	c         *Client
	path      string
	problemID string
}

// Path Set path
func (s *GetProblemFile) Path(path string) *GetProblemFile {
	s.path = path
	return s
}

// ProblemID Set problem id
func (s *GetProblemFile) ProblemID(problemID string) *GetProblemFile {
	s.problemID = problemID
	return s
}

func (s *GetProblemFile) validate() error {
	if s.path == "" {
		return requiredError("path")
	}
	if s.problemID == "" {
		return requiredError("ProblemID")
	}
	return nil
}

// Do send req          todo: File ouput?
func (s *GetProblemFile) Do(ctx context.Context, opts ...RequestOption) (interface{}, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/problems"),
	}

	r.setParam("path", s.path)
	r.setParam("problemId", s.problemID)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(interface{})
	err = json.Unmarshal(data, res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
