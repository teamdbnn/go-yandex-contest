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
	contestId int64
}

// ContestID Set contest id
func (s *GetListOfGroupsForContestService) ContestID(contestId int64) *GetListOfGroupsForContestService {
	s.contestId = contestId
	return s
}

func (s *GetListOfGroupsForContestService) validate() error {
	if s.contestId == 0 {
		return requiredError("contestId")
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
		endpoint: fmt.Sprintf("/contests/%v/groups", s.contestId),
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
	contestId int64
	groupId   int64
	body      struct {
		roles []string
	}
}

// ContestID Set contest id
func (s *RegisterGroupForContestService) ContestID(contestId int64) *RegisterGroupForContestService {
	s.contestId = contestId
	return s
}

// GroupID Set group Id
func (s *RegisterGroupForContestService) GroupID(groupId int64) *RegisterGroupForContestService {
	s.groupId = groupId
	return s
}

// Roles Set roles
func (s *RegisterGroupForContestService) Roles(roles []string) *RegisterGroupForContestService {
	s.body.roles = roles
	return s
}

func (s *RegisterGroupForContestService) validate() error {
	if s.contestId == 0 {
		return requiredError("contestId")
	}
	if s.groupId == 0 {
		return requiredError("groupId")
	}
	return nil
}

// Do send req
func (s *RegisterGroupForContestService) Do(ctx context.Context, opts ...RequestOption) (interface{}, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("/contests/%v/groups/%v", s.contestId, s.groupId),
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	fmt.Println(data)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// DeleteGroupFromContestService Delete group from contest
type DeleteGroupFromContestService struct {
	c         *Client
	contestId int64
	groupId   int64
}

// ContestID Set contest id
func (s *DeleteGroupFromContestService) ContestID(contestId int64) *DeleteGroupFromContestService {
	s.contestId = contestId
	return s
}

// GroupID Set group id
func (s *DeleteGroupFromContestService) GroupID(groupId int64) *DeleteGroupFromContestService {
	s.groupId = groupId
	return s
}

func (s *DeleteGroupFromContestService) validate() error {
	if s.contestId == 0 {
		return requiredError("contestId")
	}
	if s.groupId == 0 {
		return requiredError("groupId")
	}
	return nil
}

// Do send req
func (s *DeleteGroupFromContestService) Do(ctx context.Context, opts ...RequestOption) (interface{}, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodDelete,
		endpoint: fmt.Sprintf("/contests/%v/groups/%v", s.contestId, s.groupId),
	}

	_, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

type UpdateGroupForContestService struct {
	c         *Client
	contestId int64
	groupId   int64
	body      struct {
		roles []string
	}
}

// ContestID Set contest id
func (s *UpdateGroupForContestService) ContestID(contestId int64) *UpdateGroupForContestService {
	s.contestId = contestId
	return s
}

// GroupID Set group Id
func (s *UpdateGroupForContestService) GroupID(groupId int64) *UpdateGroupForContestService {
	s.groupId = groupId
	return s
}

// Roles Set roles
func (s *UpdateGroupForContestService) Roles(roles []string) *UpdateGroupForContestService {
	s.body.roles = roles
	return s
}

func (s *UpdateGroupForContestService) validate() error {
	if s.contestId == 0 {
		return requiredError("contestId")
	}
	if s.groupId == 0 {
		return requiredError("groupId")
	}
	return nil
}

// Do send req
func (s *UpdateGroupForContestService) Do(ctx context.Context, opts ...RequestOption) (interface{}, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPatch,
		endpoint: fmt.Sprintf("/contests/%v/groups/%v", s.contestId, s.groupId),
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	fmt.Println(data)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// GetParticipantsOfContestService Get contest participants
type GetParticipantsOfContestService struct {
	c           *Client
	contestId   int64
	displayName string
	login       string
}

// ContestID Set contest id
func (s *GetParticipantsOfContestService) ContestID(contestId int64) *GetParticipantsOfContestService {
	s.contestId = contestId
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
	if s.contestId == 0 {
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
		endpoint: fmt.Sprintf("/contests/%v/participants", s.contestId),
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
	contestId int64
	login     string
	uid       int64
}

// ContestID Set contest id
func (s *RegisterParticipantForContestService) ContestID(contestId int64) *RegisterParticipantForContestService {
	s.contestId = contestId
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
	if s.contestId == 0 {
		return requiredError("contestId")
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
		endpoint: fmt.Sprintf("/contests/%v/participants", s.contestId),
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
	contestId     int64
	participantId int64
}

// ContestID Set contest id
func (s *GetInfoOfParticipant) ContestID(contestId int64) *GetInfoOfParticipant {
	s.contestId = contestId
	return s
}

// ParticipantID Set participant id
func (s *GetInfoOfParticipant) ParticipantID(participantId int64) *GetInfoOfParticipant {
	s.participantId = participantId
	return s
}

func (s *GetInfoOfParticipant) validate() error {
	if s.contestId == 0 {
		return requiredError("contestId")
	}
	if s.participantId == 0 {
		return requiredError("participantId")
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
		endpoint: fmt.Sprintf("/contests/%v/participants/%v", s.contestId, s.participantId),
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
	contestId     int64
	participantId int64
}

// ContestID Set contest id
func (s *StartContestForParticipant) ContestID(contestId int64) *StartContestForParticipant {
	s.contestId = contestId
	return s
}

// ParticipantID Set participant id
func (s *StartContestForParticipant) ParticipantID(participantId int64) *StartContestForParticipant {
	s.participantId = participantId
	return s
}

func (s *StartContestForParticipant) validate() error {
	if s.contestId == 0 {
		return requiredError("contest")
	}
	if s.participantId == 0 {
		return requiredError("participant")
	}
	return nil
}

// Do send req
func (s *StartContestForParticipant) Do(ctx context.Context, opts ...RequestOption) (interface{}, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("/contests/%v/participants/%v", s.contestId, s.participantId),
	}

	_, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// UnregisterParticipantFromContestService Unregister participant from contest
type UnregisterParticipantFromContestService struct {
	c             *Client
	contestId     int64
	participantId int64
}

// ContestID Set contest id
func (s *UnregisterParticipantFromContestService) ContestID(contestId int64) *UnregisterParticipantFromContestService {
	s.contestId = contestId
	return s
}

// ParticipantID Set participant id
func (s *UnregisterParticipantFromContestService) ParticipantID(participantId int64) *UnregisterParticipantFromContestService {
	s.participantId = participantId
	return s
}

func (s *UnregisterParticipantFromContestService) validate() error {
	if s.contestId == 0 {
		return requiredError("contestId")
	}
	if s.participantId == 0 {
		return requiredError("participantId")
	}
	return nil
}

// Do send req
func (s *UnregisterParticipantFromContestService) Do(ctx context.Context, opts ...RequestOption) (interface{}, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodDelete,
		endpoint: fmt.Sprintf("/contests/%v/participants/%v", s.contestId, s.participantId),
	}

	_, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// UpdateParticipantForContestService Update participant for contest
type UpdateParticipantForContestService struct {
	c         *Client
	contestId int64
	body      struct {
		displayedName string
	}
	participantId int64
}

// ContestID Set contest id
func (s *UpdateParticipantForContestService) ContestID(contestId int64) *UpdateParticipantForContestService {
	s.contestId = contestId
	return s
}

// ParticipantID Set participant id
func (s *UpdateParticipantForContestService) ParticipantID(participantId int64) *UpdateParticipantForContestService {
	s.participantId = participantId
	return s
}

// DisplayedName Set displayedName
func (s *UpdateParticipantForContestService) DisplayedName(name string) *UpdateParticipantForContestService {
	s.body.displayedName = name
	return s
}

func (s *UpdateParticipantForContestService) validate() error {
	if s.contestId == 0 {
		return requiredError("contestId")
	}
	if s.participantId == 0 {
		return requiredError("participantId")
	}
	if s.body.displayedName == "" {
		return requiredError("displayedName")
	}
	return nil
}

// Do send req
func (s *UpdateParticipantForContestService) Do(ctx context.Context, opts ...RequestOption) (interface{}, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPatch,
		endpoint: fmt.Sprintf("/contests/%v/participants/%v", s.contestId, s.participantId),
	}

	r.setJSONBody(s.body)

	_, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// GetInfoOfYourParticipation Get info of your participation
type GetInfoOfYourParticipation struct {
	c         *Client
	contestId int64
}

// ContestID Set contest id
func (s *GetInfoOfYourParticipation) ContestID(contestId int64) *GetInfoOfYourParticipation {
	s.contestId = contestId
	return s
}

func (s *GetInfoOfYourParticipation) validate() error {
	if s.contestId == 0 {
		return requiredError("contestId")
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
		endpoint: fmt.Sprintf("/contests/%v/participation", s.contestId),
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
	contestId int64
}

// ContestID Set contest id
func (s *StartContest) ContestID(contestId int64) *StartContest {
	s.contestId = contestId
	return s
}

func (s *StartContest) validate() error {
	if s.contestId == 0 {
		return requiredError("contestId")
	}
	return nil
}

// Do send req
func (s *StartContest) Do(ctx context.Context, opts ...RequestOption) (interface{}, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("/contests/%v/participation", s.contestId),
	}
	_, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// UnregisterYourselfFromContest Unregister yourself from contest
type UnregisterYourselfFromContest struct {
	c         *Client
	contestId int64
}

// ContestID Set contest id
func (s *UnregisterYourselfFromContest) ContestID(contestId int64) *UnregisterYourselfFromContest {
	s.contestId = contestId
	return s
}

func (s *UnregisterYourselfFromContest) validate() error {
	if s.contestId == 0 {
		return requiredError("contest")
	}
	return nil
}

// Do send req
func (s *UnregisterYourselfFromContest) Do(ctx context.Context, opts ...RequestOption) (interface{}, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodDelete,
		endpoint: fmt.Sprintf("/contests/%v/participation", s.contestId),
	}
	_, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
