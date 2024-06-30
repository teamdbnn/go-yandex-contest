package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetSubmissionsOfContestService Get contest participants
type GetSubmissionsOfContestService struct {
	c                *Client
	contestId        int64
	page             int32
	pageFuncCall     bool
	pageSize         int32
	pageSizeFuncCall bool
}

// ContestID Set contest id
func (s *GetSubmissionsOfContestService) ContestID(contestId int64) *GetSubmissionsOfContestService {
	s.contestId = contestId
	return s
}

// Page Set page, default 1
func (s *GetSubmissionsOfContestService) Page(page int32) *GetSubmissionsOfContestService {
	s.page = page
	s.pageFuncCall = true
	return s
}

// PageSize Set page size, default 100
func (s *GetSubmissionsOfContestService) PageSize(pageSize int32) *GetSubmissionsOfContestService {
	s.pageSize = pageSize
	s.pageSizeFuncCall = true
	return s
}

func (s *GetSubmissionsOfContestService) validate() error {
	if s.contestId == 0 {
		return requiredError("contestId")
	}
	return nil
}

// Do send req
func (s *GetSubmissionsOfContestService) Do(ctx context.Context, opts ...RequestOption) ([]*Submission, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/submissions", s.contestId),
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

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := make([]*Submission, 0)

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SendSubmissionFromURL Send submission from URL
type SendSubmissionFromURL struct {
	c         *Client
	contestId int64
	body      struct {
		fileUrl      string
		fileName     string
		problemAlias string
		compilerId   string
		meta         string
	}
}

// ContestID Set contest id
func (s *SendSubmissionFromURL) ContestID(contestId int64) *SendSubmissionFromURL {
	s.contestId = contestId
	return s
}

// FileUrl Set file URL
func (s *SendSubmissionFromURL) FileUrl(fileUrl string) *SendSubmissionFromURL {
	s.body.fileUrl = fileUrl
	return s
}

// FileName Set file name
func (s *SendSubmissionFromURL) FileName(fileName string) *SendSubmissionFromURL {
	s.body.fileName = fileName
	return s
}

// ProblemAlias Set problem alias
func (s *SendSubmissionFromURL) ProblemAlias(problemAlias string) *SendSubmissionFromURL {
	s.body.problemAlias = problemAlias
	return s
}

// CompilerID Set compiler id
func (s *SendSubmissionFromURL) CompilerID(compilerId string) *SendSubmissionFromURL {
	s.body.compilerId = compilerId
	return s
}

// Meta Set meta
func (s *SendSubmissionFromURL) Meta(meta string) *SendSubmissionFromURL {
	s.body.meta = meta
	return s
}

func (s *SendSubmissionFromURL) validate() error {
	if s.contestId == 0 {
		return requiredError("contestId")
	}
	if s.body.fileUrl == "" {
		return requiredError("fileUrl")
	}
	if s.body.fileName == "" {
		return requiredError("fileName")
	}
	if s.body.problemAlias == "" {
		return requiredError("problemAlias")
	}
	if s.body.compilerId == "" {
		return requiredError("compilerId")
	}
	if s.body.meta == "" {
		return requiredError("meta")
	}
	return nil
}

// Do send req
func (s *SendSubmissionFromURL) Do(ctx context.Context, opts ...RequestOption) (*RunId, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("/contests/%v/submissions/lazy", s.contestId),
	}
	r.setJSONBody(s.body)
	data, err := s.c.callAPI(ctx, r, opts...)

	if err != nil {
		return nil, err
	}
	res := new(RunId)

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetReportForMultipleSubmissions Get report for multiple submissions
type GetReportForMultipleSubmissions struct {
	c         *Client
	contestId int64
	locale    string
	runIds    []int
}

// ContestID Set contest id
func (s *GetReportForMultipleSubmissions) ContestID(contestId int64) *GetReportForMultipleSubmissions {
	s.contestId = contestId
	return s
}

// Locale Set locale, default "ru"
func (s *GetReportForMultipleSubmissions) Locale(locale string) *GetReportForMultipleSubmissions {
	s.locale = locale
	return s
}

// RunIds Set run ids
func (s *GetReportForMultipleSubmissions) RunIds(runIds []int) *GetReportForMultipleSubmissions {
	s.runIds = runIds
	return s
}

func (s *GetReportForMultipleSubmissions) validate() error {
	if s.contestId == 0 {
		return requiredError("contestId")
	}
	return nil
}

// Do send req
func (s *GetReportForMultipleSubmissions) Do(ctx context.Context, opts ...RequestOption) ([]*MultiRunReport, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/submissions/multiple", s.contestId),
	}
	if s.locale == "" {
		r.setParam("locale", "ru")
	} else {
		r.setParam("locale", s.locale)
	}

	r.setParam("runIds", s.runIds)

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := make([]*MultiRunReport, 0)

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetBriefReportForSubmission Get brief report for submission
type GetBriefReportForSubmission struct {
	c            *Client
	contestId    int64
	submissionId int64
}

// ContestID Set contest id
func (s *GetBriefReportForSubmission) ContestID(contestId int64) *GetBriefReportForSubmission {
	s.contestId = contestId
	return s
}

// SubmissionID Set submission id
func (s *GetBriefReportForSubmission) SubmissionID(submissionId int64) *GetBriefReportForSubmission {
	s.submissionId = submissionId
	return s
}

func (s *GetBriefReportForSubmission) validate() error {
	if s.contestId == 0 {
		return requiredError("contestId")
	}
	if s.submissionId == 0 {
		return requiredError("submissionId")
	}
	return nil
}

// Do send req
func (s *GetBriefReportForSubmission) Do(ctx context.Context, opts ...RequestOption) (*BriefRunReport, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/submissions/%v", s.contestId, s.submissionId),
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(BriefRunReport)

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetFullReportForSubmission Get full report for submission
type GetFullReportForSubmission struct {
	c            *Client
	contestId    int64
	submissionId int64
	locale       string
}

// ContestID Set contest id
func (s *GetFullReportForSubmission) ContestID(contestId int64) *GetFullReportForSubmission {
	s.contestId = contestId
	return s
}

// SubmissionID Set submission id
func (s *GetFullReportForSubmission) SubmissionID(submissionId int64) *GetFullReportForSubmission {
	s.submissionId = submissionId
	return s
}
func (s *GetFullReportForSubmission) Locale(locale string) *GetFullReportForSubmission {
	s.locale = locale
	return s
}

func (s *GetFullReportForSubmission) validate() error {
	if s.contestId == 0 {
		return requiredError("contestId")
	}
	if s.submissionId == 0 {
		return requiredError("submissionId")
	}
	return nil
}

// Do send req
func (s *GetFullReportForSubmission) Do(ctx context.Context, opts ...RequestOption) (*FullRunReport, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/submissions/%v/full", s.contestId, s.submissionId),
	}
	if s.locale == "" {
		r.setParam("locale", "ru")
	} else {
		r.setParam("locale", s.locale)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(FullRunReport)

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetSubmissionSourceCode Get submission source code
type GetSubmissionSourceCode struct {
	c            *Client
	contestId    int64
	submissionId int64
}

// ContestID Set contest id
func (s *GetSubmissionSourceCode) ContestID(contestId int64) *GetSubmissionSourceCode {
	s.contestId = contestId
	return s
}

// SubmissionID Set submission id
func (s *GetSubmissionSourceCode) SubmissionID(submissionId int64) *GetSubmissionSourceCode {
	s.submissionId = submissionId
	return s
}

func (s *GetSubmissionSourceCode) validate() error {
	if s.contestId == 0 {
		return requiredError("contestId")
	}
	if s.submissionId == 0 {
		return requiredError("submissionId")
	}
	return nil
}

// Do send req
func (s *GetSubmissionSourceCode) Do(ctx context.Context, opts ...RequestOption) (interface{}, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/submissions/%v/source", s.contestId, s.submissionId),
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(interface{})

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetMetadataOfSubmissionSourceCode Get metadata of submission source code
type GetMetadataOfSubmissionSourceCode struct {
	c            *Client
	contestId    int64
	submissionId int64
}

// ContestID Set contest id
func (s *GetMetadataOfSubmissionSourceCode) ContestID(contestId int64) *GetMetadataOfSubmissionSourceCode {
	s.contestId = contestId
	return s
}

// SubmissionID Set submission id
func (s *GetMetadataOfSubmissionSourceCode) SubmissionID(submissionId int64) *GetMetadataOfSubmissionSourceCode {
	s.submissionId = submissionId
	return s
}

func (s *GetMetadataOfSubmissionSourceCode) validate() error {
	if s.contestId == 0 {
		return requiredError("contestId")
	}
	if s.submissionId == 0 {
		return requiredError("submissionId")
	}
	return nil
}

// Do send req
func (s *GetMetadataOfSubmissionSourceCode) Do(ctx context.Context, opts ...RequestOption) (interface{}, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodHead,
		endpoint: fmt.Sprintf("/contests/%v/submissions/%v/source", s.contestId, s.submissionId),
	}

	_, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// GetFullAnswerFileForTest Get full answer file for test
type GetFullAnswerFileForTest struct {
	c                 *Client
	contestId         int64
	submissionId      int64
	testName          string
	useReportSettings bool
}

// ContestID Set contest id
func (s *GetFullAnswerFileForTest) ContestID(contestId int64) *GetFullAnswerFileForTest {
	s.contestId = contestId
	return s
}

// SubmissionID Set submission id
func (s *GetFullAnswerFileForTest) SubmissionID(submissionId int64) *GetFullAnswerFileForTest {
	s.submissionId = submissionId
	return s
}

// TestName Set testName
func (s *GetFullAnswerFileForTest) TestName(testName string) *GetFullAnswerFileForTest {
	s.testName = testName
	return s
}

// UseReportSettings Set use report settings
func (s *GetFullAnswerFileForTest) UseReportSettings(useReportSettings bool) *GetFullAnswerFileForTest {
	s.useReportSettings = useReportSettings
	return s
}

func (s *GetFullAnswerFileForTest) validate() error {
	if s.contestId == 0 {
		return requiredError("contestId")
	}
	if s.submissionId == 0 {
		return requiredError("submissionId")
	}
	if s.testName == "" {
		return requiredError("testName")
	}
	return nil
}

// Do send req
func (s *GetFullAnswerFileForTest) Do(ctx context.Context, opts ...RequestOption) (interface{}, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/submissions/%v/%v/answer", s.contestId, s.submissionId, s.testName),
	}
	r.setParam("useReportSettings", s.useReportSettings)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(interface{})
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetFullInputFileForTest Get full input file for test
type GetFullInputFileForTest struct {
	c                 *Client
	contestId         int64
	submissionId      int64
	testName          string
	useReportSettings bool
}

// ContestID Set contest id
func (s *GetFullInputFileForTest) ContestID(contestId int64) *GetFullInputFileForTest {
	s.contestId = contestId
	return s
}

// SubmissionID Set submission id
func (s *GetFullInputFileForTest) SubmissionID(submissionId int64) *GetFullInputFileForTest {
	s.submissionId = submissionId
	return s
}

// TestName Set test name
func (s *GetFullInputFileForTest) TestName(testName string) *GetFullInputFileForTest {
	s.testName = testName
	return s
}

// UseReportSettings Set use report settings
func (s *GetFullInputFileForTest) UseReportSettings(useReportSettings bool) *GetFullInputFileForTest {
	s.useReportSettings = useReportSettings
	return s
}

func (s *GetFullInputFileForTest) validate() error {
	if s.contestId == 0 {
		return requiredError("contestId")
	}
	if s.submissionId == 0 {
		return requiredError("submissionId")
	}
	if s.testName == "" {
		return requiredError("testName")
	}
	return nil
}

// Do send req
func (s *GetFullInputFileForTest) Do(ctx context.Context, opts ...RequestOption) (interface{}, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/submissions/%v/%v/input", s.contestId, s.submissionId, s.testName),
	}
	r.setParam("useReportSettings", s.useReportSettings)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(interface{})

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetParticipantOutputForTest Get participant output for test
type GetParticipantOutputForTest struct {
	c                 *Client
	contestId         int64
	submissionId      int64
	testName          string
	useReportSettings bool
}

// ContestID Set contest id
func (s *GetParticipantOutputForTest) ContestID(contestId int64) *GetParticipantOutputForTest {
	s.contestId = contestId
	return s
}

// SubmissionId Set submission id
func (s *GetParticipantOutputForTest) SubmissionId(submissionId int64) *GetParticipantOutputForTest {
	s.submissionId = submissionId
	return s
}

// TestName Set testName
func (s *GetParticipantOutputForTest) TestName(testName string) *GetParticipantOutputForTest {
	s.testName = testName
	return s
}

// UseReportSettings Set use report settings
func (s *GetParticipantOutputForTest) UseReportSettings(useReportSettings bool) *GetParticipantOutputForTest {
	s.useReportSettings = useReportSettings
	return s
}

func (s *GetParticipantOutputForTest) validate() error {
	if s.contestId == 0 {
		return requiredError("contestId")
	}
	if s.submissionId == 0 {
		return requiredError("submissionId")
	}
	if s.testName == "" {
		return requiredError("testName")
	}
	return nil
}

// Do send req
func (s *GetParticipantOutputForTest) Do(ctx context.Context, opts ...RequestOption) (interface{}, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/submissions/%v/%v/output", s.contestId, s.submissionId, s.testName),
	}
	r.setParam("useReportSettings", s.useReportSettings)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(interface{})
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SubmissionRejudgeService Rejudge submission
type SubmissionRejudgeService struct {
	c            *Client
	submissionId int64
}

// SubmissionID Set submission id
func (s *SubmissionRejudgeService) SubmissionID(submissionId int64) *SubmissionRejudgeService {
	s.submissionId = submissionId
	return s
}

func (s *SubmissionRejudgeService) validate() error {
	if s.submissionId == 0 {
		return requiredError("submissionId")
	}
	return nil
}

// Do send req
func (s *SubmissionRejudgeService) Do(ctx context.Context, opts ...RequestOption) (interface{}, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("/submissions/%v/rejudge", s.submissionId),
	}
	_, err := s.c.callAPI(ctx, r, opts...)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
