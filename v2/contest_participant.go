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
		return requiredFieldError("contestID")
	}
	return nil
}

// Do List groups registered for contest
//
// Docs: https://api.contest.yandex.net/api/public/swagger-ui.html#/participant/groupsUsingGET
// meta:operation GET /contests/{contestId}/groups
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
	if err = json.Unmarshal(data, &res); err != nil {
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
		return requiredFieldError("contestID")
	}
	if s.groupID == 0 {
		return requiredFieldError("groupID")
	}
	if s.body.Roles == nil {
		return requiredFieldError("Roles")
	}
	return nil
}

// Do Register group for contest
//
// Docs: https://api.contest.yandex.net/api/public/swagger-ui.html#/participant/registerGroupUsingPOST
// meta:operation POST /contests/{contestId}/groups/{groupId}
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

	return err
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
		return requiredFieldError("contestID")
	}
	if s.groupID == 0 {
		return requiredFieldError("groupID")
	}
	return nil
}

// Do Delete group from contest
//
// Docs: https://api.contest.yandex.net/api/public/swagger-ui.html#/participant/deleteGroupUsingDELETE
// meta:operation DELETE /contests/{contestId}/groups/{groupId}
func (s *DeleteGroupFromContest) Do(ctx context.Context, opts ...RequestOption) error {
	if err := s.validate(); err != nil {
		return err
	}

	r := &request{
		method:   http.MethodDelete,
		endpoint: fmt.Sprintf("/contests/%v/groups/%v", s.contestID, s.groupID),
	}
	_, err := s.c.callAPI(ctx, r, opts...)

	return err
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
		return requiredFieldError("contestID")
	}
	if s.groupID == 0 {
		return requiredFieldError("groupID")
	}
	if len(s.body.Roles) == 0 {
		return requiredFieldError("Roles")
	}
	return nil
}

// Do Change group registration info
//
// Docs: https://api.contest.yandex.net/api/public/swagger-ui.html#/participant/updateGroupUsingPATCH
// meta:operation PATCH /contests/{contestId}/groups/{groupId}
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

	return err
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
		return requiredFieldError("contestId")
	}
	return nil
}

// Do Get contest participants
//
// Docs: https://api.contest.yandex.net/api/public/swagger-ui.html#/participant/getParticipantsUsingGET
// meta:operation GET /contests/{contestId}/participants
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
	if err = json.Unmarshal(data, &res); err != nil {
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
		return requiredFieldError("contestID")
	}
	if s.login == "" {
		return requiredFieldError("login")
	}
	return nil
}

// Do Register for contest
//
// Docs: https://api.contest.yandex.net/api/public/swagger-ui.html#/participant/registerForContestUsingPOST
// meta:operation POST /contests/{contestId}/participants
func (s *RegisterParticipantForContest) Do(ctx context.Context, opts ...RequestOption) (int64, error) {
	if err := s.validate(); err != nil {
		return 0, err
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
		return 0, err
	}

	i, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return 0, err
	}

	return i, nil
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
		return requiredFieldError("contestID")
	}
	if s.participantID == 0 {
		return requiredFieldError("participantID")
	}
	return nil
}

// Do Get information about participant
//
// Docs: https://api.contest.yandex.net/api/public/swagger-ui.html#/participant/getParticipantStatusUsingGET
// meta:operation GET /contests/{contestId}/participants/{participantId}
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
	if err = json.Unmarshal(data, res); err != nil {
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
		return requiredFieldError("contestID")
	}
	if s.participantID == 0 {
		return requiredFieldError("participantID")
	}
	return nil
}

// Do Start the contest for participant
//
// Docs: https://api.contest.yandex.net/api/public/swagger-ui.html#/participant/startContestForParticipantUsingPUT
// meta:operation PUT /contests/{contestId}/participants/{participantId}
func (s *StartContestForParticipant) Do(ctx context.Context, opts ...RequestOption) error {
	if err := s.validate(); err != nil {
		return err
	}

	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("/contests/%v/participants/%v", s.contestID, s.participantID),
	}
	_, err := s.c.callAPI(ctx, r, opts...)

	return err
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
		return requiredFieldError("contestID")
	}
	if s.participantID == 0 {
		return requiredFieldError("participantID")
	}
	return nil
}

