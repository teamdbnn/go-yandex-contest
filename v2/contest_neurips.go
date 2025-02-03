package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	defaultSubmissionPageSize  = 30
	defaultSubmissionSortField = "ID"
	defaultSubmissionSortDesc  = "ASC"
)

// GetAllSubmissionsForContest Get all submissions for contest
type GetAllSubmissionsForContest struct {
	c            *Client
	contestID    int64
	onlyFinished bool
	pageNumber   int64
	pageSize     int64
	sortBy       string
	sortDesc     string
}

// ContestID Set contest id
func (s *GetAllSubmissionsForContest) ContestID(contestID int64) *GetAllSubmissionsForContest {
	s.contestID = contestID
	return s
}

// OnlyFinished only finished or not, default false
func (s *GetAllSubmissionsForContest) OnlyFinished(onlyFinished bool) *GetAllSubmissionsForContest {
	s.onlyFinished = onlyFinished
	return s
}

// PageNumber Set page number, default 0
func (s *GetAllSubmissionsForContest) PageNumber(pageNumber int64) *GetAllSubmissionsForContest {
	s.pageNumber = pageNumber
	return s
}

// PageSize Set page size, default 30
func (s *GetAllSubmissionsForContest) PageSize(pageSize int64) *GetAllSubmissionsForContest {
	s.pageSize = pageSize
	return s
}

// Sort by field and direction, default "ID" and "ASC"
func (s *GetAllSubmissionsForContest) Sort(field, direction string) *GetAllSubmissionsForContest {
	s.sortBy = field
	s.sortDesc = direction
	return s
}

func (s *GetAllSubmissionsForContest) validate() error {
	if s.contestID == 0 {
		return requiredFieldError("contestID")
	}
	return nil
}

// Do Get all submissions for contest
//
// Docs: https://api.contest.yandex.net/api/public/swagger-ui.html#/neurips/getAllNeuripsSubmissionsReportsUsingGET
// meta:operation GET /contests/neurips/{contestId}/submissions/all
func (s *GetAllSubmissionsForContest) Do(ctx context.Context, opts ...RequestOption) (*NeuripsSubmissionsReportResponse, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/neurips/%v/submissions/all", s.contestID),
	}

	if s.pageSize <= 0 {
		s.pageSize = defaultSubmissionPageSize
	}
	if s.sortBy == "" {
		s.sortBy = defaultSubmissionSortField
	}
	if s.sortDesc == "" {
		s.sortDesc = defaultSubmissionSortDesc
	}
	if s.pageNumber <= 0 {
		s.pageNumber = 1
	}

	r.setParam("onlyFinished", s.onlyFinished)
	r.setParam("pageNumber", s.pageNumber)
	r.setParam("pageSize", s.pageSize)
	r.setParam("sortBy", s.sortBy)
	r.setParam("sortDesc", s.sortDesc)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(NeuripsSubmissionsReportResponse)
	if err = json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetYourSubmissionsForContest Get your contest submissions
type GetYourSubmissionsForContest struct {
	c            *Client
	contestID    int64
	onlyFinished bool
	pageNumber   int64
	pageSize     int64
	sortBy       string
	sortDesc     string
}

// ContestID Set contest id
func (s *GetYourSubmissionsForContest) ContestID(contestID int64) *GetYourSubmissionsForContest {
	s.contestID = contestID
	return s
}

// OnlyFinished only finished or not, default false
func (s *GetYourSubmissionsForContest) OnlyFinished(onlyFinished bool) *GetYourSubmissionsForContest {
	s.onlyFinished = onlyFinished
	return s
}

// PageNumber Set page number, default 0
func (s *GetYourSubmissionsForContest) PageNumber(pageNumber int64) *GetYourSubmissionsForContest {
	s.pageNumber = pageNumber
	return s
}

// PageSize Set page size, default 30
func (s *GetYourSubmissionsForContest) PageSize(pageSize int64) *GetYourSubmissionsForContest {
	s.pageSize = pageSize
	return s
}

// SortBy Sort by, default "ID"
func (s *GetYourSubmissionsForContest) SortBy(sortBy string) *GetYourSubmissionsForContest {
	s.sortBy = sortBy
	return s
}

// SortDesc Sort desc, default "ASC"
func (s *GetYourSubmissionsForContest) SortDesc(sortDesc string) *GetYourSubmissionsForContest {
	s.sortDesc = sortDesc
	return s
}

func (s *GetYourSubmissionsForContest) validate() error {
	if s.contestID == 0 {
		return requiredFieldError("contestID")
	}
	return nil
}

// Do Get your submissions for contest
//
// Docs: https://api.contest.yandex.net/api/public/swagger-ui.html#/neurips/getMyNeuripsSubmissionsReportsUsingGET
// meta:operation GET /contests/neurips/{contestId}/submissions/my
func (s *GetYourSubmissionsForContest) Do(ctx context.Context, opts ...RequestOption) (*NeuripsSubmissionsReportResponse, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/neurips/%v/submissions/my", s.contestID),
	}
	if s.pageSize <= 0 {
		s.pageSize = defaultSubmissionPageSize
	}
	if s.sortBy == "" {
		s.sortBy = defaultSubmissionSortField
	}
	if s.sortDesc == "" {
		s.sortDesc = defaultSubmissionSortDesc
	}
	if s.pageNumber <= 0 {
		s.pageNumber = 1
	}

	r.setParam("onlyFinished", s.onlyFinished)
	r.setParam("pageNumber", s.pageNumber)
	r.setParam("pageSize", s.pageSize)
	r.setParam("sortBy", s.sortBy)
	r.setParam("sortDesc", s.sortDesc)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(NeuripsSubmissionsReportResponse)
	if err = json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}
