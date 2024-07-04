package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// GetListOfGroupsForContest Get list of groups for contest
type GetListOfGroupsForContest struct {
	c         *Client
	contestID int64
}

// ContestID Set contest id
func (s *GetListOfGroupsForContest) ContestID(contestID int64) *GetListOfGroupsForContest {
	s.contestID = contestID
	return s
}

func (s *GetListOfGroupsForContest) validate() error {
	if s.contestID == 0 {
		return requiredError("contestID")
	}
	return nil
}

// Do Send GET request
func (s *GetListOfGroupsForContest) Do(ctx context.Context, opts ...RequestOption) ([]*ContestGroup, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/groups", s.contestID),
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := make([]*ContestGroup, 0)

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// RegisterGroupForContest Register group for contest
type RegisterGroupForContest struct {
	c         *Client
	contestID int64
	groupID   int64
	body      struct {
		Roles []string `json:"roles"`
	}
}

// ContestID Set contest id
func (s *RegisterGroupForContest) ContestID(contestID int64) *RegisterGroupForContest {
	s.contestID = contestID
	return s
}

// GroupID Set group Id
func (s *RegisterGroupForContest) GroupID(groupID int64) *RegisterGroupForContest {
	s.groupID = groupID
	return s
}

// Roles Set roles
func (s *RegisterGroupForContest) Roles(roles []string) *RegisterGroupForContest {
	s.body.Roles = roles
	return s
}

func (s *RegisterGroupForContest) validate() error {
	if s.contestID == 0 {
		return requiredError("contestID")
	}
	if s.groupID == 0 {
		return requiredError("groupID")
	}
	if s.body.Roles == nil {
		return requiredError("Roles")
	}
	return nil
}

// Do Send POST request
func (s *RegisterGroupForContest) Do(ctx context.Context, opts ...RequestOption) error {
	if err := s.validate(); err != nil {
		return err
	}

	r := &request{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("/contests/%v/groups/%v", s.contestID, s.groupID),
	}
	r.setJSONBody(s.body)
	_, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}

	return nil
}

// DeleteGroupFromContest Delete group from contest
type DeleteGroupFromContest struct {
	c         *Client
	contestID int64
	groupID   int64
}

// ContestID Set contest id
func (s *DeleteGroupFromContest) ContestID(contestID int64) *DeleteGroupFromContest {
	s.contestID = contestID
	return s
}

// GroupID Set group id
func (s *DeleteGroupFromContest) GroupID(groupID int64) *DeleteGroupFromContest {
	s.groupID = groupID
	return s
}

func (s *DeleteGroupFromContest) validate() error {
	if s.contestID == 0 {
		return requiredError("contestID")
	}
	if s.groupID == 0 {
		return requiredError("groupID")
	}
	return nil
}

// Do Send DELETE request
func (s *DeleteGroupFromContest) Do(ctx context.Context, opts ...RequestOption) error {
	if err := s.validate(); err != nil {
		return err
	}

	r := &request{
		method:   http.MethodDelete,
		endpoint: fmt.Sprintf("/contests/%v/groups/%v", s.contestID, s.groupID),
	}

	_, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}

	return nil
}

// UpdateGroupForContest Update group for contest
type UpdateGroupForContest struct {
	c         *Client
	contestID int64
	groupID   int64
	body      struct {
		Roles []string `json:"roles"`
	}
}

// ContestID Set contest id
func (s *UpdateGroupForContest) ContestID(contestID int64) *UpdateGroupForContest {
	s.contestID = contestID
	return s
}

// GroupID Set group Id
func (s *UpdateGroupForContest) GroupID(groupID int64) *UpdateGroupForContest {
	s.groupID = groupID
	return s
}

// Roles Set roles
func (s *UpdateGroupForContest) Roles(roles []string) *UpdateGroupForContest {
	s.body.Roles = roles
	return s
}

func (s *UpdateGroupForContest) validate() error {
	if s.contestID == 0 {
		return requiredError("contestID")
	}
	if s.groupID == 0 {
		return requiredError("groupID")
	}
	if s.body.Roles == nil {
		return requiredError("Roles")
	}
	return nil
}

// Do Send PATCH request
func (s *UpdateGroupForContest) Do(ctx context.Context, opts ...RequestOption) error {
	if err := s.validate(); err != nil {
		return err
	}

	r := &request{
		method:   http.MethodPatch,
		endpoint: fmt.Sprintf("/contests/%v/groups/%v", s.contestID, s.groupID),
	}
	r.setJSONBody(s.body)
	_, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}

	return nil
}

