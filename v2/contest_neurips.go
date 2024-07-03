package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetAllSubmissionsForContest Get all contest submissions
type GetAllSubmissionsForContest struct {
	c                *Client
	contestID        int64
	onlyFinished     bool
	pageNumber       int32
	pageSize         int32
	pageSizeFuncCall bool
	sortBy           string
	sortDesc         string
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
func (s *GetAllSubmissionsForContest) PageNumber(pageNumber int32) *GetAllSubmissionsForContest {
	s.pageNumber = pageNumber
	return s
}

// PageSize Set page size, default 30
func (s *GetAllSubmissionsForContest) PageSize(pageSize int32) *GetAllSubmissionsForContest {
	s.pageSize = pageSize
	s.pageSizeFuncCall = true
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
		return requiredError("contestID")
	}
	return nil
}

// Do send req
func (s *GetAllSubmissionsForContest) Do(ctx context.Context, opts ...RequestOption) (*NeuripsSubmissionsReportResponse, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/neurips/%v/submissions/all", s.contestID),
	}
	const pageSizeDefault int32 = 30
	r.setParam("onlyFinished", s.onlyFinished)
	r.setParam("pageNumber", s.pageNumber)

	if s.pageSize == 0 && !s.pageSizeFuncCall {
		r.setParam("pageSize", pageSizeDefault)
	} else {
		r.setParam("pageSize", s.pageSize)
	}

	if s.sortBy == "" {
		r.setParam("sortBy", "ID")
	} else {
		r.setParam("sortBy", s.sortBy)
	}

	if s.sortDesc == "" {
		r.setParam("sortDesc", "ASC")
	} else {
		r.setParam("sortDesc", s.sortDesc)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(NeuripsSubmissionsReportResponse)

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetYourSubmissionsForContest struct {
	c                *Client
	contestID        int64
	onlyFinished     bool
	pageNumber       int32
	pageSize         int32
	sortBy           string
	sortDesc         string
	pageSizeFuncCall bool
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
func (s *GetYourSubmissionsForContest) PageNumber(pageNumber int32) *GetYourSubmissionsForContest {
	s.pageNumber = pageNumber
	return s
}

// PageSize Set page size, default 30
func (s *GetYourSubmissionsForContest) PageSize(pageSize int32) *GetYourSubmissionsForContest {
	s.pageSize = pageSize
	s.pageSizeFuncCall = true
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
		return requiredError("contestID")
	}
	return nil
}

// Do send req
func (s *GetYourSubmissionsForContest) Do(ctx context.Context, opts ...RequestOption) (*NeuripsSubmissionsReportResponse, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/neurips/%v/submissions/my", s.contestID),
	}
	const pageSizeDefault int32 = 30
	r.setParam("onlyFinished", s.onlyFinished)
	r.setParam("pageNumber", s.pageNumber)

	if s.pageSize == 0 && !s.pageSizeFuncCall {
		r.setParam("pageSize", pageSizeDefault)
	} else {
		r.setParam("pageSize", s.pageSize)
	}

	if s.sortBy == "" {
		r.setParam("sortBy", "ID")
	} else {
		r.setParam("sortBy", s.sortBy)
	}

	if s.sortDesc == "" {
		r.setParam("sortDesc", "ASC")
	} else {
		r.setParam("sortDesc", s.sortDesc)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(NeuripsSubmissionsReportResponse)

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
