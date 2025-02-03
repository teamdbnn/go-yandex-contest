package yacontest

import (
	"github.com/go-openapi/strfmt"
)

const (
	defaultPage = 1
)

type AdditionalSolutionCheckResult struct {
	CheckerName  string `json:"checkerName,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	Message      string `json:"message,omitempty"`
	Verdict      string `json:"verdict,omitempty"` // Enum: ["OK","FAILED"]
}

type BriefRunReport struct {
	AuthorID       int64   `json:"authorId"`
	CompileLog     string  `json:"compileLog"`
	Compiler       string  `json:"compiler"`
	Diff           string  `json:"diff,omitempty"`
	MaxMemoryUsage int64   `json:"maxMemoryUsage"`
	MaxTimeUsage   int64   `json:"maxTimeUsage"`
	ProblemAlias   string  `json:"problemAlias"`
	ProblemID      string  `json:"problemId"`
	RunID          int64   `json:"runId"`
	Score          float64 `json:"score,omitempty"`
	Source         string  `json:"source,omitempty"`
	SubmissionTime string  `json:"submissionTime"`
	TestNumber     int64   `json:"testNumber"`
	TimeFromStart  int64   `json:"timeFromStart"`
	Verdict        string  `json:"verdict"`
}

type Clarification struct {
	Message string `json:"message,omitempty"`
	Subject string `json:"subject,omitempty"`
}

type Clarifications struct {
	Clarifications []*Clarification `json:"clarifications"`
}

type CompetitionRegisterResponse struct {
	MissedUIDs []string `json:"missedUids"`
}

type CompilerLimit struct {
	CompilerName  string `json:"compilerName,omitempty"`
	IdlenessLimit int64  `json:"idlenessLimit,omitempty"`
	MemoryLimit   int64  `json:"memoryLimit,omitempty"`
	OutputLimit   int64  `json:"outputLimit,omitempty"`
	TimeLimit     int64  `json:"timeLimit,omitempty"`
}

type CompilerListResponse struct {
	Compilers []*PublicCompilerInfo `json:"compilers"`
}

type ContestDescription struct {
	ContestSettings *ContestReportSettings `json:"contestSettings,omitempty"`
	Duration        int64                  `json:"duration,omitempty"`
	FreezeTime      int64                  `json:"freezeTime,omitempty"`
	Name            string                 `json:"name"`
	StartTime       string                 `json:"startTime,omitempty"`
	Type            string                 `json:"type,omitempty"`
	// Upsolving settings
	//  - `ALLOWED_AFTER_PARTICIPATION_ENDS` — participant can upsolve after their participation ends
	//  - `ALLOWED_AFTER_CONTEST_ENDS` — participant can upsolve after contest is finished for everyone
	//  - `PROHIBITED` — participant can not upsolve
	// Enum: ["ALLOWED_AFTER_PARTICIPATION_ENDS","ALLOWED_AFTER_CONTEST_ENDS","PROHIBITED"]
	UpsolvingAllowance string `json:"upsolvingAllowance,omitempty"`
}

type ContestReportSettings struct {
	ShowAnswer               string `json:"showAnswer"`        // Enum: ["NEVER","ALL_FAILED","FIRST_FAILED","ALL_PASSED","ALL","ONLY_SAMPLES"]
	ShowCheckerOutput        string `json:"showCheckerOutput"` // Enum: ["NEVER","ALL_FAILED","FIRST_FAILED","ALL_PASSED","ALL","ONLY_SAMPLES"]
	ShowInput                string `json:"showInput"`         // Enum: ["NEVER","ALL_FAILED","FIRST_FAILED","ALL_PASSED","ALL","ONLY_SAMPLES"]
	ShowLightweightInterface bool   `json:"showLightweightInterface"`
	ShowOutput               string `json:"showOutput"`            // Enum: ["NEVER","ALL_FAILED","FIRST_FAILED","ALL_PASSED","ALL","ONLY_SAMPLES"]
	ShowPostprocessOutput    string `json:"showPostprocessOutput"` // Enum: ["NEVER","ALL_FAILED","FIRST_FAILED","ALL_PASSED","ALL","ONLY_SAMPLES"]
	ShowReport               bool   `json:"showReport"`
	ShowScore                string `json:"showScore"`  // Enum: ["NEVER","SCORED_TESTS","SCORED_TESTS_AND_SAMPLES"]
	ShowStderr               string `json:"showStderr"` // Enum: ["NEVER","ALL_FAILED","FIRST_FAILED","ALL_PASSED","ALL","ONLY_SAMPLES"]
	ShowTestNumber           bool   `json:"showTestNumber"`
	ShowUsedResources        string `json:"showUsedResources"` // Enum: ["NEVER","ALL_FAILED","FIRST_FAILED","ALL_PASSED","ALL","ONLY_SAMPLES"]
	ShowVerdict              string `json:"showVerdict"`       // Enum: ["NEVER","ALL_FAILED","FIRST_FAILED","ALL_PASSED","ALL","ONLY_SAMPLES"]
	StopOnFirstFail          bool   `json:"stopOnFirstFail"`
	StopOnFirstFailInTestSet bool   `json:"stopOnFirstFailInTestSet"`
	StopOnSampleFail         bool   `json:"stopOnSampleFail"`
	TestOnlySamples          bool   `json:"testOnlySamples"`
	UseAcNotOk               bool   `json:"useAcNotOk"`
}

type ContestGroup struct {
	ID    int64    `json:"id"`
	Name  string   `json:"name"`
	Roles []string `json:"roles"`
}

type ContestProblem struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Alias       string           `json:"alias"`
	Compilers   []string         `json:"compilers"`
	Statements  []*Statement     `json:"statements"`
	Limits      []*CompilerLimit `json:"limits"`
	TestCount   int64            `json:"testCount,omitempty"`
	ProblemType string           `json:"problemType"`
}

type ContestProblems struct {
	Problems []*ContestProblem `json:"problems"`
}

type ContestStandings struct {
	Rows       []*ContestStandingsRow   `json:"rows"`
	Titles     []*ContestStandingsTitle `json:"titles"`
	Statistics *ContestStatistics       `json:"statistics"`
}

type ContestStandingsRow struct {
	ParticipantInfo *ParticipantInfo `json:"participantInfo"`
	PlaceFrom       []int64          `json:"placeFrom"`
	PlaceTo         []int64          `json:"placeTo"`
	ProblemResults  []*ProblemResult `json:"problemResults"`
	Score           string           `json:"score"`
}

type ContestStandingsTitle struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Title string `json:"title"`
}

type ContestStatistics struct {
	LastSubmit  *SubmitInfo `json:"lastSubmit,omitempty"`
	LastSuccess *SubmitInfo `json:"lastSuccess,omitempty"`
}

type CreateGroupRequest struct {
	Name string `json:"name"`
}

type FullRunReport struct {
	CheckerLog       []*TestLog                       `json:"checkerLog"`
	CompileLog       string                           `json:"compileLog"`
	Compiler         string                           `json:"compiler"`
	ContestID        int64                            `json:"contestId"`
	ContestName      string                           `json:"contestName"`
	Diff             string                           `json:"diff"`
	FinalScore       string                           `json:"finalScore"`
	GlobalError      string                           `json:"globalError,omitempty"` // Enum: ["NO_TESTS_IN_PROBLEM","NO_CHECKER_IN_PROBLEM","GLOBAL_TL"]
	IP               string                           `json:"ip,omitempty"`
	MaxMemoryUsage   int64                            `json:"maxMemoryUsage"`
	MaxTimeUsage     int64                            `json:"maxTimeUsage"`
	ParticipantInfo  *ParticipantInfo                 `json:"participantInfo"`
	PrecompileChecks []*AdditionalSolutionCheckResult `json:"precompileChecks"`
	PreliminaryScore string                           `json:"preliminaryScore"`
	ProblemAlias     string                           `json:"problemAlias"`
	ProblemID        string                           `json:"problemId"`
	ProblemName      string                           `json:"problemName"`
	RunID            int64                            `json:"runId"`
	Source           string                           `json:"source"`
	Status           string                           `json:"status"` // Enum: ["WAITING","FAILED","RUNNING","FINISHED"]
	SubmissionTime   string                           `json:"submissionTime"`
	TestFileType     string                           `json:"testFileType"` // Enum: ["BINARY","TEXT"]
	TimeFromStart    int64                            `json:"timeFromStart"`
	Verdict          string                           `json:"verdict"`
}

type GroupID struct {
	ID int64 `json:"id"`
}

type GroupInfo struct {
	MemberCount int64  `json:"memberCount"`
	Name        string `json:"name"`
}

type JSONNode any

type LazySubmissionRequest struct {
	FileURL      string `json:"fileUrl"`
	FileName     string `json:"fileName"`
	ProblemAlias string `json:"problemAlias"`
	CompilerID   string `json:"compilerId"`
	Meta         string `json:"meta,omitempty"`
}

type Message struct {
	Answers  []string `json:"answers"`
	Problem  string   `json:"problem,omitempty"`
	Question string   `json:"question,omitempty"`
	Subject  string   `json:"subject,omitempty"`
}

type Messages struct {
	Messages []*Message `json:"messages"`
}

type MultiRunReport struct {
	CompileLog           string                           `json:"compileLog"`
	Compiler             string                           `json:"compiler"`
	ContestID            int64                            `json:"contestId"`
	ContestName          string                           `json:"contestName"`
	FinalScore           string                           `json:"finalScore"`
	GlobalError          string                           `json:"globalError,omitempty"` // Enum: ["NO_TESTS_IN_PROBLEM","NO_CHECKER_IN_PROBLEM","GLOBAL_TL"]
	IP                   string                           `json:"ip,omitempty"`
	MaxMemoryUsage       int64                            `json:"maxMemoryUsage"`
	MaxTimeUsage         int64                            `json:"maxTimeUsage"`
	ParticipantInfo      *ParticipantInfo                 `json:"participantInfo"`
	PostprocessorMessage string                           `json:"postprocessorMessage,omitempty"`
	PrecompileChecks     []*AdditionalSolutionCheckResult `json:"precompileChecks"`
	PreliminaryScore     string                           `json:"preliminaryScore"`
	ProblemAlias         string                           `json:"problemAlias"`
	ProblemID            string                           `json:"problemId"`
	RunID                int64                            `json:"runId"`
	Status               string                           `json:"status,omitempty"` // Enum: ["WAITING","FAILED","RUNNING","FINISHED"]
	SubmissionTime       string                           `json:"submissionTime"`
	TestFileType         string                           `json:"testFileType"` // Enum: ["BINARY","TEXT"]
	Tests                []*TestLog                       `json:"tests"`
	TimeFromStart        int64                            `json:"timeFromStart"`
	Verdict              string                           `json:"verdict"`
}

type NeuripsSubmissionDetails struct {
	Metrics JSONNode `json:"metrics,omitempty"`
}

type NeuripsSubmissionReport struct {
	Compiler            string                   `json:"compiler"`
	ContestID           int64                    `json:"contestId"`
	CreateAt            *strfmt.DateTime         `json:"createAt"`
	Meta                string                   `json:"meta,omitempty"`
	ParticipantName     string                   `json:"participantName"`
	ProblemID           string                   `json:"problemId"`
	Rank                int64                    `json:"rank"`
	Result              *NeuripsSubmissionResult `json:"result,omitempty"`
	RunID               int64                    `json:"runId"`
	Status              string                   `json:"status,omitempty"` // Enum: ["WAITING","FAILED","RUNNING","FINISHED"]
	TimeFromStartMillis int64                    `json:"timeFromStartMillis,omitempty"`
}

type NeuripsSubmissionResult struct {
	Details         *NeuripsSubmissionDetails `json:"details,omitempty"`
	Score           float64                   `json:"score,omitempty"`
	UsedMemoryBytes int64                     `json:"usedMemoryBytes,omitempty"`
	UsedTimeMillis  int64                     `json:"usedTimeMillis,omitempty"`
	Verdict         string                    `json:"verdict"`
}

type NeuripsSubmissionsReportResponse struct {
	Submissions []*NeuripsSubmissionReport `json:"submissions"`
	TotalCount  int64                      `json:"totalCount"`
}

type ParticipantInfo struct {
	ID        int64  `json:"id,omitempty"`
	Login     string `json:"login,omitempty"`
	Name      string `json:"name,omitempty"`
	StartTime string `json:"startTime,omitempty"`
	UID       string `json:"uid,omitempty"`
}

type ParticipantStats struct {
	ContestID           int64             `json:"contestId"`
	FirstSubmissionTime string            `json:"firstSubmissionTime,omitempty"`
	ID                  int64             `json:"id"`
	Login               string            `json:"login"`
	Name                string            `json:"name"`
	Runs                []*BriefRunReport `json:"runs"`
	StartedAt           string            `json:"startedAt,omitempty"`
}

type ParticipantStatus struct {
	ContestDuration       string `json:"contestDuration,omitempty"`   // Example: PT4H30M
	ContestFinishTime     string `json:"contestFinishTime,omitempty"` // Example: 2022-01-31T09:00:00.000Z
	ContestInfinite       bool   `json:"contestInfinite,omitempty"`
	ContestStartTime      string `json:"contestStartTime,omitempty"`      // Example: 2021-12-16T17:00:00.000Z
	ContestState          string `json:"contestState,omitempty"`          // Enum: ["FINISHED","NOT_STARTED","IN_PROGRESS"]
	ParticipantFinishTime string `json:"participantFinishTime,omitempty"` // Example: 2022-01-01T16:30:00.000Z
	// Deprecated. Switch to `participantLeftTimeMillis`.
	ParticipantLeftTime       string   `json:"participantLeftTime,omitempty"`       // Example: PT5M4.123S
	ParticipantLeftTimeMillis int64    `json:"participantLeftTimeMillis,omitempty"` // Example: 304123
	ParticipantName           string   `json:"participantName,omitempty"`
	ParticipantStartTime      string   `json:"participantStartTime,omitempty"` // Example: 2022-01-01T12:00:00.000Z
	Roles                     []string `json:"roles"`
	// This property is `null` unless contestant is participating in a team.
	TeamID           int64 `json:"teamId,omitempty"` // Example: 10
	UpsolvingAllowed bool  `json:"upsolvingAllowed,omitempty"`
}

type ProblemResult struct {
	Score           string `json:"score,omitempty"`
	Status          string `json:"status,omitempty"`
	SubmissionCount string `json:"submissionCount,omitempty"`
	SubmitDelay     int64  `json:"submitDelay,omitempty"`
}

type PublicCompilerInfo struct {
	Deprecated bool   `json:"deprecated"`
	ID         string `json:"id"`
	Name       string `json:"name"`
	Style      string `json:"style"`
}

type RegisterGroupRequest struct {
	Roles []string `json:"roles"`
}

type RunID struct {
	RunID int64 `json:"runId"`
}

type ServiceCapacity struct {
	SubmissionsCount int64 `json:"submissionsCount"`
}

type Statement struct {
	Locale string `json:"locale,omitempty"`
	Path   string `json:"path,omitempty"`
	Type   string `json:"type,omitempty"` // Enum: ["TEX","PDF","BINARY","MARKDOWN"]
}

type Submission struct {
	Author         string  `json:"author"`
	AuthorID       int64   `json:"authorId"`
	Compiler       string  `json:"compiler"`
	ID             int64   `json:"id"`
	Memory         int64   `json:"memory"`
	ProblemAlias   string  `json:"problemAlias"`
	ProblemID      string  `json:"problemId"`
	Score          float64 `json:"score,omitempty"`
	SubmissionTime string  `json:"submissionTime"`
	Test           int64   `json:"test"`
	Time           int64   `json:"time"`
	// Time of submission, in milliseconds from participant start time
	TimeFromStart int64  `json:"timeFromStart,omitempty"`
	Verdict       string `json:"verdict"`
}

type Submissions struct {
	Count       int64         `json:"count"`
	Submissions []*Submission `json:"submissions"`
}

type SubmitInfo struct {
	ParticipantID   int64  `json:"participantId"`
	ParticipantName string `json:"participantName"`
	ProblemTitle    string `json:"problemTitle"`
	SubmitTime      int64  `json:"submitTime"`
}

type TeamView struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type TestLog struct {
	Answer         string `json:"answer"`
	CheckerError   string `json:"checkerError"`
	Error          string `json:"error"`
	Input          string `json:"input"`
	IsSample       bool   `json:"isSample"`
	MemoryUsed     int64  `json:"memoryUsed"`
	Message        string `json:"message"`
	Output         string `json:"output"`
	RunningTime    int64  `json:"runningTime"`
	SequenceNumber int64  `json:"sequenceNumber"`
	TestName       string `json:"testName"`
	TestsetIdx     int64  `json:"testsetIdx"`
	Verdict        string `json:"verdict"`
}

type TokenInfo struct {
	Active bool   `json:"active"`
	Scope  string `json:"scope"`
}

type UpdateGroupParticipationRequest struct {
	Roles []string `json:"roles"`
}

type UpdateParticipantRequest struct {
	DisplayedName string `json:"displayedName,omitempty"`
}

type UserIdentifier struct {
	Login string `json:"login,omitempty"`

	UID int64 `json:"uid,omitempty"`
}

type UserWithPasswordResponse struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	UserID   int64  `json:"userId"`
}