// GetParticipantsOfContest Get contest participants
type GetParticipantsOfContest struct {
	c           *Client
	contestID   int64
	displayName string
	login       string
}

// ContestID Set contest id
func (s *GetParticipantsOfContest) ContestID(contestID int64) *GetParticipantsOfContest {
	s.contestID = contestID
	return s
}

// DisplayName Set display name
func (s *GetParticipantsOfContest) DisplayName(displayName string) *GetParticipantsOfContest {
	s.displayName = displayName
	return s
}

// Login Set login
func (s *GetParticipantsOfContest) Login(login string) *GetParticipantsOfContest {
	s.login = login
	return s
}

func (s *GetParticipantsOfContest) validate() error {
	if s.contestID == 0 {
		return requiredError("contestId")
	}
	return nil
}

// Do Send GET request
func (s *GetParticipantsOfContest) Do(ctx context.Context, opts ...RequestOption) ([]*ParticipantInfo, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/participants", s.contestID),
	}
	if s.displayName != "" {
		r.setParam("display_name", s.displayName)
	}
	if s.login != "" {
		r.setParam("login", s.login)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := make([]*ParticipantInfo, 0)

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// RegisterParticipantForContest Register participant for contest
type RegisterParticipantForContest struct {
	c         *Client
	contestID int64
	login     string
	uid       int64
}

// ContestID Set contest id
func (s *RegisterParticipantForContest) ContestID(contestID int64) *RegisterParticipantForContest {
	s.contestID = contestID
	return s
}

// Login Set login
func (s *RegisterParticipantForContest) Login(login string) *RegisterParticipantForContest {
	s.login = login
	return s
}

// UID Set uid
func (s *RegisterParticipantForContest) UID(uid int64) *RegisterParticipantForContest {
	s.uid = uid
	return s
}

func (s *RegisterParticipantForContest) validate() error {
	if s.contestID == 0 {
		return requiredError("contestID")
	}
	if s.login == "" {
		return requiredError("login")
	}
	return nil
}

// Do Send POST request
func (s *RegisterParticipantForContest) Do(ctx context.Context, opts ...RequestOption) (*int64, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("/contests/%v/participants", s.contestID),
	}
	if s.login != "" {
		r.setParam("login", s.login)
	}

	if s.uid != 0 {
		r.setParam("uid", s.uid) // todo: Check usage
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := string(data)
	i, err := strconv.ParseInt(res, 10, 64)
	if err != nil {
		return nil, err
	}

	return &i, nil
}

// GetInfoOfParticipant Get info about participant
type GetInfoOfParticipant struct {
	c             *Client
	contestID     int64
	participantID int64
}

// ContestID Set contest id
func (s *GetInfoOfParticipant) ContestID(contestID int64) *GetInfoOfParticipant {
	s.contestID = contestID
	return s
}

// ParticipantID Set participant id
func (s *GetInfoOfParticipant) ParticipantID(participantID int64) *GetInfoOfParticipant {
	s.participantID = participantID
	return s
}

func (s *GetInfoOfParticipant) validate() error {
	if s.contestID == 0 {
		return requiredError("contestID")
	}
	if s.participantID == 0 {
		return requiredError("participantID")
	}
	return nil
}

// Do Send GET request
func (s *GetInfoOfParticipant) Do(ctx context.Context, opts ...RequestOption) (*ParticipantStatus, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/participants/%v", s.contestID, s.participantID),
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(ParticipantStatus)

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// StartContestForParticipant Start contest for participant
type StartContestForParticipant struct {
	c             *Client
	contestID     int64
	participantID int64
}

// ContestID Set contest id
func (s *StartContestForParticipant) ContestID(contestID int64) *StartContestForParticipant {
	s.contestID = contestID
	return s
}

// ParticipantID Set participant id
func (s *StartContestForParticipant) ParticipantID(participantID int64) *StartContestForParticipant {
	s.participantID = participantID
	return s
}

func (s *StartContestForParticipant) validate() error {
	if s.contestID == 0 {
		return requiredError("contestID")
	}
	if s.participantID == 0 {
		return requiredError("participantID")
	}
	return nil
}

// Do Send PUT request
func (s *StartContestForParticipant) Do(ctx context.Context, opts ...RequestOption) error {
	if err := s.validate(); err != nil {
		return err
	}

	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("/contests/%v/participants/%v", s.contestID, s.participantID),
	}

	_, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}

	return nil
}

// UnregisterParticipantFromContest Unregister participant from contest
type UnregisterParticipantFromContest struct {
	c             *Client
	contestID     int64
	participantID int64
}

