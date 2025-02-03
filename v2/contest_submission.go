package yacontest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// GetSubmissionsOfContest Get contest participants
type GetSubmissionsOfContest struct {
	c                *Client
	contestID        int64
	page             int32
	pageFuncCall     bool
	pageSize         int32
	pageSizeFuncCall bool
}

// ContestID Set contest id
func (s *GetSubmissionsOfContest) ContestID(contestID int64) *GetSubmissionsOfContest {
	s.contestID = contestID
	return s
}

// Page Set page, default 1
func (s *GetSubmissionsOfContest) Page(page int32) *GetSubmissionsOfContest {
	s.page = page
	return s
}

// PageSize Set page size, default 100
func (s *GetSubmissionsOfContest) PageSize(pageSize int32) *GetSubmissionsOfContest {
	s.pageSize = pageSize
	return s
}

func (s *GetSubmissionsOfContest) validate() error {
	if s.contestID == 0 {
		return requiredFieldError("contestID")
	}
	return nil
}

// Do Send GET request
//
// Docs:
// meta:operation
func (s *GetSubmissionsOfContest) Do(ctx context.Context, opts ...RequestOption) (*Submissions, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/submissions", s.contestID),
	}
	if s.page == 0 {
		s.page = defaultPage
	}
	if s.pageSize == 0 {
		s.pageSize = defaultSubmissionPageSize
	}
	r.setParam("pageSize", s.pageSize)
	r.setParam("page", s.page)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(Submissions)
	if err = json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// SendSubmission Send submission
type SendSubmission struct {
	c                 *Client
	compiler          string
	file              string
	problem           string
	submissionMeta    string
	tcmSubmissionType string
	contestID         int64
}

// Compiler Set compiler id
func (s *SendSubmission) Compiler(compiler string) *SendSubmission {
	s.compiler = compiler
	return s
}

// File Set file URL
func (s *SendSubmission) File(file string) *SendSubmission {
	s.file = file
	return s
}

// Problem Set problem
func (s *SendSubmission) Problem(problem string) *SendSubmission {
	s.problem = problem
	return s
}

// SubmissionMeta Set submission meta
func (s *SendSubmission) SubmissionMeta(submissionMeta string) *SendSubmission {
	s.submissionMeta = submissionMeta
	return s
}

// TcmSubmissionType Set tcm submission type
func (s *SendSubmission) TcmSubmissionType(tcmSubmissionType string) *SendSubmission {
	s.tcmSubmissionType = tcmSubmissionType
	return s
}

// ContestID Set contest id
func (s *SendSubmission) ContestID(contestID int64) *SendSubmission {
	s.contestID = contestID
	return s
}

func (s *SendSubmission) validate() error {
	if s.compiler == "" {
		return requiredFieldError("compiler")
	}
	if s.file == "" {
		return requiredFieldError("file")
	}
	if s.problem == "" {
		return requiredFieldError("problem")
	}
	if s.contestID == 0 {
		return requiredFieldError("contestID")
	}
	return nil
}

// Do Send POST request
//
// Docs:
// meta:operation
func (s *SendSubmission) Do(ctx context.Context, opts ...RequestOption) (*RunID, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("/contests/%v/submissions", s.contestID),
	}

	f, err := os.Open(s.file)
	if err != nil {
		return nil, err
	}

	defer func() {
		cerr := f.Close()
		if err == nil {
			err = cerr
		}
	}()

	// todo: check file uploads
	r.setFormParam("file", f)
	r.setFormParam("compiler", s.compiler)
	r.setFormParam("problem", s.problem)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(RunID)
	if err = json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// SendSubmissionFromURL Send submission from URL
type SendSubmissionFromURL struct {
	c         *Client
	contestID int64
	body      LazySubmissionRequest
}

// ContestID Set contest id
func (s *SendSubmissionFromURL) ContestID(contestID int64) *SendSubmissionFromURL {
	s.contestID = contestID
	return s
}

