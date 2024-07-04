package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateGroupForCompetition Create group for competition
type CreateGroupForCompetition struct {
	c    *Client
	body struct {
		Name string `json:"name"`
	}
}

// Name Set name
func (s *CreateGroupForCompetition) Name(name string) *CreateGroupForCompetition {
	s.body.Name = name
	return s
}

func (s *CreateGroupForCompetition) validate() error {
	if s.body.Name == "" {
		return requiredError("Name")
	}
	return nil
}

// Do Send POST request
func (s *CreateGroupForCompetition) Do(ctx context.Context, opts ...RequestOption) (*GroupID, error) {
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

// GetGroupInfoOfCompetition Get group info of competition
type GetGroupInfoOfCompetition struct {
	c       *Client
	groupID int64
}

// GroupID Set group id
func (s *GetGroupInfoOfCompetition) GroupID(groupID int64) *GetGroupInfoOfCompetition {
	s.groupID = groupID
	return s
}

func (s *GetGroupInfoOfCompetition) validate() error {
	if s.groupID == 0 {
		return requiredError("groupID")
	}
	return nil
}

// Do Send GET request
func (s *GetGroupInfoOfCompetition) Do(ctx context.Context, opts ...RequestOption) (*GroupInfo, error) {
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

// AddGroupMemberForCompetition Add group member for competition
type AddGroupMemberForCompetition struct {
	c       *Client
	groupID int64
	body    struct {
		Login string `json:"login"`
		Uid   int64  `json:"uid"`
	}
}

// GroupID Set group id
func (s *AddGroupMemberForCompetition) GroupID(groupID int64) *AddGroupMemberForCompetition {
	s.groupID = groupID
	return s
}

// Login Set login
func (s *AddGroupMemberForCompetition) Login(Login string) *AddGroupMemberForCompetition {
	s.body.Login = Login
	return s
}

// UID Set uid
func (s *AddGroupMemberForCompetition) UID(Uid int64) *AddGroupMemberForCompetition {
	s.body.Uid = Uid
	return s
}

func (s *AddGroupMemberForCompetition) validate() error {
	if s.groupID == 0 {
		return requiredError("groupID")
	}
	if s.body.Login == "" {
		return requiredError("Login")
	}
	if s.body.Uid == 0 {
		return requiredError("Uid")
	}
	return nil
}

// Do Send POST request
func (s *AddGroupMemberForCompetition) Do(ctx context.Context, opts ...RequestOption) error {
	if err := s.validate(); err != nil {
		return err
	}

	r := &request{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("/groups/%v/members", s.groupID),
	}

	r.setJSONBody(s.body)
	_, err := s.c.callAPI(ctx, r, opts...)

	if err != nil {
		return err
	}

	return nil
}

// RemoveGroupOfCompetition Remove group of competition
type RemoveGroupOfCompetition struct {
	c       *Client
	groupID int64
	body    struct {
		Login string `json:"login"`
		Uid   int64  `json:"uid"`
	}
}

// GroupID Set group id
func (s *RemoveGroupOfCompetition) GroupID(groupID int64) *RemoveGroupOfCompetition {
	s.groupID = groupID
	return s
}

// Login Set login
func (s *RemoveGroupOfCompetition) Login(login string) *RemoveGroupOfCompetition {
	s.body.Login = login
	return s
}

// UID Set uid
func (s *RemoveGroupOfCompetition) UID(uid int64) *RemoveGroupOfCompetition {
	s.body.Uid = uid
	return s
}

func (s *RemoveGroupOfCompetition) validate() error {
	if s.groupID == 0 {
		return requiredError("groupID")
	}
	if s.body.Login == "" {
		return requiredError("Login")
	}
	if s.body.Uid == 0 {
		return requiredError("Uid")
	}
	return nil
}

// Do Send DELETE request
func (s *RemoveGroupOfCompetition) Do(ctx context.Context, opts ...RequestOption) error {
	if err := s.validate(); err != nil {
		return err
	}

	r := &request{
		method:   http.MethodDelete,
		endpoint: fmt.Sprintf("/groups/%v/members", s.groupID),
	}
	r.setJSONBody(s.body)
	_, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return nil
}
