package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// GetSubmissionsOfContestService Get contest participants
type GetSubmissionsOfContestService struct {
	c                *Client
	contestID        int64
	page             int32
	pageFuncCall     bool
	pageSize         int32
	pageSizeFuncCall bool
}

// ContestID Set contest id
func (s *GetSubmissionsOfContestService) ContestID(contestID int64) *GetSubmissionsOfContestService {
	s.contestID = contestID
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
	if s.contestID == 0 {
		return requiredError("contestID")
	}
	return nil
}

// Do send req
func (s *GetSubmissionsOfContestService) Do(ctx context.Context, opts ...RequestOption) (*Submissions, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/submissions", s.contestID),
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

	res := new(Submissions)

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SendSubmissionFromURL Send submission from URL
type SendSubmissionFromURL struct {
	c         *Client
	contestID int64
	body      struct {
		FileUrl      string `json:"fileUrl"`
		FileName     string `json:"fileName"`
		ProblemAlias string `json:"problemAlias"`
		CompilerID   string `json:"compilerID"`
		Meta         string `json:"meta"`
	}
}

// ContestID Set contest id
func (s *SendSubmissionFromURL) ContestID(contestID int64) *SendSubmissionFromURL {
	s.contestID = contestID
	return s
}

// FileUrl Set file URL
func (s *SendSubmissionFromURL) FileUrl(FileUrl string) *SendSubmissionFromURL {
	s.body.FileUrl = FileUrl
	return s
}

// FileName Set file name
func (s *SendSubmissionFromURL) FileName(FileName string) *SendSubmissionFromURL {
	s.body.FileName = FileName
	return s
}

// ProblemAlias Set problem alias
func (s *SendSubmissionFromURL) ProblemAlias(ProblemAlias string) *SendSubmissionFromURL {
	s.body.ProblemAlias = ProblemAlias
	return s
}

// CompilerID Set compiler id
func (s *SendSubmissionFromURL) CompilerID(CompilerID string) *SendSubmissionFromURL {
	s.body.CompilerID = CompilerID
	return s
}

// Meta Set meta
func (s *SendSubmissionFromURL) Meta(Meta string) *SendSubmissionFromURL {
	s.body.Meta = Meta
	return s
}

func (s *SendSubmissionFromURL) validate() error {
	if s.contestID == 0 {
		return requiredError("contestID")
	}
	if s.body.FileUrl == "" {
		return requiredError("FileUrl")
	}
	if s.body.FileName == "" {
		return requiredError("FileName")
	}
	if s.body.ProblemAlias == "" {
		return requiredError("ProblemAlias")
	}
	if s.body.CompilerID == "" {
		return requiredError("CompilerID")
	}
	if s.body.Meta == "" {
		return requiredError("Meta")
	}
	return nil
}

// Do send req
func (s *SendSubmissionFromURL) Do(ctx context.Context, opts ...RequestOption) (*RunID, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("/contests/%v/submissions/lazy", s.contestID),
	}
	r.setJSONBody(s.body)
	data, err := s.c.callAPI(ctx, r, opts...)

	if err != nil {
		return nil, err
	}
	res := new(RunID)

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetReportForMultipleSubmissions Get report for multiple submissions
type GetReportForMultipleSubmissions struct {
	c         *Client
	contestID int64
	locale    string
	runIDs    []int
}

// ContestID Set contest id
func (s *GetReportForMultipleSubmissions) ContestID(contestID int64) *GetReportForMultipleSubmissions {
	s.contestID = contestID
	return s
}

// Locale Set locale, default "ru"
func (s *GetReportForMultipleSubmissions) Locale(locale string) *GetReportForMultipleSubmissions {
	s.locale = locale
	return s
}

// RunIds Set run ids
func (s *GetReportForMultipleSubmissions) RunIds(runIDs []int) *GetReportForMultipleSubmissions {
	s.runIDs = runIDs
	return s
}

func (s *GetReportForMultipleSubmissions) validate() error {
	if s.contestID == 0 {
		return requiredError("contestID")
	}
	if s.runIDs == nil {
		return requiredError("runIDs")
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
		endpoint: fmt.Sprintf("/contests/%v/submissions/multiple", s.contestID),
	}
	if s.locale == "" {
		r.setParam("locale", "ru")
	} else {
		r.setParam("locale", s.locale)
	}

	r.setParam("runIds", s.runIDs)

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
	contestID    int64
	submissionID int64
}

// ContestID Set contest id
func (s *GetBriefReportForSubmission) ContestID(contestID int64) *GetBriefReportForSubmission {
	s.contestID = contestID
	return s
}

// SubmissionID Set submission id
func (s *GetBriefReportForSubmission) SubmissionID(submissionID int64) *GetBriefReportForSubmission {
	s.submissionID = submissionID
	return s
}

func (s *GetBriefReportForSubmission) validate() error {
	if s.contestID == 0 {
		return requiredError("contestID")
	}
	if s.submissionID == 0 {
		return requiredError("submissionID")
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
		endpoint: fmt.Sprintf("/contests/%v/submissions/%v", s.contestID, s.submissionID),
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
	contestID    int64
	submissionID int64
	locale       string
}

// ContestID Set contest id
func (s *GetFullReportForSubmission) ContestID(contestID int64) *GetFullReportForSubmission {
	s.contestID = contestID
	return s
}

// SubmissionID Set submission id
func (s *GetFullReportForSubmission) SubmissionID(submissionID int64) *GetFullReportForSubmission {
	s.submissionID = submissionID
	return s
}
func (s *GetFullReportForSubmission) Locale(locale string) *GetFullReportForSubmission {
	s.locale = locale
	return s
}

func (s *GetFullReportForSubmission) validate() error {
	if s.contestID == 0 {
		return requiredError("contestID")
	}
	if s.submissionID == 0 {
		return requiredError("submissionID")
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
		endpoint: fmt.Sprintf("/contests/%v/submissions/%v/full", s.contestID, s.submissionID),
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
	contestID    int64
	submissionID int64
}

// ContestID Set contest id
func (s *GetSubmissionSourceCode) ContestID(contestID int64) *GetSubmissionSourceCode {
	s.contestID = contestID
	return s
}

// SubmissionID Set submission id
func (s *GetSubmissionSourceCode) SubmissionID(submissionID int64) *GetSubmissionSourceCode {
	s.submissionID = submissionID
	return s
}

func (s *GetSubmissionSourceCode) validate() error {
	if s.contestID == 0 {
		return requiredError("contestID")
	}
	if s.submissionID == 0 {
		return requiredError("submissionID")
	}
	return nil
}

// Do send req
func (s *GetSubmissionSourceCode) Do(ctx context.Context, opts ...RequestOption) ([]byte, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/submissions/%v/source", s.contestID, s.submissionID),
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetMetadataOfSubmissionSourceCode Get metadata of submission source code
type GetMetadataOfSubmissionSourceCode struct {
	c            *Client
	contestID    int64
	submissionID int64
}

// ContestID Set contest id
func (s *GetMetadataOfSubmissionSourceCode) ContestID(contestID int64) *GetMetadataOfSubmissionSourceCode {
	s.contestID = contestID
	return s
}

// SubmissionID Set submission id
func (s *GetMetadataOfSubmissionSourceCode) SubmissionID(submissionID int64) *GetMetadataOfSubmissionSourceCode {
	s.submissionID = submissionID
	return s
}

func (s *GetMetadataOfSubmissionSourceCode) validate() error {
	if s.contestID == 0 {
		return requiredError("contestID")
	}
	if s.submissionID == 0 {
		return requiredError("submissionID")
	}
	return nil
}

// Do send req
func (s *GetMetadataOfSubmissionSourceCode) Do(ctx context.Context, opts ...RequestOption) ([]byte, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodHead,
		endpoint: fmt.Sprintf("/contests/%v/submissions/%v/source", s.contestID, s.submissionID),
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetFullAnswerFileForTest Get full answer file for test
type GetFullAnswerFileForTest struct {
	c                 *Client
	contestID         int64
	submissionID      int64
	testName          string
	useReportSettings bool
}

// ContestID Set contest id
func (s *GetFullAnswerFileForTest) ContestID(contestID int64) *GetFullAnswerFileForTest {
	s.contestID = contestID
	return s
}

// SubmissionID Set submission id
func (s *GetFullAnswerFileForTest) SubmissionID(submissionID int64) *GetFullAnswerFileForTest {
	s.submissionID = submissionID
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
	if s.contestID == 0 {
		return requiredError("contestID")
	}
	if s.submissionID == 0 {
		return requiredError("submissionID")
	}
	if s.testName == "" {
		return requiredError("testName")
	}
	return nil
}

// Do send req
func (s *GetFullAnswerFileForTest) Do(ctx context.Context, opts ...RequestOption) (error, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/submissions/%v/%v/answer", s.contestID, s.submissionID, s.testName),
	}
	r.setParam("useReportSettings", s.useReportSettings)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	f, err := os.Create("fullAnswerFileForTest")
	if err != nil {
		return nil, err
	}

	_, err = f.Write(data)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// GetFullInputFileForTest Get full input file for test
type GetFullInputFileForTest struct {
	c                 *Client
	contestID         int64
	submissionID      int64
	testName          string
	useReportSettings bool
}

// ContestID Set contest id
func (s *GetFullInputFileForTest) ContestID(contestID int64) *GetFullInputFileForTest {
	s.contestID = contestID
	return s
}

// SubmissionID Set submission id
func (s *GetFullInputFileForTest) SubmissionID(submissionID int64) *GetFullInputFileForTest {
	s.submissionID = submissionID
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
	if s.contestID == 0 {
		return requiredError("contestID")
	}
	if s.submissionID == 0 {
		return requiredError("submissionID")
	}
	if s.testName == "" {
		return requiredError("testName")
	}
	return nil
}

// Do send req
func (s *GetFullInputFileForTest) Do(ctx context.Context, opts ...RequestOption) (error, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/submissions/%v/%v/input", s.contestID, s.submissionID, s.testName),
	}
	r.setParam("useReportSettings", s.useReportSettings)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	f, err := os.Create("fullInputFileForTest")
	if err != nil {
		return nil, err
	}

	_, err = f.Write(data)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// GetParticipantOutputForTest Get participant output for test
type GetParticipantOutputForTest struct {
	c                 *Client
	contestID         int64
	submissionID      int64
	testName          string
	useReportSettings bool
}

// ContestID Set contest id
func (s *GetParticipantOutputForTest) ContestID(contestID int64) *GetParticipantOutputForTest {
	s.contestID = contestID
	return s
}

// SubmissionID Set submission id
func (s *GetParticipantOutputForTest) SubmissionID(submissionID int64) *GetParticipantOutputForTest {
	s.submissionID = submissionID
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
	if s.contestID == 0 {
		return requiredError("contestID")
	}
	if s.submissionID == 0 {
		return requiredError("submissionID")
	}
	if s.testName == "" {
		return requiredError("testName")
	}
	return nil
}

// Do send req
func (s *GetParticipantOutputForTest) Do(ctx context.Context, opts ...RequestOption) ([]byte, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/submissions/%v/%v/output", s.contestID, s.submissionID, s.testName),
	}
	r.setParam("useReportSettings", s.useReportSettings)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// SubmissionRejudgeService Rejudge submission
type SubmissionRejudgeService struct {
	c            *Client
	submissionID int64
}

// SubmissionID Set submission id
func (s *SubmissionRejudgeService) SubmissionID(submissionID int64) *SubmissionRejudgeService {
	s.submissionID = submissionID
	return s
}

func (s *SubmissionRejudgeService) validate() error {
	if s.submissionID == 0 {
		return requiredError("submissionID")
	}
	return nil
}

// Do send req
func (s *SubmissionRejudgeService) Do(ctx context.Context, opts ...RequestOption) (error, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("/submissions/%v/rejudge", s.submissionID),
	}
	_, err := s.c.callAPI(ctx, r, opts...)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
