package yacontest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client Define API client
type Client struct {
	OAuthToken string
	BaseURL    string
	UserAgent  string
	HTTPClient *http.Client
	Debug      bool
	Logger     logger
	TimeOffset int64
	do         doFunc
}
type doFunc func(req *http.Request) (*http.Response, error)
type logger func(format string, v ...interface{})

const (
	baseAPIMainURL = "https://api.contest.yandex.net/api/public/v2"
)

// NewClient Create new http client
func NewClient(oauthToken string) *Client {
	return &Client{
		OAuthToken: oauthToken,
		BaseURL:    baseAPIMainURL,
		UserAgent:  "Yandex-Contest/golang",
		HTTPClient: http.DefaultClient,
	}
}

// WithLogger add logger
func (c *Client) WithLogger(l logger) *Client {
	c.Logger = l
	return c
}

// debug Put data in log
func (c *Client) debug(format string, v ...interface{}) {
	if !c.Debug || c.Logger == nil {
		return
	}

	c.Logger(format, v...)
}

func (c *Client) parseRequest(r *request, opts ...RequestOption) (err error) {
	for _, opt := range opts {
		opt(r)
	}
	err = r.validate()
	if err != nil {
		return err
	}
	fullURL := fmt.Sprintf("%s%s", c.BaseURL, r.endpoint)

	queryString := r.query.Encode()
	body := &bytes.Buffer{}
	bodyString := r.form.Encode()
	header := http.Header{}
	if r.header != nil {
		header = r.header.Clone()
	}

	if bodyString != "" {
		header.Set("Content-Type", "application/x-www-form-urlencoded")
		body = bytes.NewBufferString(bodyString)
	} else if r.json != nil {
		header.Set("Content-Type", "application/json")
		jsonString, _ := json.Marshal(r.json)
		body = bytes.NewBuffer(jsonString)
	}

	header.Set("Authorization", fmt.Sprintf("%s %s", "OAuth", c.OAuthToken))

	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}
	c.debug("full url: %s, body: %s", fullURL, bodyString)
	r.fullURL = fullURL
	r.header = header
	r.body = body
	return nil
}

func (c *Client) callAPI(ctx context.Context, r *request, opts ...RequestOption) (data []byte, err error) {
	err = c.parseRequest(r, opts...)
	if err != nil {
		return []byte{}, err
	}
	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		return []byte{}, err
	}
	req = req.WithContext(ctx)
	req.Header = r.header
	c.debug("request: %#v", req)
	f := c.do
	if f == nil {
		f = c.HTTPClient.Do
	}
	res, err := f(req)
	if err != nil {
		return []byte{}, err
	}
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	defer func() {
		cerr := res.Body.Close()
		if err == nil && cerr != nil {
			err = cerr
		}
	}()
	c.debug("response: %#v", res)
	c.debug("response body: %#v", string(data))
	c.debug("response status code: %#v", res.StatusCode)
	if res.StatusCode >= http.StatusBadRequest {
		apiErr := &APIError{
			Code:    int64(res.StatusCode),
			Message: string(data),
		}
		return nil, apiErr
	}

	return data, nil
}

// NewGetParticipantsOfCompetition Get participants of competition
func (c *Client) NewGetParticipantsOfCompetition() *GetParticipantsOfCompetition {
	return &GetParticipantsOfCompetition{c: c}
}

// NewRegisterParticipantIntoCompetition Register participant into competition
func (c *Client) NewRegisterParticipantIntoCompetition() *RegisterParticipantIntoCompetition {
	return &RegisterParticipantIntoCompetition{c: c}
}

// NewGetCompilersList Get compilers list
func (c *Client) NewGetCompilersList() *GetCompilersList {
	return &GetCompilersList{c: c}
}

// NewGetContestInfo Get contest info
func (c *Client) NewGetContestInfo() *GetContestInfo {
	return &GetContestInfo{c: c}
}

// NewCreateGroupForCompetition Create group for competition
func (c *Client) NewCreateGroupForCompetition() *CreateGroupForCompetition {
	return &CreateGroupForCompetition{c: c}
}