// FileURL Set file URL
func (s *SendSubmissionFromURL) FileURL(fileUrl string) *SendSubmissionFromURL {
	s.body.FileURL = fileUrl
	return s
}

// FileName Set file name
func (s *SendSubmissionFromURL) FileName(fileName string) *SendSubmissionFromURL {
	s.body.FileName = fileName
	return s
}

// ProblemAlias Set problem alias
func (s *SendSubmissionFromURL) ProblemAlias(problemAlias string) *SendSubmissionFromURL {
	s.body.ProblemAlias = problemAlias
	return s
}

// CompilerID Set compiler id
func (s *SendSubmissionFromURL) CompilerID(compilerID string) *SendSubmissionFromURL {
	s.body.CompilerID = compilerID
	return s
}

// Meta Set meta
func (s *SendSubmissionFromURL) Meta(meta string) *SendSubmissionFromURL {
	s.body.Meta = meta
	return s
}

func (s *SendSubmissionFromURL) validate() error {
	if s.contestID == 0 {
		return requiredFieldError("contestID")
	}
	if s.body.FileURL == "" {
		return requiredFieldError("FileURL")
	}
	if s.body.FileName == "" {
		return requiredFieldError("FileName")
	}
	if s.body.ProblemAlias == "" {
		return requiredFieldError("ProblemAlias")
	}
	if s.body.CompilerID == "" {
		return requiredFieldError("CompilerID")
	}
	if s.body.Meta == "" {
		return requiredFieldError("Meta")
	}
	return nil
}

// Do Send POST request
//
// Docs:
// meta:operation
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
	if err = json.Unmarshal(data, res); err != nil {
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

// RunIDs Set run ids
func (s *GetReportForMultipleSubmissions) RunIDs(runIDs []int) *GetReportForMultipleSubmissions {
	s.runIDs = runIDs
	return s
}

func (s *GetReportForMultipleSubmissions) validate() error {
	if s.contestID == 0 {
		return requiredFieldError("contestID")
	}
	if s.runIDs == nil {
		return requiredFieldError("runIDs")
	}
	return nil
}

// Do Send GET request
//
// Docs:
// meta:operation
func (s *GetReportForMultipleSubmissions) Do(ctx context.Context, opts ...RequestOption) ([]*MultiRunReport, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/submissions/multiple", s.contestID),
	}
	if s.locale == "" {
		s.locale = defaultLocale
	}

	r.setParam("locale", s.locale)
	for i := 0; i < len(s.runIDs); i++ {
		r.addParam("runIds", s.runIDs[i])
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := make([]*MultiRunReport, 0)
	if err = json.Unmarshal(data, &res); err != nil {
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
		return requiredFieldError("contestID")
	}
	if s.submissionID == 0 {
		return requiredFieldError("submissionID")
	}
	return nil
}

// Do Send GET request
//
// Docs:
// meta:operation
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
	if err = json.Unmarshal(data, res); err != nil {
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
		return requiredFieldError("contestID")
	}
	if s.submissionID == 0 {
		return requiredFieldError("submissionID")
	}
	return nil
}

// Do Send GET request
//
// Docs:
// meta:operation
func (s *GetFullReportForSubmission) Do(ctx context.Context, opts ...RequestOption) (*FullRunReport, error) {
	if err := s.validate(); err != nil {
		return nil, err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/submissions/%v/full", s.contestID, s.submissionID),
	}
	if s.locale == "" {
		s.locale = defaultLocale
	}
	r.setParam("locale", s.locale)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}

	res := new(FullRunReport)
	if err = json.Unmarshal(data, res); err != nil {
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
		return requiredFieldError("contestID")
	}
	if s.submissionID == 0 {
		return requiredFieldError("submissionID")
	}
	return nil
}