// Do Unregister participant from contest
//
// Docs: https://api.contest.yandex.net/api/public/swagger-ui.html#/participant/unregisterFromContestUsingDELETE
// meta:operation DELETE /contests/{contestId}/participants/{participantId}
func (s *UnregisterParticipantFromContest) Do(ctx context.Context, opts ...RequestOption) error {
	if err := s.validate(); err != nil {
		return err
	}

	r := &request{
		method:   http.MethodDelete,
		endpoint: fmt.Sprintf("/contests/%v/participants/%v", s.contestID, s.participantID),
	}
	_, err := s.c.callAPI(ctx, r, opts...)

	return err
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
		return requiredFieldError("contestID")
	}
	if s.participantID == 0 {
		return requiredFieldError("participantID")
	}
	if s.body.DisplayedName == "" {
		return requiredFieldError("DisplayedName")
	}
	return nil
}

// Do Send PATCH request
//
// Docs: https://api.contest.yandex.net/api/public/swagger-ui.html#/participant/updateParticipantUsingPATCH
// meta:operation PATCH /contests/{contestId}/participants/{participantId}
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

	return err
}

// GetParticipantsStats Get participant statistics
type GetParticipantsStats struct {
	c             *Client
	contestID     int64
	participantID int64
}

// ContestID Set contest id
func (s *GetParticipantsStats) ContestID(contestID int64) *GetParticipantsStats {
	s.contestID = contestID
	return s
}

// ParticipantID Set participant id
func (s *GetParticipantsStats) ParticipantID(participantID int64) *GetParticipantsStats {
	s.participantID = participantID
	return s
}

func (s *GetParticipantsStats) validate() error {
	if s.contestID == 0 {
		return requiredFieldError("contestID")
	}
	if s.participantID == 0 {
		return requiredFieldError("participantID")
	}
	return nil
}

// Do Get participant statistics
//
// Docs: https://api.contest.yandex.net/api/public/swagger-ui.html#/participant/getParticipantStatsUsingGET
// meta:operation GET /contests/{contestId}/participants/{participantId}/stats
func (s *GetParticipantsStats) Do(ctx context.Context, opts ...RequestOption) (*ParticipantStats, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/participants/%v/stats", s.contestID, s.participantID),
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(ParticipantStats)
	if err = json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, err
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
		return requiredFieldError("contestID")
	}
	return nil
}

// Do Get information about your participation
//
// Docs: https://api.contest.yandex.net/api/public/swagger-ui.html#/participant/getYourStatusUsingGET
// meta:operation GET /contests/{contestId}/participation
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
	if err = json.Unmarshal(data, res); err != nil {
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
		return requiredFieldError("contestID")
	}
	return nil
}

// Do Start the contest
//
// Docs: https://api.contest.yandex.net/api/public/swagger-ui.html#/participant/startContestUsingPUT
// meta:operation PUT /contests/{contestId}/participation
func (s *StartContest) Do(ctx context.Context, opts ...RequestOption) error {
	if err := s.validate(); err != nil {
		return err
	}

	r := &request{
		method:   http.MethodPut,
		endpoint: fmt.Sprintf("/contests/%v/participation", s.contestID),
	}
	_, err := s.c.callAPI(ctx, r, opts...)

	return err
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
		return requiredFieldError("contestID")
	}
	return nil
}

// Do Unregister yourself from contest
//
// Docs: https://api.contest.yandex.net/api/public/swagger-ui.html#/participant/unregisterYourselfFromContestUsingDELETE
// meta:operation DELETE /contests/{contestId}/participation
func (s *UnregisterYourselfFromContest) Do(ctx context.Context, opts ...RequestOption) error {
	if err := s.validate(); err != nil {
		return err
	}

	r := &request{
		method:   http.MethodDelete,
		endpoint: fmt.Sprintf("/contests/%v/participation", s.contestID),
	}
	_, err := s.c.callAPI(ctx, r, opts...)

	return err
}