// NewGetGroupInfoOfCompetition Get group info of competition
func (c *Client) NewGetGroupInfoOfCompetition() *GetGroupInfoOfCompetition {
	return &GetGroupInfoOfCompetition{c: c}
}

// NewAddGroupMemberForCompetition Add group member for competition
func (c *Client) NewAddGroupMemberForCompetition() *AddGroupMemberForCompetition {
	return &AddGroupMemberForCompetition{c: c}
}

// NewRemoveGroupOfCompetition Remove group of competition
func (c *Client) NewRemoveGroupOfCompetition() *RemoveGroupOfCompetition {
	return &RemoveGroupOfCompetition{c: c}
}

// NewUserGeneratePassword Generate user password
func (c *Client) NewUserGeneratePassword() *UserGeneratePassword {
	return &UserGeneratePassword{c: c}
}

// NewGetJuryClarifications Get jury clarifications
func (c *Client) NewGetJuryClarifications() *GetJuryClarifications {
	return &GetJuryClarifications{c: c}
}

// NewGetContestMessages Get contest messages
func (c *Client) NewGetContestMessages() *GetContestMessages {
	return &GetContestMessages{c: c}
}

// NewSendQuestionToJury Send question to jury
func (c *Client) NewSendQuestionToJury() *SendQuestionToJury {
	return &SendQuestionToJury{c: c}
}

// NewGetAllSubmissionsForContest Get all submissions for contest
func (c *Client) NewGetAllSubmissionsForContest() *GetAllSubmissionsForContest {
	return &GetAllSubmissionsForContest{c: c}
}

// NewGetYourSubmissionsForContest Get your submissions for contest
func (c *Client) NewGetYourSubmissionsForContest() *GetYourSubmissionsForContest {
	return &GetYourSubmissionsForContest{c: c}
}

// NewGetListOfGroupsForContest Get groups list for contest
func (c *Client) NewGetListOfGroupsForContest() *GetListOfGroupsForContest {
	return &GetListOfGroupsForContest{c: c}
}

// NewRegisterGroupForContest Register group for contest
func (c *Client) NewRegisterGroupForContest() *RegisterGroupForContest {
	return &RegisterGroupForContest{c: c}
}

// NewDeleteGroupFromContest Delete group from contest
func (c *Client) NewDeleteGroupFromContest() *DeleteGroupFromContest {
	return &DeleteGroupFromContest{c: c}
}

// NewUpdateGroupForContest Update group for contest
func (c *Client) NewUpdateGroupForContest() *UpdateGroupForContest {
	return &UpdateGroupForContest{c: c}
}

// NewGetParticipantsOfContest Get participants of contest
func (c *Client) NewGetParticipantsOfContest() *GetParticipantsOfContest {
	return &GetParticipantsOfContest{c: c}
}

// NewRegisterParticipantForContest Register participant for contest
func (c *Client) NewRegisterParticipantForContest() *RegisterParticipantForContest {
	return &RegisterParticipantForContest{c: c}
}

// NewGetInfoOfParticipant Get info of participant
func (c *Client) NewGetInfoOfParticipant() *GetInfoOfParticipant {
	return &GetInfoOfParticipant{c: c}
}

// NewStartContestForParticipant Start contest for participant
func (c *Client) NewStartContestForParticipant() *StartContestForParticipant {
	return &StartContestForParticipant{c: c}
}

// NewUnregisterParticipantFromContest Unregister participant from contest
func (c *Client) NewUnregisterParticipantFromContest() *UnregisterParticipantFromContest {
	return &UnregisterParticipantFromContest{c: c}
}

// NewUpdateParticipantForContest Update participant for contest
func (c *Client) NewUpdateParticipantForContest() *UpdateParticipantForContest {
	return &UpdateParticipantForContest{c: c}
}

// NewGetInfoOfYourParticipation Get info of your participation
func (c *Client) NewGetInfoOfYourParticipation() *GetInfoOfYourParticipation {
	return &GetInfoOfYourParticipation{c: c}
}

// NewStartContest Start contest
func (c *Client) NewStartContest() *StartContest {
	return &StartContest{c: c}
}

// NewUnregisterYourselfFromContest Unregister yourself from contest
func (c *Client) NewUnregisterYourselfFromContest() *UnregisterYourselfFromContest {
	return &UnregisterYourselfFromContest{c: c}
}

