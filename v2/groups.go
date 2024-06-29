package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetGroupInfoOfCompetitionService Get group info of competition
type GetGroupInfoOfCompetitionService struct {
	c       *Client
	groupId int64
}

// GroupId Set groupId
func (s *GetGroupInfoOfCompetitionService) GroupId(groupId int64) *GetGroupInfoOfCompetitionService {
	s.groupId = groupId
	return s
}

func (s *GetGroupInfoOfCompetitionService) validate() error {
	if s.groupId == 0 {
		return requiredError("groupId")
	}
	return nil
}

// Do send req
func (s *GetGroupInfoOfCompetitionService) Do(ctx context.Context, opts ...RequestOption) (*Group, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/groups/%v", s.groupId),
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(Group)

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type CreateGroupForCompetitionService struct {
	c    *Client
	body struct {
		name string
	}
}

func (s *CreateGroupForCompetitionService) Name(name string) *CreateGroupForCompetitionService {
	s.body.name = name
	return s
}

func (s *CreateGroupForCompetitionService) validate() error {
	if s.body.name == "" {
		return requiredError("name")
	}
	return nil
}

// Do send req
func (s *CreateGroupForCompetitionService) Do(ctx context.Context, opts ...RequestOption) ([]byte, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("/groups/"),
	}

	r.setJSONBody(s.body)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	return data, nil
}
