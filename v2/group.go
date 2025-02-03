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
		return requiredFieldError("Name")
	}
	return nil
}

// Do Create a new group
// Docs: https://api.contest.yandex.net/api/public/swagger-ui.html#/group/createGroupUsingPOST
// meta:operation POST /groups/
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
	if err = json.Unmarshal(data, res); err != nil {
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
		return requiredFieldError("groupID")
	}
	return nil
}

// Do View group
// Docs: https://api.contest.yandex.net/api/public/swagger-ui.html#/group/getGroupUsingGET
// meta:operation GET /groups/{groupId}
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
	if err = json.Unmarshal(data, res); err != nil {
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
		return requiredFieldError("groupID")
	}
	if s.body.Login == "" {
		return requiredFieldError("Login")
	}
	if s.body.Uid == 0 {
		return requiredFieldError("Uid")
	}
	return nil
}

// Do Add group member
// Docs: https://api.contest.yandex.net/api/public/swagger-ui.html#/group/addMemberUsingPOST
// meta:operation POST /groups/{groupId}/members
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

	return err
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
		return requiredFieldError("groupID")
	}
	if s.body.Login == "" {
		return requiredFieldError("Login")
	}
	if s.body.Uid == 0 {
		return requiredFieldError("Uid")
	}
	return nil
}

// Do Remove group member
// Docs: https://api.contest.yandex.net/api/public/swagger-ui.html#/group/deleteMemberUsingDELETE
// meta:operation DELETE /groups/{groupId}/members
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

	return err
}
