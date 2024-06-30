package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetContestStandings Get contest standings
type GetContestStandings struct {
	c                 *Client
	contestId         int64
	forJudge          bool
	locale            string
	page              int32
	pageFuncCall      bool
	pageSize          int32
	pageSizeFuncCall  bool
	participantSearch string
	showExternal      bool
	showVirtual       bool
	userGroupId       int64
}

// ContestID Set contest id
func (s *GetContestStandings) ContestID(contestId int64) *GetContestStandings {
	s.contestId = contestId
	return s
}

// ForJudge for judge or not, default false
func (s *GetContestStandings) ForJudge(forJudge bool) *GetContestStandings {
	s.forJudge = forJudge
	return s
}

// Locale Set locale, default "ru"
func (s *GetContestStandings) Locale(locale string) *GetContestStandings {
	s.locale = locale
	return s
}

// Page Set page, default 1
func (s *GetContestStandings) Page(page int32) *GetContestStandings {
	s.page = page
	s.pageFuncCall = true
	return s
}

// PageSize Set page size, default 100
func (s *GetContestStandings) PageSize(pageSize int32) *GetContestStandings {
	s.pageSize = pageSize
	s.pageSizeFuncCall = true
	return s
}

// ParticipantSearch Set participant search
func (s *GetContestStandings) ParticipantSearch(participantSearch string) *GetContestStandings {
	s.participantSearch = participantSearch
	return s
}

// ShowExternal Show external, default false
func (s *GetContestStandings) ShowExternal(showExternal bool) *GetContestStandings {
	s.showExternal = showExternal
	return s
}

// ShowVirtual Show virtual, default false
func (s *GetContestStandings) ShowVirtual(showVirtual bool) *GetContestStandings {
	s.showVirtual = showVirtual
	return s
}

// UserGroupId User group id
func (s *GetContestStandings) UserGroupId(userGroupId int64) *GetContestStandings {
	s.userGroupId = userGroupId
	return s
}

func (s *GetContestStandings) validate() error {
	if s.contestId == 0 {
		return requiredError("contestId")
	}
	return nil
}

// Do send req
func (s *GetContestStandings) Do(ctx context.Context, opts ...RequestOption) (*ContestStandings, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/standings", s.contestId),
	}
	r.setParam("forJudge", s.forJudge)

	if s.locale == "" {
		r.setParam("locale", "ru")
	} else {
		r.setParam("locale", s.locale)
	}

	if s.page == 0 && !s.pageFuncCall {
		r.setParam("page", 1)
	} else {
		r.setParam("page", s.page)
	}

	if s.pageSize == 0 && !s.pageSizeFuncCall {
		r.setParam("pageSize", 100)
	} else {
		r.setParam("pageSize", s.pageSize)
	}

	r.setParam("showExternal", s.showExternal)

	r.setParam("showVirtual", s.showVirtual)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(ContestStandings)

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetContestStandingsExtended Get extended contest standings
type GetContestStandingsExtended struct {
	c                 *Client
	contestId         int64
	forJudge          bool
	locale            string
	page              int32
	pageFuncCall      bool
	pageSize          int32
	pageSizeFuncCall  bool
	participantSearch string
	showExternal      bool
	showVirtual       bool
	userGroupId       int64
}

// ContestID Set contest id
func (s *GetContestStandingsExtended) ContestID(contestId int64) *GetContestStandingsExtended {
	s.contestId = contestId
	return s
}

// ForJudge for judge or not, default false
func (s *GetContestStandingsExtended) ForJudge(forJudge bool) *GetContestStandingsExtended {
	s.forJudge = forJudge
	return s
}

// Locale Set locale, default "ru"
func (s *GetContestStandingsExtended) Locale(locale string) *GetContestStandingsExtended {
	s.locale = locale
	return s
}