// Do Send GET request
//
// Docs:
// meta:operation
func (s *GetSubmissionSourceCode) Do(ctx context.Context, opts ...RequestOption) (string, error) {
	if err := s.validate(); err != nil {
		return "", err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/submissions/%v/source", s.contestID, s.submissionID),
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return "", err
	}

	return string(data), nil
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
		return requiredFieldError("contestID")
	}
	if s.submissionID == 0 {
		return requiredFieldError("submissionID")
	}
	return nil
}

// Do Send HEAD request
//
// Docs:
// meta:operation
func (s *GetMetadataOfSubmissionSourceCode) Do(ctx context.Context, opts ...RequestOption) (string, error) {
	if err := s.validate(); err != nil {
		return "", err
	}

	r := &request{
		method:   http.MethodHead,
		endpoint: fmt.Sprintf("/contests/%v/submissions/%v/source", s.contestID, s.submissionID),
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return "", err
	}

	return string(data), nil
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
		return requiredFieldError("contestID")
	}
	if s.submissionID == 0 {
		return requiredFieldError("submissionID")
	}
	if s.testName == "" {
		return requiredFieldError("testName")
	}
	return nil
}

// Do Send GET request todo: check file download
//
// Docs:
// meta:operation
func (s *GetFullAnswerFileForTest) Do(ctx context.Context, opts ...RequestOption) error {
	if err := s.validate(); err != nil {
		return err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/submissions/%v/%v/answer", s.contestID, s.submissionID, s.testName),
	}
	r.setParam("useReportSettings", s.useReportSettings)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}

	f, err := os.Create("fullAnswerFileForTest")
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	if err == nil {
		return nil
	}

	return err
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
		return requiredFieldError("contestID")
	}
	if s.submissionID == 0 {
		return requiredFieldError("submissionID")
	}
	if s.testName == "" {
		return requiredFieldError("testName")
	}
	return nil
}

// Do Send GET request
//
// Docs:
// meta:operation
func (s *GetFullInputFileForTest) Do(ctx context.Context, opts ...RequestOption) error {
	if err := s.validate(); err != nil {
		return err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/submissions/%v/%v/input", s.contestID, s.submissionID, s.testName),
	}
	r.setParam("useReportSettings", s.useReportSettings)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return err
	}

	f, err := os.Create("fullInputFileForTest")
	if err != nil {
		return err
	}

	_, err = f.Write(data)
	if err == nil {
		return nil
	}

	return err
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
		return requiredFieldError("contestID")
	}
	if s.submissionID == 0 {
		return requiredFieldError("submissionID")
	}
	if s.testName == "" {
		return requiredFieldError("testName")
	}
	return nil
}

// Do Send GET request
//
// Docs:
// meta:operation
func (s *GetParticipantOutputForTest) Do(ctx context.Context, opts ...RequestOption) (string, error) {
	if err := s.validate(); err != nil {
		return "", err
	}

	r := &request{
		method:   http.MethodGet,
		endpoint: fmt.Sprintf("/contests/%v/submissions/%v/%v/output", s.contestID, s.submissionID, s.testName),
	}
	r.setParam("useReportSettings", s.useReportSettings)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// SubmissionRejudge Rejudge submission
type SubmissionRejudge struct {
	c            *Client
	submissionID int64
}

// SubmissionID Set submission id
func (s *SubmissionRejudge) SubmissionID(submissionID int64) *SubmissionRejudge {
	s.submissionID = submissionID
	return s
}

func (s *SubmissionRejudge) validate() error {
	if s.submissionID == 0 {
		return requiredFieldError("submissionID")
	}
	return nil
}

// Do Send POST request
//
// Docs:
// meta:operation
func (s *SubmissionRejudge) Do(ctx context.Context, opts ...RequestOption) error {
	if err := s.validate(); err != nil {
		return err
	}

	r := &request{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("/submissions/%v/rejudge", s.submissionID),
	}
	_, err := s.c.callAPI(ctx, r, opts...)
	if err == nil {
		return nil
	}

	return err
}