// NewGetContestProblems Get contest problems
func (c *Client) NewGetContestProblems() *GetContestProblems {
	return &GetContestProblems{c: c}
}

// NewGetProblemStatement Get problem statement
func (c *Client) NewGetProblemStatement() *GetProblemStatement {
	return &GetProblemStatement{c: c}
}

// NewGetProblemFile Get problem file
func (c *Client) NewGetProblemFile() *GetProblemFile {
	return &GetProblemFile{c: c}
}

// NewGetSubmissionsQueueCapacity Get submissions queue capacity
func (c *Client) NewGetSubmissionsQueueCapacity() *GetSubmissionsQueueCapacity {
	return &GetSubmissionsQueueCapacity{c: c}
}

// NewGetAvailableScopes Get available scopes
func (c *Client) NewGetAvailableScopes() *GetAvailableScopes {
	return &GetAvailableScopes{c: c}
}

// NewGetContestStandings Get contest standings
func (c *Client) NewGetContestStandings() *GetContestStandings {
	return &GetContestStandings{c: c}
}

// NewGetContestStandingsExtended Get extended contest standings
func (c *Client) NewGetContestStandingsExtended() *GetContestStandingsExtended {
	return &GetContestStandingsExtended{c: c}
}

// NewGetYourPositionInContestStandings Get your position in contest standings
func (c *Client) NewGetYourPositionInContestStandings() *GetYourPositionInContestStandings {
	return &GetYourPositionInContestStandings{c: c}
}

// NewGetSubmissionsOfContest Get submissions of contest
func (c *Client) NewGetSubmissionsOfContest() *GetSubmissionsOfContest {
	return &GetSubmissionsOfContest{c: c}
}

// NewSendSubmission Send submission
func (c *Client) NewSendSubmission() *SendSubmission {
	return &SendSubmission{c: c}
}

// NewSendSubmissionFromURL Send submissions from URL
func (c *Client) NewSendSubmissionFromURL() *SendSubmissionFromURL {
	return &SendSubmissionFromURL{c: c}
}

// NewGetReportForMultipleSubmissions Get report for multiple submissions
func (c *Client) NewGetReportForMultipleSubmissions() *GetReportForMultipleSubmissions {
	return &GetReportForMultipleSubmissions{c: c}
}

// NewGetBriefReportForSubmission Get brief report for submission
func (c *Client) NewGetBriefReportForSubmission() *GetBriefReportForSubmission {
	return &GetBriefReportForSubmission{c: c}
}

// NewGetFullReportForSubmission Get full report for submission
func (c *Client) NewGetFullReportForSubmission() *GetFullReportForSubmission {
	return &GetFullReportForSubmission{c: c}
}

// NewGetSubmissionSourceCode Get submission source code
func (c *Client) NewGetSubmissionSourceCode() *GetSubmissionSourceCode {
	return &GetSubmissionSourceCode{c: c}
}

// NewGetMetadataOfSubmissionSourceCode Get metadata of submission source code
func (c *Client) NewGetMetadataOfSubmissionSourceCode() *GetMetadataOfSubmissionSourceCode {
	return &GetMetadataOfSubmissionSourceCode{c: c}
}

// NewGetFullAnswerFileForTest Get full answer file for test
func (c *Client) NewGetFullAnswerFileForTest() *GetFullAnswerFileForTest {
	return &GetFullAnswerFileForTest{c: c}
}

// NewGetFullInputFileForTest Get full input file for test
func (c *Client) NewGetFullInputFileForTest() *GetFullInputFileForTest {
	return &GetFullInputFileForTest{c: c}
}

// NewGetParticipantOutputForTest Get participant output for test
func (c *Client) NewGetParticipantOutputForTest() *GetParticipantOutputForTest {
	return &GetParticipantOutputForTest{c: c}
}

// NewSubmissionRejudge Submission rejudge
func (c *Client) NewSubmissionRejudge() *SubmissionRejudge {
	return &SubmissionRejudge{c: c}
}

// NewGetUserTeams Get user teams
func (c *Client) NewGetUserTeams() *GetUserTeams {
	return &GetUserTeams{c: c}
}