// Page Set page, default 1
func (s *GetContestStandingsExtended) Page(page int32) *GetContestStandingsExtended {
	s.page = page
	s.pageFuncCall = true
	return s
}

// PageSize Set page size, default 100
func (s *GetContestStandingsExtended) PageSize(pageSize int32) *GetContestStandingsExtended {
	s.pageSize = pageSize
	s.pageSizeFuncCall = true
	return s
}

// ParticipantSearch Set participant search
func (s *GetContestStandingsExtended) ParticipantSearch(participantSearch string) *GetContestStandingsExtended {
	s.participantSearch = participantSearch
	return s
}

// ShowExternal Show external, default false
func (s *GetContestStandingsExtended) ShowExternal(showExternal bool) *GetContestStandingsExtended {
	s.showExternal = showExternal
	return s
}

// ShowVirtual Show virtual, default false
func (s *GetContestStandingsExtended) ShowVirtual(showVirtual bool) *GetContestStandingsExtended {
	s.showVirtual = showVirtual
	return s
}

// UserGroupId User group id
func (s *GetContestStandingsExtended) UserGroupId(userGroupId int64) *GetContestStandingsExtended {
	s.userGroupId = userGroupId
	return s
}

func (s *GetContestStandingsExtended) validate() error {
	if s.contestId == 0 {
		return requiredError("contestId")
	}
	return nil
}

// Do send req
func (s *GetContestStandingsExtended) Do(ctx context.Context, opts ...RequestOption) (*ContestStandings, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/standings-extended", s.contestId),
	}
	r.setParam("forJudge", s.forJudge)

	if s.locale == "" {
		r.setParam("locale", "ru")
	} else {
		r.setParam("locale", s.locale)
	}

	if s.page == 0 && !s.pageFuncCall {
		r.setParam("page", 1)
	} else {
		r.setParam("page", s.page)
	}

	if s.pageSize == 0 && !s.pageSizeFuncCall {
		r.setParam("pageSize", 100)
	} else {
		r.setParam("pageSize", s.pageSize)
	}

	r.setParam("showExternal", s.showExternal)

	r.setParam("showVirtual", s.showVirtual)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(ContestStandings)

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type GetYourPositionInContestStandings struct {
	c            *Client
	contestId    int64
	forJudge     bool
	locale       string
	showExternal bool
	showVirtual  bool
}

// ContestID Set contest id
func (s *GetYourPositionInContestStandings) ContestID(contestId int64) *GetYourPositionInContestStandings {
	s.contestId = contestId
	return s
}

// ForJudge for judge or not, default false
func (s *GetYourPositionInContestStandings) ForJudge(forJudge bool) *GetYourPositionInContestStandings {
	s.forJudge = forJudge
	return s
}

// Locale Set locale, default "ru"
func (s *GetYourPositionInContestStandings) Locale(locale string) *GetYourPositionInContestStandings {
	s.locale = locale
	return s
}

// ShowExternal Show external, default false
func (s *GetYourPositionInContestStandings) ShowExternal(showExternal bool) *GetYourPositionInContestStandings {
	s.showExternal = showExternal
	return s
}

// ShowVirtual Show virtual, default false
func (s *GetYourPositionInContestStandings) ShowVirtual(showVirtual bool) *GetYourPositionInContestStandings {
	s.showVirtual = showVirtual
	return s
}

func (s *GetYourPositionInContestStandings) validate() error {
	if s.contestId == 0 {
		return requiredError("contest")
	}
	return nil
}

// Do send req
func (s *GetYourPositionInContestStandings) Do(ctx context.Context, opts ...RequestOption) (*ContestStandings, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/standings/my", s.contestId),
	}
	r.setParam("forJudge", s.forJudge)
	r.setParam("showExternal", s.showExternal)
	r.setParam("showVirtual", s.showVirtual)
	if s.locale != "" {
		r.setParam("locale", s.locale)
	} else {
		r.setParam("locale", "ru")
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(ContestStandings)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
