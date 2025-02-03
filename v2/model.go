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
	Verdict      string `json:"verdict,omitempty"`
}

type BriefRunReport struct {
	AuthorId       string  `json:"authorId"`
	CompileLog     string  `json:"compileLog"`
	Compiler       string  `json:"compiler"`
	Diff           string  `json:"diff,omitempty"`
	MaxMemoryUsage int64   `json:"maxMemoryUsage"`
	MaxTimeUsage   int64   `json:"maxTimeUsage"`
	ProblemAlias   string  `json:"problemAlias"`
	ProblemId      string  `json:"problemId"`
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

type PublicCompilerInfo struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Deprecated bool   `json:"deprecated"`
	Style      string `json:"style"`
}

type CompilerListResponse struct {
	Compilers []*PublicCompilerInfo `json:"compilers"`
}

type ContestDescription struct {
	Duration           int64                  `json:"duration,omitempty"`
	FreezeTime         int64                  `json:"freezeTime,omitempty"`
	Name               string                 `json:"name"`
	StartTime          string                 `json:"startTime,omitempty"`
	Type               string                 `json:"type,omitempty"`
	UpsolvingAllowance string                 `json:"upsolvingAllowance,omitempty"`
	ContestSettings    *ContestReportSettings `json:"contestSettings,omitempty"`
}

type ContestReportSettings struct {
	ShowAnswer               string `json:"showAnswer"`
	ShowCheckerOutput        string `json:"showCheckerOutput"`
	ShowInput                string `json:"showInput"`
	ShowLightweightInterface bool   `json:"showLightweightInterface"`
	ShowOutput               string `json:"showOutput"`
	ShowPostprocessOutput    string `json:"showPostprocessOutput"`
	ShowReport               bool   `json:"showReport"`
	ShowScore                string `json:"showScore"`
	ShowStderr               string `json:"showStderr"`
	ShowTestNumber           bool   `json:"showTestNumber"`
	ShowUsedResources        string `json:"showUsedResources"`
	ShowVerdict              string `json:"showVerdict"`
	StopOnFirstFail          bool   `json:"stopOnFirstFail"`
	StopOnFirstFailInTestSet bool   `json:"stopOnFirstFailInTestSet"`
	StopOnSampleFail         bool   `json:"stopOnSampleFail"`
	TestOnlySamples          bool   `json:"testOnlySamples"`
	YseAcNotOk               bool   `json:"useAcNotOk"`
}

type ContestProblem struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Alias       string           `json:"alias"`
	Compilers   []string         `json:"compilers"`
	Statements  []*Statement     `json:"statements"`
	Limits      []*CompilerLimit `json:"limits"`
	TestCount   int64            `json:"testCount"`
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

type ContestStatistics struct {
	LastSubmit  *SubmitInfo `json:"lastSubmit"`
	LastSuccess *SubmitInfo `json:"lastSuccess"`
}

type SubmitInfo struct {
	ParticipantId   int64  `json:"participantId"`
	ParticipantName string `json:"participantName"`
	ProblemTitle    string `json:"problemTitle"`
	SubmitTime      int64  `json:"submitTime"`
}

type ContestStandingsTitle struct {
	Name  string `json:"name,omitempty"`
	Title string `json:"title"`
}