// ContestID Set contest id
func (s *UnregisterParticipantFromContest) ContestID(contestID int64) *UnregisterParticipantFromContest {
	s.contestID = contestID
	return s
}

// ParticipantID Set participant id
func (s *UnregisterParticipantFromContest) ParticipantID(participantID int64) *UnregisterParticipantFromContest {
	s.participantID = participantID
	return s
}

func (s *UnregisterParticipantFromContest) validate() error {
	if s.contestID == 0 {
		return requiredError("contestID")
	}
	if s.participantID == 0 {
		return requiredError("participantID")
	}
	return nil
}

// Do Send DELETE request
func (s *UnregisterParticipantFromContest) Do(ctx context.Context, opts ...RequestOption) error {
	if err := s.validate(); err != nil {
		return err
	}

	r := &request{
		method:   http.MethodDelete,
		endpoint: fmt.Sprintf("/contests/%v/participants/%v", s.contestID, s.participantID),
	}

	_, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}

	return nil
}

// UpdateParticipantForContest Update participant for contest
type UpdateParticipantForContest struct {
	c         *Client
	contestID int64
	body      struct {
		DisplayedName string `json:"displayedName"`
	}
	participantID int64
}

// ContestID Set contest id
func (s *UpdateParticipantForContest) ContestID(contestID int64) *UpdateParticipantForContest {
	s.contestID = contestID
	return s
}

// ParticipantID Set participant id
func (s *UpdateParticipantForContest) ParticipantID(participantID int64) *UpdateParticipantForContest {
	s.participantID = participantID
	return s
}

// DisplayedName Set displayedName
func (s *UpdateParticipantForContest) DisplayedName(displayedName string) *UpdateParticipantForContest {
	s.body.DisplayedName = displayedName
	return s
}

func (s *UpdateParticipantForContest) validate() error {
	if s.contestID == 0 {
		return requiredError("contestID")
	}
	if s.participantID == 0 {
		return requiredError("participantID")
	}
	if s.body.DisplayedName == "" {
		return requiredError("DisplayedName")
	}
	return nil
}

// Do Send PATCH request
func (s *UpdateParticipantForContest) Do(ctx context.Context, opts ...RequestOption) error {
	if err := s.validate(); err != nil {
		return err
	}

	r := &request{
		method:   http.MethodPatch,
		endpoint: fmt.Sprintf("/contests/%v/participants/%v", s.contestID, s.participantID),
	}

	r.setJSONBody(s.body)

	_, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}

	return nil
}

// GetInfoOfYourParticipation Get info of your participation
type GetInfoOfYourParticipation struct {
	c         *Client
	contestID int64
}

// ContestID Set contest id
func (s *GetInfoOfYourParticipation) ContestID(contestID int64) *GetInfoOfYourParticipation {
	s.contestID = contestID
	return s
}

func (s *GetInfoOfYourParticipation) validate() error {
	if s.contestID == 0 {
		return requiredError("contestID")
	}
	return nil
}

// Do Send GET request
func (s *GetInfoOfYourParticipation) Do(ctx context.Context, opts ...RequestOption) (*ParticipantStatus, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/participation", s.contestID),
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(ParticipantStatus)

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// StartContest Start contest
type StartContest struct {
	c         *Client
	contestID int64
}

// ContestID Set contest id
func (s *StartContest) ContestID(contestID int64) *StartContest {
	s.contestID = contestID
	return s
}

func (s *StartContest) validate() error {
	if s.contestID == 0 {
		return requiredError("contestID")
	}
	return nil
}

// Do Send PUT request
func (s *StartContest) Do(ctx context.Context, opts ...RequestOption) error {
	if err := s.validate(); err != nil {
		return err
	}

	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("/contests/%v/participation", s.contestID),
	}
	_, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}
	return nil
}

// UnregisterYourselfFromContest Unregister yourself from contest
type UnregisterYourselfFromContest struct {
	c         *Client
	contestID int64
}

// ContestID Set contest id
func (s *UnregisterYourselfFromContest) ContestID(contestID int64) *UnregisterYourselfFromContest {
	s.contestID = contestID
	return s
}

func (s *UnregisterYourselfFromContest) validate() error {
	if s.contestID == 0 {
		return requiredError("contestID")
	}
	return nil
}

// Do Send DELETE request
func (s *UnregisterYourselfFromContest) Do(ctx context.Context, opts ...RequestOption) error {
	if err := s.validate(); err != nil {
		return err
	}

	r := &request{
		method:   http.MethodDelete,
		endpoint: fmt.Sprintf("/contests/%v/participation", s.contestID),
	}
	_, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}

	return nil
}
