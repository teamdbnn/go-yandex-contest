package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// GetListOfGroupsForContestService Get list of groups for contest
type GetListOfGroupsForContestService struct {
	c         *Client
	contestID int64
}

// ContestID Set contest id
func (s *GetListOfGroupsForContestService) ContestID(contestID int64) *GetListOfGroupsForContestService {
	s.contestID = contestID
	return s
}

func (s *GetListOfGroupsForContestService) validate() error {
	if s.contestID == 0 {
		return requiredError("contestID")
	}
	return nil
}

// Do send req
func (s *GetListOfGroupsForContestService) Do(ctx context.Context, opts ...RequestOption) ([]*ContestGroup, error) {
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

type RegisterGroupForContestService struct {
	c         *Client
	contestID int64
	groupID   int64
	body      struct {
		Roles []string `json:"roles"`
	}
}

// ContestID Set contest id
func (s *RegisterGroupForContestService) ContestID(contestID int64) *RegisterGroupForContestService {
	s.contestID = contestID
	return s
}

// GroupID Set group Id
func (s *RegisterGroupForContestService) GroupID(groupID int64) *RegisterGroupForContestService {
	s.groupID = groupID
	return s
}

// Roles Set roles
func (s *RegisterGroupForContestService) Roles(roles []string) *RegisterGroupForContestService {
	s.body.Roles = roles
	return s
}

func (s *RegisterGroupForContestService) validate() error {
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

// Do send req
func (s *RegisterGroupForContestService) Do(ctx context.Context, opts ...RequestOption) error {
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

// DeleteGroupFromContestService Delete group from contest
type DeleteGroupFromContestService struct {
	c         *Client
	contestID int64
	groupID   int64
}

// ContestID Set contest id
func (s *DeleteGroupFromContestService) ContestID(contestID int64) *DeleteGroupFromContestService {
	s.contestID = contestID
	return s
}

// GroupID Set group id
func (s *DeleteGroupFromContestService) GroupID(groupID int64) *DeleteGroupFromContestService {
	s.groupID = groupID
	return s
}

func (s *DeleteGroupFromContestService) validate() error {
	if s.contestID == 0 {
		return requiredError("contestID")
	}
	if s.groupID == 0 {
		return requiredError("groupID")
	}
	return nil
}

// Do send req
func (s *DeleteGroupFromContestService) Do(ctx context.Context, opts ...RequestOption) error {
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

type UpdateGroupForContestService struct {
	c         *Client
	contestID int64
	groupID   int64
	body      struct {
		Roles []string `json:"roles"`
	}
}

// ContestID Set contest id
func (s *UpdateGroupForContestService) ContestID(contestID int64) *UpdateGroupForContestService {
	s.contestID = contestID
	return s
}

// GroupID Set group Id
func (s *UpdateGroupForContestService) GroupID(groupID int64) *UpdateGroupForContestService {
	s.groupID = groupID
	return s
}

// Roles Set roles
func (s *UpdateGroupForContestService) Roles(roles []string) *UpdateGroupForContestService {
	s.body.Roles = roles
	return s
}

func (s *UpdateGroupForContestService) validate() error {
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

// Do send req
func (s *UpdateGroupForContestService) Do(ctx context.Context, opts ...RequestOption) error {
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

// GetParticipantsOfContestService Get contest participants
type GetParticipantsOfContestService struct {
	c           *Client
	contestID   int64
	displayName string
	login       string
}

// ContestID Set contest id
func (s *GetParticipantsOfContestService) ContestID(contestID int64) *GetParticipantsOfContestService {
	s.contestID = contestID
	return s
}

// DisplayName Set display name
func (s *GetParticipantsOfContestService) DisplayName(displayName string) *GetParticipantsOfContestService {
	s.displayName = displayName
	return s
}

// Login Set login
func (s *GetParticipantsOfContestService) Login(login string) *GetParticipantsOfContestService {
	s.login = login
	return s
}

func (s *GetParticipantsOfContestService) validate() error {
	if s.contestID == 0 {
		return requiredError("contestId")
	}
	return nil
}

// Do send req
func (s *GetParticipantsOfContestService) Do(ctx context.Context, opts ...RequestOption) ([]*ParticipantInfo, error) {
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

// RegisterParticipantForContestService Register participant for contest
type RegisterParticipantForContestService struct {
	c         *Client
	contestID int64
	login     string
	uid       int64
}

// ContestID Set contest id
func (s *RegisterParticipantForContestService) ContestID(contestID int64) *RegisterParticipantForContestService {
	s.contestID = contestID
	return s
}

// Login Set login
func (s *RegisterParticipantForContestService) Login(login string) *RegisterParticipantForContestService {
	s.login = login
	return s
}

// UID Set uid
func (s *RegisterParticipantForContestService) UID(uid int64) *RegisterParticipantForContestService {
	s.uid = uid
	return s
}

func (s *RegisterParticipantForContestService) validate() error {
	if s.contestID == 0 {
		return requiredError("contestID")
	}
	if s.login == "" {
		return requiredError("login")
	}
	return nil
}

// Do send req
func (s *RegisterParticipantForContestService) Do(ctx context.Context, opts ...RequestOption) (*int64, error) {
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

// Do send req
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

// Do send req
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

// UnregisterParticipantFromContestService Unregister participant from contest
type UnregisterParticipantFromContestService struct {
	c             *Client
	contestID     int64
	participantID int64
}

// ContestID Set contest id
func (s *UnregisterParticipantFromContestService) ContestID(contestID int64) *UnregisterParticipantFromContestService {
	s.contestID = contestID
	return s
}

// ParticipantID Set participant id
func (s *UnregisterParticipantFromContestService) ParticipantID(participantID int64) *UnregisterParticipantFromContestService {
	s.participantID = participantID
	return s
}

func (s *UnregisterParticipantFromContestService) validate() error {
	if s.contestID == 0 {
		return requiredError("contestID")
	}
	if s.participantID == 0 {
		return requiredError("participantID")
	}
	return nil
}

// Do send req
func (s *UnregisterParticipantFromContestService) Do(ctx context.Context, opts ...RequestOption) error {
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

// UpdateParticipantForContestService Update participant for contest
type UpdateParticipantForContestService struct {
	c         *Client
	contestID int64
	body      struct {
		DisplayedName string `json:"displayedName"`
	}
	participantID int64
}

// ContestID Set contest id
func (s *UpdateParticipantForContestService) ContestID(contestID int64) *UpdateParticipantForContestService {
	s.contestID = contestID
	return s
}

// ParticipantID Set participant id
func (s *UpdateParticipantForContestService) ParticipantID(participantID int64) *UpdateParticipantForContestService {
	s.participantID = participantID
	return s
}

// DisplayedName Set displayedName
func (s *UpdateParticipantForContestService) DisplayedName(displayedName string) *UpdateParticipantForContestService {
	s.body.DisplayedName = displayedName
	return s
}

func (s *UpdateParticipantForContestService) validate() error {
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

// Do send req
func (s *UpdateParticipantForContestService) Do(ctx context.Context, opts ...RequestOption) error {
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

// Do send req
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

// Do send req
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

// Do send req
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
