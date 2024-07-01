package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateGroupForCompetitionService Create group for competition
type CreateGroupForCompetitionService struct {
	c    *Client
	body struct {
		name string
	}
}

// Name Set name
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
func (s *CreateGroupForCompetitionService) Do(ctx context.Context, opts ...RequestOption) (*GroupID, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("/groups/"),
	}

	r.setJSONBody(s.body)
	data, err := s.c.callAPI(ctx, r, opts...)

	res := new(GroupID)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetGroupInfoOfCompetitionService Get group info of competition
type GetGroupInfoOfCompetitionService struct {
	c       *Client
	groupID int64
}

// GroupID Set group id
func (s *GetGroupInfoOfCompetitionService) GroupID(groupID int64) *GetGroupInfoOfCompetitionService {
	s.groupID = groupID
	return s
}

func (s *GetGroupInfoOfCompetitionService) validate() error {
	if s.groupID == 0 {
		return requiredError("groupID")
	}
	return nil
}

// Do send req
func (s *GetGroupInfoOfCompetitionService) Do(ctx context.Context, opts ...RequestOption) (*GroupInfo, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/groups/%v", s.groupID),
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(GroupInfo)

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// AddGroupMemberForCompetitionService Add group member for competition
type AddGroupMemberForCompetitionService struct {
	c       *Client
	groupID int64
	body    struct {
		login string
		uid   int64
	}
}

// GroupID Set group id
func (s *AddGroupMemberForCompetitionService) GroupID(groupID int64) *AddGroupMemberForCompetitionService {
	s.groupID = groupID
	return s
}

// Login Set login
func (s *AddGroupMemberForCompetitionService) Login(login string) *AddGroupMemberForCompetitionService {
	s.body.login = login
	return s
}

// UID Set uid
func (s *AddGroupMemberForCompetitionService) UID(uid int64) *AddGroupMemberForCompetitionService {
	s.body.uid = uid
	return s
}

func (s *AddGroupMemberForCompetitionService) validate() error {
	if s.groupID == 0 {
		return requiredError("groupID")
	}
	if s.body.login == "" {
		return requiredError("login")
	}
	if s.body.uid == 0 {
		return requiredError("uid")
	}
	return nil
}

// Do send req
func (s *AddGroupMemberForCompetitionService) Do(ctx context.Context, opts ...RequestOption) (interface{}, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("/groups/%v/members", s.groupID),
	}

	r.setJSONBody(s.body)
	data, err := s.c.callAPI(ctx, r, opts...)

	fmt.Println(data)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// RemoveGroupOfCompetitionService Remove group of competition
type RemoveGroupOfCompetitionService struct {
	c       *Client
	groupID int64
	body    struct {
		login string
		uid   int64
	}
}

// GroupID Set group id
func (s *RemoveGroupOfCompetitionService) GroupID(groupID int64) *RemoveGroupOfCompetitionService {
	s.groupID = groupID
	return s
}

// Login Set login
func (s *RemoveGroupOfCompetitionService) Login(login string) *RemoveGroupOfCompetitionService {
	s.body.login = login
	return s
}

// UID Set uid
func (s *RemoveGroupOfCompetitionService) UID(uid int64) *RemoveGroupOfCompetitionService {
	s.body.uid = uid
	return s
}

func (s *RemoveGroupOfCompetitionService) validate() error {
	if s.groupID == 0 {
		return requiredError("groupID")
	}
	if s.body.login == "" {
		return requiredError("login")
	}
	if s.body.uid == 0 {
		return requiredError("uid")
	}
	return nil
}

func (s *RemoveGroupOfCompetitionService) Do(ctx context.Context, opts ...RequestOption) (interface{}, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodDelete,
		endpoint: fmt.Sprintf("/groups/%v/members", s.groupID),
	}
	r.setJSONBody(s.body)
	_, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