type FullRunReport struct {
	CheckerLog       []*TestLog                       `json:"checkerLog"`
	CompileLog       string                           `json:"compileLog"`
	Compiler         string                           `json:"compiler"`
	ContestID        int64                            `json:"contestId"`
	ContestName      string                           `json:"contestName"`
	Diff             string                           `json:"diff"`
	FinalScore       string                           `json:"finalScore"`
	IP               string                           `json:"ip,omitempty"`
	MaxMemoryUsage   int64                            `json:"maxMemoryUsage"`
	MaxTimeUsage     int64                            `json:"maxTimeUsage"`
	Participant      *ParticipantInfo                 `json:"participant"`
	PrecompileChecks []*AdditionalSolutionCheckResult `json:"precompileChecks"`
	PreliminaryScore string                           `json:"preliminaryScore"`
	ProblemAlias     string                           `json:"problemAlias"`
	ProblemID        string                           `json:"problemId"`
	ProblemName      string                           `json:"problemName"`
	RunID            int64                            `json:"runId"`
	Source           string                           `json:"source"`
	Status           string                           `json:"status"`
	SubmissionTime   string                           `json:"submissionTime"`
	TestFileType     string                           `json:"testFileType"`
	TimeFromStart    int64                            `json:"timeFromStart"`
	Verdict          string                           `json:"verdict"`
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

type ContestStandingsRow struct {
	ParticipantInfo *ParticipantInfo `json:"participantInfo,omitempty"`
	PlaceFrom       *int64           `json:"placeFrom,omitempty"`
	PlaceTo         *int64           `json:"placeTo,omitempty"`
	ProblemResults  []*ProblemResult `json:"problemResults"`
	Score           string           `json:"score,omitempty"`
}

type GroupInfo struct {
	Name        string `json:"name,omitempty"`
	MemberCount int64  `json:"memberCount,omitempty"`
}

type ContestGroup struct {
	ID    int64    `json:"id"`
	Name  string   `json:"name"`
	Roles []string `json:"roles,omitempty"`
}

type CreateGroupRequest struct {
	Name string `json:"name"`
}

type GroupID struct {
	ID int64 `json:"id"`
}

type Service struct {
	Scope  string `json:"scope,omitempty"`
	Active bool   `json:"active,omitempty"`
}

type InputStream struct {
}

type LazySubmissionRequest struct {
	FileUrl      string `json:"fileUrl"`
	FileName     string `json:"fileName"`
	ProblemAlias string `json:"problemAlias"`
	CompilerId   string `json:"compilerId"`
	Meta         string `json:"meta,omitempty"`
}

type MultiRunReport struct {
	CompileLog           string                           `json:"compileLog"`
	Compiler             string                           `json:"compiler"`
	ContestID            int64                            `json:"contestId"`
	ContestName          string                           `json:"contestName"`
	FinalScore           string                           `json:"finalScore"`
	Ip                   string                           `json:"ip,omitempty"`
	MaxMemoryUsage       int64                            `json:"maxMemoryUsage"`
	MaxTimeUsage         int64                            `json:"MaxTimeUsage"`
	ParticipantInfo      *ParticipantInfo                 `json:"ParticipantInfo"`
	PostprocessorMessage string                           `json:"postprocessorMessage,omitempty"`
	PrecompileChecks     *[]AdditionalSolutionCheckResult `json:"precompileChecks"`
	PreliminaryScore     string                           `json:"preliminaryScore"`
	ProblemAlias         string                           `json:"problemAlias"`
	ProblemID            string                           `json:"problemId"`
	RunId                int64                            `json:"runId"`
	Status               string                           `json:"status,omitempty"`
	SubmissionTime       string                           `json:"submissionTime"`
	TestFileType         string                           `json:"testFileType"`
	Tests                *[]TestLog                       `json:"tests"`
	TimeFromStart        int64                            `json:"TimeFromStart"`
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
	Status              string                   `json:"status,omitempty"`
	TimeFromStartMillis int64                    `json:"timeFromStartMillis"`
}

type NeuripsSubmissionResult struct {
	Details         *NeuripsSubmissionDetails `json:"details,omitempty"`
	Score           float64                   `json:"score,omitempty"` // number(double)?
	UsedMemoryBytes int64                     `json:"usedMemoryBytes,omitempty"`
	UsedTimeMillis  int64                     `json:"UsedTimeMillis,omitempty"`
	Verdict         string                    `json:"verdict"`
}

type NeuripsSubmissionsReportResponse struct {
	Submissions *[]NeuripsSubmissionReport `json:"submissions"`
	TotalCount  *int64                     `json:"totalCount"`
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
	ContestDuration       string `json:"contestDuration,omitempty"`
	ContestFinishTime     string `json:"contestFinishTime,omitempty"`
	ContestInfinite       bool   `json:"contestInfinite,omitempty"`
	ContestStartTime      string `json:"contestStartTime,omitempty"`
	ContestState          string `json:"contestState,omitempty"`
	ParticipantFinishTime string `json:"participantFinishTime,omitempty"`
	// Deprecated. Switch to participantLeftTimeMillis.
	ParticipantLeftTime       string   `json:"participantLeftTime,omitempty"`
	ParticipantLeftTimeMillis int64    `json:"ParticipantLeftTimeMillis,omitempty"`
	ParticipantName           string   `json:"participantName,omitempty"`
	ParticipantStartTime      string   `json:"participantStartTime,omitempty"`
	Roles                     []string `json:"roles"`
	// This property is null unless contestant is participating in a team.
	TeamID           int64 `json:"teamId,omitempty"`
	UpsolvingAllowed bool  `json:"upsolvingAllowed,omitempty"`
}

type ProblemResult struct {
	Score           string `json:"score,omitempty"`
	Status          string `json:"status,omitempty"`
	SubmissionCount string `json:"submissionCount,omitempty"`
	SubmitDelay     int64  `json:"submitDelay,omitempty"`
}

type RegisterGroupRequest struct {
	Roles *[]string `json:"roles"`
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
	Type   string `json:"type,omitempty"`
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
	// time of submission, in milliseconds from participant start time
	TimeFromStart int64  `json:"timeFromStart,omitempty"`
	Verdict       string `json:"verdict"`
}

type Submissions struct {
	Count       int64         `json:"count"`
	Submissions []*Submission `json:"submissions"`
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

type UserWithPasswordResponse struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	UserID   int64  `json:"userId"`
}

type JSONNode struct {
	Array               bool   `json:"array,omitempty"`
	BigDecimal          bool   `json:"bigDecimal,omitempty"`
	BigInteger          bool   `json:"bigInteger,omitempty"`
	Binary              bool   `json:"binary,omitempty"`
	Boolean             bool   `json:"boolean,omitempty"`
	ContainerNode       bool   `json:"containerNode,omitempty"`
	Double              bool   `json:"double,omitempty"`
	Empty               bool   `json:"empty,omitempty"`
	Float               bool   `json:"float,omitempty"`
	FloatingPointNumber bool   `json:"floatingPointNumber,omitempty"`
	Int                 bool   `json:"int,omitempty"`
	IntegralNumber      bool   `json:"integralNumber,omitempty"`
	Long                bool   `json:"long,omitempty"`
	MissingNode         bool   `json:"missingNode,omitempty"`
	NodeType            string `json:"nodeType,omitempty"`
	Null                bool   `json:"null,omitempty"`
	Number              bool   `json:"number,omitempty"`
	Object              bool   `json:"object,omitempty"`
	Pojo                bool   `json:"pojo,omitempty"`
	Short               bool   `json:"short,omitempty"`
	Textual             bool   `json:"textual,omitempty"`
	ValueNode           bool   `json:"valueNode,omitempty"`
}
