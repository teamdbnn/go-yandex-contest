package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetContestProblems Get contest problems
type GetContestProblems struct {
	c         *Client
	contestId int64
	locale    string
}

// ContestID Set contest id
func (s *GetContestProblems) ContestID(contestId int64) *GetContestProblems {
	s.contestId = contestId
	return s
}

// Locale Set locale
func (s *GetContestProblems) Locale(locale string) *GetContestProblems {
	s.locale = locale
	return s
}

func (s *GetContestProblems) validate() error {
	if s.contestId == 0 {
		return requiredError("contest")
	}
	return nil
}

// Do send req
func (s *GetContestProblems) Do(ctx context.Context, opts ...RequestOption) (*ContestProblems, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/problems", s.contestId),
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

// GetStatementProblem Get statement problem
type GetStatementProblem struct {
	c         *Client
	contestId int64
	alias     string
	locale    string
	types     string
}

// ContestID Set contest id
func (s *GetStatementProblem) ContestID(contestId int64) *GetStatementProblem {
	s.contestId = contestId
	return s
}

// Alias Set alias
func (s *GetStatementProblem) Alias(alias string) *GetStatementProblem {
	s.alias = alias
	return s
}

// Locale Set locale
func (s *GetStatementProblem) Locale(locale string) *GetStatementProblem {
	s.locale = locale
	return s
}

// Type Set type
func (s *GetStatementProblem) Type(types string) *GetStatementProblem {
	s.types = types
	return s
}

func (s *GetStatementProblem) validate() error {
	if s.contestId == 0 {
		return requiredError("contestId")
	}
	if s.alias == "" {
		return requiredError("alias")
	}
	return nil
}

// Do Send GET request
func (s *GetStatementProblem) Do(ctx context.Context, opts ...RequestOption) (*ContestProblems, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/problems/%v/statement", s.contestId, s.alias),
	}
	if s.locale == "" {
		r.setParam("locale", "ru")
	} else {
		r.setParam("locale", s.locale)
	}
	r.setParam("type", s.types)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(ContestProblems)
	err = json.Unmarshal(data, res)
	fmt.Println(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetProblemFile Get Problem file
type GetProblemFile struct {
	c         *Client
	path      string
	ProblemId string
}

// Path Set path
func (s *GetProblemFile) Path(path string) *GetProblemFile {
	s.path = path
	return s
}

// ProblemID Set problem id
func (s *GetProblemFile) ProblemID(ProblemId string) *GetProblemFile {
	s.ProblemId = ProblemId
	return s
}

func (s *GetProblemFile) validate() error {
	if s.path == "" {
		return requiredError("path")
	}
	if s.ProblemId == "" {
		return requiredError("ProblemId")
	}
	return nil
}

// Do send req          todo: File ouput?
func (s *GetProblemFile) Do(ctx context.Context, opts ...RequestOption) (interface{}, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/problems"),
	}

	r.setParam("path", s.path)
	r.setParam("problemId", s.ProblemId)
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
