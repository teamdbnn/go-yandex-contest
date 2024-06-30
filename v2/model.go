package yacontest

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"io"
)

// --------------------------
type AdditionalSolutionCheckResult struct {
	CheckerName  string `json:"checkerName,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
	Message      string `json:"message,omitempty"`
	Verdict      string `json:"verdict,omitempty"`
}

type BriefRunReport struct {
	CompileLog     *string `json:"compileLog"`
	Compiler       *string `json:"compiler"`
	Diff           *string `json:"diff"`
	ProblemAlias   *string `json:"problemAlias"`
	ProblemId      *string `json:"problemId"`
	Score          float64 `json:"score"`
	Source         *string `json:"source"`
	SubmissionTime *string `json:"submissionTime"`
	TimeFromStart  *int64  `json:"timeFromStart"`
	Verdict        *string `json:"verdict"`
}

type Clarification struct {
	Message string `json:"message,omitempty"`
	Subject string `json:"subject,omitempty"`
}

type Clarifications struct {
	Clarifications []*Clarification `json:"clarifications"`
}

type CompetitionRegisterResponse struct {
	MissedUids []string `json:"missedUids"`
}

type CompilerLimit struct {
	CompilerName  string `json:"compilerName,omitempty"`
	IdlenessLimit int64  `json:"idlenessLimit,omitempty"`
	MemoryLimit   int64  `json:"memoryLimit,omitempty"`
	OutputLimit   int64  `json:"outputLimit,omitempty"`
	TimeLimit     int64  `json:"timeLimit,omitempty"`
}
type PublicCompilerInfo struct {
	ID         *string `json:"id"`
	Name       *string `json:"name"`
	Deprecated *bool   `json:"deprecated"`
	Style      *string `json:"style"`
}
type CompilerListResponse struct {
	Compilers []*PublicCompilerInfo `json:"compilers"`
}

type ContestDescription struct {
	Duration           int64                  `json:"duration,omitempty"`
	FreezeTime         int64                  `json:"freezeTime,omitempty"`
	Name               *string                `json:"name"`
	StartTime          string                 `json:"startTime,omitempty"`
	Type               string                 `json:"type,omitempty"`
	UpsolvingAllowance string                 `json:"upsolvingAllowance,omitempty"`
	ContestSettings    *ContestReportSettings `json:"contestSettings,omitempty"`
}

type ContestReportSettings struct {
	ShowAnswer               *string `json:"showAnswer"`
	ShowCheckerOutput        *string `json:"showCheckerOutput"`
	ShowInput                *string `json:"showInput"`
	showLightweightInterface *bool   `json:"showLightweightInterface"`
	ShowOutput               *string `json:"showOutput"`
	ShowPostprocessOutput    *string `json:"showPostprocessOutput"`
	ShowReport               *bool   `json:"showReport"`
	ShowScore                *string `json:"showScore"`
	ShowStderr               *string `json:"showStderr"`
	ShowTestNumber           *bool   `json:"showTestNumber"`
	ShowUsedResources        *string `json:"showUsedResources"`
	ShowVerdict              *string `json:"showVerdict"`
	StopOnFirstFail          *bool   `json:"stopOnFirstFail"`
	StopOnFirstFailInTestSet *bool   `json:"stopOnFirstFailInTestSet"`
	StopOnSampleFail         *bool   `json:"stopOnSampleFail"`
	TestOnlySamples          *bool   `json:"testOnlySamples"`
	YseAcNotOk               *bool   `json:"useAcNotOk"`
}

type ContestProblem struct {
	ID          *string          `json:"id"`
	Name        *string          `json:"name"`
	Alias       *string          `json:"alias"`
	Compilers   []string         `json:"compilers"`
	Statements  []*Statement     `json:"statements"`
	Limits      []*CompilerLimit `json:"limits"`
	TestCount   int32            `json:"testCount"`
	ProblemType *string          `json:"problemType"`
}

type ContestProblems struct {
	Problems []*ContestProblem `json:"problems"`
}

type ContestSettings struct {
	ContestName string     `json:"contestName,omitempty"`
	ContestType string     `json:"contestType,omitempty"`
	Duration    string     `json:"duration,omitempty"`
	Languages   []Language `json:"languages"`
}

type ContestStandings struct {
	Rows       []*ContestStandingsRow `json:"rows"`
	Titles     []*RowTitle            `json:"titles"`
	Statistics *ContestStatistics     `json:"statistics"`
}

type ContestStatistics struct {
	LastSubmit  *SubmitInfo `json:"lastSubmit"`
	LastSuccess *SubmitInfo `json:"lastSuccess"`
}

type SubmitInfo struct {
	ParticipantId   *int64  `json:"participantId"`
	ParticipantName *string `json:"participantName"`
	ProblemTitle    *string `json:"problemTitle"`
	SubmitTime      *int64  `json:"submitTime"`
}

type ContestStandingsTitle struct {
	Name  string  `json:"name,omitempty"`
	Title *string `json:"title"`
}

type FullRunReport struct {
	CheckerLog       []*TestLog                       `json:"checkerLog"`
	CompileLog       *string                          `json:"compileLog"`
	Compiler         *string                          `json:"compiler"`
	ContestID        *int64                           `json:"contestId"`
	ContestName      *string                          `json:"contestName"`
	Diff             *string                          `json:"diff"`
	FinalScore       *string                          `json:"finalScore"`
	IP               string                           `json:"ip,omitempty"`
	MaxMemoryUsage   *int64                           `json:"maxMemoryUsage"`
	MaxTimeUsage     *int64                           `json:"maxTimeUsage"`
	Participant      *ParticipantInfo                 `json:"participant"`
	PrecompileChecks []*AdditionalSolutionCheckResult `json:"precompileChecks"`
	PreliminaryScore *string                          `json:"preliminaryScore"`
	ProblemAlias     *string                          `json:"problemAlias"`
	ProblemID        *string                          `json:"problemId"`
	ProblemName      *string                          `json:"problemName"`
	RunID            *int64                           `json:"runId"`
	Source           *string                          `json:"source"`
	Status           *string                          `json:"status"`
	SubmissionTime   *string                          `json:"submissionTime"`
	TestFileType     *string                          `json:"testFileType"`
	TimeFromStart    *int64                           `json:"timeFromStart"`
	Verdict          *string                          `json:"verdict"`
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
	PlaceFrom       *int32           `json:"placeFrom,omitempty"`
	PlaceTo         *int32           `json:"placeTo,omitempty"`
	ProblemResults  []*ProblemResult `json:"problemResults"`
	Score           *string          `json:"score,omitempty"`
}

type GroupInfo struct {
	Name        *string `json:"name,omitempty"`
	MemberCount *int64  `json:"memberCount,omitempty"`
}

type ContestGroup struct {
	ID    *int64   `json:"id"`
	Name  *string  `json:"name"`
	Roles []string `json:"roles,omitempty"`
}

type CreateGroupRequest struct {
	Name *string `json:"name"`
}

type GroupId struct {
	ID *int64 `json:"id"`
}

type Service struct {
	Scope  string `json:"scope,omitempty"`
	Active bool   `json:"active,omitempty"`
}

type InputStream struct {
}

type LazySubmissionRequest struct {
	FileUrl      *string `json:"fileUrl"`
	FileName     *string `json:"fileName"`
	ProblemAlias *string `json:"problemAlias"`
	CompilerId   *string `json:"compilerId"`
	Meta         *string `json:"meta,omitempty"`
}

type MultiRunReport struct {
	CompileLog           *string                          `json:"compileLog"`
	Compiler             *string                          `json:"compiler"`
	ContestID            *int64                           `json:"contestId"`
	ContestName          *string                          `json:"contestName"`
	FinalScore           *string                          `json:"finalScore"`
	Ip                   string                           `json:"ip,omitempty"`
	MaxMemoryUsage       *int64                           `json:"maxMemoryUsage"`
	MaxTimeUsage         *int64                           `json:"MaxTimeUsage"`
	ParticipantInfo      *ParticipantInfo                 `json:"ParticipantInfo"`
	PostprocessorMessage string                           `json:"postprocessorMessage,omitempty"`
	PrecompileChecks     *[]AdditionalSolutionCheckResult `json:"precompileChecks"`
	PreliminaryScore     *string                          `json:"preliminaryScore"`
	ProblemAlias         *string                          `json:"problemAlias"`
	ProblemID            *string                          `json:"problemId"`
	RunId                *int64                           `json:"runId"`
	Status               string                           `json:"status,omitempty"`
	SubmissionTime       *string                          `json:"submissionTime"`
	TestFileType         *string                          `json:"testFileType"`
	Tests                *[]TestLog                       `json:"tests"`
	TimeFromStart        *int64                           `json:"TimeFromStart"`
	Verdict              *string                          `json:"verdict"`
}

type NeuripsSubmissionDetails struct {
	Metrics JSONNode `json:"metrics,omitempty"`
}

type NeuripsSubmissionReport struct {
	Compiler            *string                 `json:"compiler"`
	ContestID           *int64                  `json:"contestId"`
	CreateAt            *strfmt.DateTime        `json:"createAt"`
	Meta                string                  `json:"meta,omitempty"`
	ParticipantName     *string                 `json:"participantName"`
	ProblemID           *string                 `json:"problemId"`
	Rank                *int32                  `json:"rank"`
	Result              NeuripsSubmissionResult `json:"result,omitempty"`
	RunID               *int64                  `json:"runId"`
	Status              string                  `json:"status,omitempty"`
	TimeFromStartMillis *int64                  `json:"timeFromStartMillis"`
}

type NeuripsSubmissionResult struct {
	Details         NeuripsSubmissionDetails `json:"details,omitempty"`
	Score           float64                  `json:"score,omitempty"` // number(double)?
	UsedMemoryBytes int64                    `json:"usedMemoryBytes,omitempty"`
	UsedTimeMillis  int64                    `json:"UsedTimeMillis,omitempty"`
	Verdict         *string                  `json:"verdict"`
}

type NeuripsSubmissionsReportResponse struct {
	Submissions *[]NeuripsSubmissionReport `json:"submissions"`
	TotalCount  *int32                     `json:"totalCount"`
}

type ParticipantInfo struct {
	ID        int64  `json:"id,omitempty"`
	Login     string `json:"login,omitempty"`
	Name      string `json:"name,omitempty"`
	StartTime string `json:"startTime,omitempty"`
	UID       string `json:"uid,omitempty"`
}

type ParticipantStats struct {
	ContestID           *int64           `json:"contestId"`
	FirstSubmissionTime string           `json:"firstSubmissionTime,omitempty"`
	ID                  *int64           `json:"id"`
	Login               *string          `json:"login"`
	Name                *string          `json:"name"`
	Runs                []*FullRunReport `json:"runs"`
	StartedAt           string           `json:"startedAt,omitempty"`
}

type ParticipantStatus struct {
	ContestDuration           string   `json:"contestDuration,omitempty"`
	ContestFinishTime         string   `json:"contestFinishTime,omitempty"`
	ContestInfinite           bool     `json:"contestInfinite,omitempty"`
	ContestStartTime          string   `json:"contestStartTime,omitempty"`
	ContestState              string   `json:"contestState,omitempty"`
	ParticipantFinishTime     string   `json:"participantFinishTime,omitempty"`
	ParticipantLeftTime       string   `json:"participantLeftTime,omitempty"`
	ParticipantLeftTimeMillis int64    `json:"ParticipantLeftTimeMillis,omitempty"`
	ParticipantName           string   `json:"participantName,omitempty"`
	ParticipantStartTime      string   `json:"participantStartTime,omitempty"`
	Roles                     []string `json:"roles"`
	TeamID                    int64    `json:"teamId,omitempty"`
	UpsolvingAllowed          bool     `json:"upsolvingAllowed,omitempty"`
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

type RunId struct {
	RunID *int64 `json:"runId"`
}

type ServiceCapacity struct {
	SubmissionsCount *int64 `json:"submissionsCount"`
}

type Statement struct {
	Locale string `json:"locale,omitempty"`
	Path   string `json:"path,omitempty"`
	Type   string `json:"type,omitempty"`
}

type Submission struct {
	Author         *string `json:"author"`
	AuthorID       *int64  `json:"authorId"`
	Compiler       *string `json:"compiler"`
	ID             *int64  `json:"id"`
	Memory         *int64  `json:"memory"`
	ProblemAlias   *string `json:"problemAlias"`
	ProblemID      *string `json:"problemId"`
	Score          float64 `json:"score,omitempty"` // !!!
	SubmissionTime *string `json:"submissionTime"`
	Test           *int32  `json:"test"`
	Time           *int64  `json:"time"`
	TimeFromStart  int64   `json:"timeFromStart,omitempty"`
	Verdict        *string `json:"verdict"`
}

type Submissions struct {
	Count       *int32        `json:"count"`
	Submissions []*Submission `json:"submissions"`
}

type TeamView struct {
	ID   *int64  `json:"id"`
	Name *string `json:"name"`
}

type TestLog struct {
	Answer         *string `json:"answer"`
	CheckerError   *string `json:"checkerError"`
	Error          *string `json:"error"`
	Input          *string `json:"input"`
	IsSample       *bool   `json:"isSample"`
	MemoryUsed     *int64  `json:"memoryUsed"`
	Message        *string `json:"message"`
	Output         *string `json:"output"`
	RunningTime    *int64  `json:"runningTime"`
	SequenceNumber *int32  `json:"sequenceNumber"`
	TestName       *string `json:"testName"`
	TestsetIdx     *int32  `json:"testsetIdx"`
	Verdict        *string `json:"verdict"`
}

type TokenInfo struct {
	Active *bool   `json:"active"`
	Scope  *string `json:"scope"`
}

type UpdateGroupParticipantRequest struct {
	Roles *[]string `json:"roles"`
}

type UpdateParticipantRequest struct {
	DisplayedName string `json:"displayedName,omitempty"`
}

type UserIdentifier struct {
	Login string `json:"login,omitempty"`
	UID   int64  `json:"uid,omitempty"`
}

type UserWithPasswordResponse struct {
	Login    *string `json:"login"`
	Password *string `json:"password"`
	UserID   *int64  `json:"userId"`
}

// ----------------------------

type AccessLevel struct {
	ID                                 *int64  `json:"id"`
	CheckerSetingsModificationAllowed  *bool   `json:"checkerSettingsModificationAllowed"`
	ContestLimit                       int32   `json:"contestLimit,omitempty"`
	FileSizeUploadLimit                int64   `json:"fileSizeUploadLimit,omitempty"`
	IsGroupCreationAllowed             *bool   `json:"isGroupCreationAllowed"`
	IsProblemsetModificationAllowed    *bool   `json:"isProblemsetModificationAllowed"`
	Name                               *string `json:"name"`
	ParticipantLimit                   int32   `json:"participantLimit,omitempty"`
	ProblemLimit                       int32   `json:"problemLimit,omitempty"`
	ProblemsetLimit                    int32   `json:"problemsetLimit,omitempty"`
	TestsetTemplateModificationAllowed *bool   `json:"testsetTemplateModificationAllowed"`
	TotalSizeUploadLimit               int64   `json:"totalSizeUploadLimit,omitempty"`
}

type AdditionalSolutionCheck struct {
	CheckerFilesField FileSystemPackage `json:"checkerFilesField,omitempty"`
	CompilerID        string            `json:"compilerId,omitempty"`
	MainFilename      string            `json:"mainFilename,omitempty"`
}

type Annotation interface{}

type CompetitionParticipantsRegisterResponse struct {
	MissedUids []string `json:"missedUids"`
}

type Compiler struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Deprecated bool   `json:"deprecated"`
	Style      string `json:"style,omitempty"`
}

type Compilers struct {
	Compilers []*Compiler `json:"compilers"`
}

type Contest struct {
	ID              *int64                  `json:"id"`
	Compilers       []string                `json:"compilers"`
	Duration        *Duration               `json:"duration,omitempty"`
	Enabled         bool                    `json:"enabled,omitempty"`
	EndTime         strfmt.DateTime         `json:"endTime,omitempty"`
	Finished        bool                    `json:"finished,omitempty"`
	Infinite        bool                    `json:"infinite,omitempty"`
	Monitor         *MonitorConfiguration   `json:"monitor,omitempty"`
	MonitorPlugin   string                  `json:"monitorPlugin,omitempty"`
	Name            string                  `json:"name,omitempty"`
	Owner           *User                   `json:"owner,omitempty"`
	ProblemSetID    string                  `json:"problemSetId,omitempty"`
	StartTime       strfmt.DateTime         `json:"startTime,omitempty"`
	TestingSettings *ContestTestingSettings `json:"testingSettings,omitempty"`
	TimeControlType string                  `json:"timeControlType,omitempty"`
	TimeLimited     bool                    `json:"timeLimited,omitempty"`
}

type ContestLog struct {
	Events         []*Event         `json:"events"`
	GenerationTime strfmt.DateTime  `json:"generationTime,omitempty"`
	Problems       []Problem        `json:"problems"`
	Settings       *ContestSettings `json:"settings,omitempty"`
	Users          []*User          `json:"users"`
}

type Statistic struct {
	LastSubmit  []*LastSuccessOrSubmit `json:"lastSubmit"`
	LastSuccess []*LastSuccessOrSubmit `json:"lastSuccess"`
}

type LastSuccessOrSubmit struct {
	ParticipantId   *int64 `json:"participantId"`
	ParticipantName string `json:"participantName,omitempty"`
	ProblemTitle    string `json:"problemTitle,omitempty"`
	SubmitTime      int64  `json:"submitTime,omitempty"`
}

type ContestTestingSettings struct {
	Contest            *Contest                   `json:"contest,omitempty"`
	ID                 *int64                     `json:"id"`
	PrecompileCheckers []*AdditionalSolutionCheck `json:"precompileCheckers"`
}

type DetailedRunReport struct {
	CompileLog     *string `json:"compileLog"`
	Compiler       *string `json:"compiler"`
	Diff           *string `json:"diff"`
	ProblemAlias   *string `json:"problemAlias"`
	ProblemID      *string `json:"problemId"`
	Source         *string `json:"source"`
	SubmissionTime *string `json:"submissionTime"`
	Verdict        *string `json:"verdict"`
}

type DirectoryPackageAllOf1 struct {
	Name          string `json:"name,omitempty"`
	packagesField []FileSystemPackage
}

type DirectoryPackage struct {
	nameField string
	DirectoryPackageAllOf1
}

type Duration struct {
	Millis          int64 `json:"millis,omitempty"`
	StandardDays    int64 `json:"standardDays,omitempty"`
	StandardHours   int64 `json:"standardHours,omitempty"`
	StandardMinutes int64 `json:"standardMinutes,omitempty"`
	StandardSeconds int64 `json:"standardSeconds,omitempty"`
}

type Empty interface{}

type Event struct {
	AbsoluteTime int64 `json:"absoluteTime,omitempty"`
	ContestTime  int64 `json:"contestTime,omitempty"`
}

type ExtendedRole struct {
	Metas     []*RoleMeta `json:"metas"`
	Principal *Principal  `json:"principal"`
	Resource  *Resource   `json:"resource,omitempty"`
	Role      *Role       `json:"role"`
}

type ExternalID struct {
	ClientID *string `json:"clientId"`
	Empty    *bool   `json:"empty"`
	ID       string  `json:"id,omitempty"`
}

type ExternalLog struct {
	ContestLog        *ContestLog     `json:"contestLog,omitempty"`
	ContestLogKey     []string        `json:"contestLogKey"`
	DownloadStartTime strfmt.DateTime `json:"downloadStartTime,omitempty"`
	Dynamic           bool            `json:"dynamic,omitempty"`
	DynamicURL        string          `json:"dynamicUrl,omitempty"`
	EventsCount       int32           `json:"eventsCount,omitempty"`
	ID                *int64          `json:"id"`

	LastUpdate        strfmt.DateTime `json:"lastUpdate,omitempty"`
	LastUpdateAttempt strfmt.DateTime `json:"lastUpdateAttempt,omitempty"`
	LastUpdateError   []string        `json:"lastUpdateError"`
	OwnerID           []int64         `json:"ownerId"`
	RawXMLLog         strfmt.Base64   `json:"rawXmlLog,omitempty"`
	Title             string          `json:"title,omitempty"`
	XMLLogKey         []string        `json:"xmlLogKey"`
}

type FilePackageAllOf1 struct {
	Data io.ReadCloser `json:"data,omitempty"`
	Name string        `json:"name,omitempty"`
}

type FilePackage struct {
	nameField string
	FilePackageAllOf1
}

type FileSystemPackage interface {
	runtime.Validatable
	runtime.ContextValidatable
	Name() string
	SetName(string)
}
type fileSystemPackage struct {
	nameField string
}

func (m *fileSystemPackage) Name() string {
	return m.nameField
}
func (m *fileSystemPackage) SetName(val string) {
	m.nameField = val
}

type GrantResponse struct {
	Roles []*ExtendedRole `json:"roles"`
}

type HTTPEntity struct {
	Body interface{} `json:"body,omitempty"`
}

type IDPrincipal struct {
	CheckedID            *int64           `json:"checkedId"`
	CheckedRawExternalID *string          `json:"checkedRawExternalId"`
	ExternalID           *ExternalID      `json:"externalId"`
	ID                   int64            `json:"id,omitempty"`
	Model                *KClassPrincipal `json:"model"`
	New                  *bool            `json:"new"`
	RawExternalID        string           `json:"rawExternalId,omitempty"`
	UID                  strfmt.UUID      `json:"uid,omitempty"`
}

type IDResource struct {
	CheckedID            *int64          `json:"checkedId"`
	CheckedRawExternalID *string         `json:"checkedRawExternalId"`
	ExternalID           *ExternalID     `json:"externalId"`
	ID                   int64           `json:"id,omitempty"`
	Model                *KClassResource `json:"model"`
	New                  *bool           `json:"new"`
	RawExternalID        string          `json:"rawExternalId,omitempty"`
	UID                  strfmt.UUID     `json:"uid,omitempty"`
}

type IDRole struct {
	CheckedID            *int64      `json:"checkedId"`
	CheckedRawExternalID *string     `json:"checkedRawExternalId"`
	ExternalID           *ExternalID `json:"externalId"`
	ID                   int64       `json:"id,omitempty"`
	Model                *KClassRole `json:"model"`
	New                  *bool       `json:"new"`
	RawExternalID        string      `json:"rawExternalId,omitempty"`
	UID                  strfmt.UUID `json:"uid,omitempty"`
}

type IDRoleMeta struct {
	CheckedID            *int64          `json:"checkedId"`
	CheckedRawExternalID *string         `json:"checkedRawExternalId"`
	ExternalID           *ExternalID     `json:"externalId"`
	ID                   int64           `json:"id,omitempty"`
	Model                *KClassRoleMeta `json:"model"`
	New                  *bool           `json:"new"`
	RawExternalID        string          `json:"rawExternalId,omitempty"`
	UID                  strfmt.UUID     `json:"uid,omitempty"`
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

type KCallableObject struct {
	Abstract       *bool             `json:"abstract"`
	Annotations    []Annotation      `json:"annotations"`
	Final          *bool             `json:"final"`
	Name           *string           `json:"name"`
	Open           *bool             `json:"open"`
	Parameters     []*KParameter     `json:"parameters"`
	ReturnType     *KType            `json:"returnType"`
	Suspend        *bool             `json:"suspend"`
	TypeParameters []*KTypeParameter `json:"typeParameters"`
	Visibility     string            `json:"visibility,omitempty"`
}

type KClassifier interface{}

type KClassObject struct {
	Abstract         *bool              `json:"abstract"`
	Annotations      []Annotation       `json:"annotations"`
	Companion        *bool              `json:"companion"`
	Constructors     []*KFunctionObject `json:"constructors"`
	Data             *bool              `json:"data"`
	Final            *bool              `json:"final"`
	Inner            *bool              `json:"inner"`
	Members          []*KCallableObject `json:"members"`
	NestedClasses    []*KClassObject    `json:"nestedClasses"`
	ObjectInstance   interface{}        `json:"objectInstance,omitempty"`
	Open             *bool              `json:"open"`
	QualifiedName    string             `json:"qualifiedName,omitempty"`
	Sealed           *bool              `json:"sealed"`
	SealedSubclasses []*KClassObject    `json:"sealedSubclasses"`
	SimpleName       string             `json:"simpleName,omitempty"`
	Supertypes       []*KType           `json:"supertypes"`
	TypeParameters   []*KTypeParameter  `json:"typeParameters"`
	Visibility       string             `json:"visibility,omitempty"`
}

type KClassPrincipal struct {
	Abstract         *bool                 `json:"abstract"`
	Annotations      []Annotation          `json:"annotations"`
	Companion        *bool                 `json:"companion"`
	Constructors     []*KFunctionPrincipal `json:"constructors"`
	Data             *bool                 `json:"data"`
	Final            *bool                 `json:"final"`
	Inner            *bool                 `json:"inner"`
	Members          []*KCallableObject    `json:"members"`
	NestedClasses    []*KClassObject       `json:"nestedClasses"`
	ObjectInstance   *Principal            `json:"objectInstance,omitempty"`
	Open             *bool                 `json:"open"`
	QualifiedName    string                `json:"qualifiedName,omitempty"`
	Sealed           *bool                 `json:"sealed"`
	SealedSubclasses []*KClassPrincipal    `json:"sealedSubclasses"`
	SimpleName       string                `json:"simpleName,omitempty"`
	Supertypes       []*KType              `json:"supertypes"`
	TypeParameters   []*KTypeParameter     `json:"typeParameters"`
	Visibility       string                `json:"visibility,omitempty"`
}

type KClassResource struct {
	Abstract         *bool                `json:"abstract"`
	Annotations      []Annotation         `json:"annotations"`
	Companion        *bool                `json:"companion"`
	Constructors     []*KFunctionResource `json:"constructors"`
	Data             *bool                `json:"data"`
	Final            *bool                `json:"final"`
	Inner            *bool                `json:"inner"`
	Members          []*KCallableObject   `json:"members"`
	NestedClasses    []*KClassObject      `json:"nestedClasses"`
	ObjectInstance   *Resource            `json:"objectInstance,omitempty"`
	Open             *bool                `json:"open"`
	QualifiedName    string               `json:"qualifiedName,omitempty"`
	Sealed           *bool                `json:"sealed"`
	SealedSubclasses []*KClassResource    `json:"sealedSubclasses"`
	SimpleName       string               `json:"simpleName,omitempty"`
	Supertypes       []*KType             `json:"supertypes"`
	TypeParameters   []*KTypeParameter    `json:"typeParameters"`
	Visibility       string               `json:"visibility,omitempty"`
}

type KClassRole struct {
	Abstract         *bool              `json:"abstract"`
	Annotations      []Annotation       `json:"annotations"`
	Companion        *bool              `json:"companion"`
	Constructors     []*KFunctionRole   `json:"constructors"`
	Data             *bool              `json:"data"`
	Final            *bool              `json:"final"`
	Inner            *bool              `json:"inner"`
	Members          []*KCallableObject `json:"members"`
	NestedClasses    []*KClassObject    `json:"nestedClasses"`
	ObjectInstance   *Role              `json:"objectInstance,omitempty"`
	Open             *bool              `json:"open"`
	QualifiedName    string             `json:"qualifiedName,omitempty"`
	Sealed           *bool              `json:"sealed"`
	SealedSubclasses []*KClassRole      `json:"sealedSubclasses"`
	SimpleName       string             `json:"simpleName,omitempty"`
	Supertypes       []*KType           `json:"supertypes"`
	TypeParameters   []*KTypeParameter  `json:"typeParameters"`
	Visibility       string             `json:"visibility,omitempty"`
}

type KClassRoleMeta struct {
	Abstract         *bool                `json:"abstract"`
	Annotations      []Annotation         `json:"annotations"`
	Companion        *bool                `json:"companion"`
	Constructors     []*KFunctionRoleMeta `json:"constructors"`
	Data             *bool                `json:"data"`
	Final            *bool                `json:"final"`
	Inner            *bool                `json:"inner"`
	Members          []*KCallableObject   `json:"members"`
	NestedClasses    []*KClassObject      `json:"nestedClasses"`
	ObjectInstance   *RoleMeta            `json:"objectInstance,omitempty"`
	Open             *bool                `json:"open"`
	QualifiedName    string               `json:"qualifiedName,omitempty"`
	Sealed           *bool                `json:"sealed"`
	SealedSubclasses []*KClassRoleMeta    `json:"sealedSubclasses"`
	SimpleName       string               `json:"simpleName,omitempty"`
	Supertypes       []*KType             `json:"supertypes"`
	TypeParameters   []*KTypeParameter    `json:"typeParameters"`
	Visibility       string               `json:"visibility,omitempty"`
}

type KFunctionObject struct {
	Abstract       *bool             `json:"abstract"`
	Annotations    []Annotation      `json:"annotations"`
	External       *bool             `json:"external"`
	Final          *bool             `json:"final"`
	Infix          *bool             `json:"infix"`
	Inline         *bool             `json:"inline"`
	Name           *string           `json:"name"`
	Open           *bool             `json:"open"`
	Operator       *bool             `json:"operator"`
	Parameters     []*KParameter     `json:"parameters"`
	ReturnType     *KType            `json:"returnType"`
	Suspend        *bool             `json:"suspend"`
	TypeParameters []*KTypeParameter `json:"typeParameters"`
	Visibility     string            `json:"visibility,omitempty"`
}

type KFunctionPrincipal struct {
	Abstract       *bool             `json:"abstract"`
	Annotations    []Annotation      `json:"annotations"`
	External       *bool             `json:"external"`
	Final          *bool             `json:"final"`
	Infix          *bool             `json:"infix"`
	Inline         *bool             `json:"inline"`
	Name           *string           `json:"name"`
	Open           *bool             `json:"open"`
	Operator       *bool             `json:"operator"`
	Parameters     []*KParameter     `json:"parameters"`
	ReturnType     *KType            `json:"returnType"`
	Suspend        *bool             `json:"suspend"`
	TypeParameters []*KTypeParameter `json:"typeParameters"`
	Visibility     string            `json:"visibility,omitempty"`
}

type KFunctionResource struct {
	Abstract       *bool             `json:"abstract"`
	Annotations    []Annotation      `json:"annotations"`
	External       *bool             `json:"external"`
	Final          *bool             `json:"final"`
	Infix          *bool             `json:"infix"`
	Inline         *bool             `json:"inline"`
	Name           *string           `json:"name"`
	Open           *bool             `json:"open"`
	Operator       *bool             `json:"operator"`
	Parameters     []*KParameter     `json:"parameters"`
	ReturnType     *KType            `json:"returnType"`
	Suspend        *bool             `json:"suspend"`
	TypeParameters []*KTypeParameter `json:"typeParameters"`
	Visibility     string            `json:"visibility,omitempty"`
}

type KFunctionRole struct {
	Abstract       *bool             `json:"abstract"`
	Annotations    []Annotation      `json:"annotations"`
	External       *bool             `json:"external"`
	Final          *bool             `json:"final"`
	Infix          *bool             `json:"infix"`
	Inline         *bool             `json:"inline"`
	Name           *string           `json:"name"`
	Open           *bool             `json:"open"`
	Operator       *bool             `json:"operator"`
	Parameters     []*KParameter     `json:"parameters"`
	ReturnType     *KType            `json:"returnType"`
	Suspend        *bool             `json:"suspend"`
	TypeParameters []*KTypeParameter `json:"typeParameters"`
	Visibility     string            `json:"visibility,omitempty"`
}

type KFunctionRoleMeta struct {
	Abstract       *bool             `json:"abstract"`
	Annotations    []Annotation      `json:"annotations"`
	External       *bool             `json:"external"`
	Final          *bool             `json:"final"`
	Infix          *bool             `json:"infix"`
	Inline         *bool             `json:"inline"`
	Name           *string           `json:"name"`
	Open           *bool             `json:"open"`
	Operator       *bool             `json:"operator"`
	Parameters     []*KParameter     `json:"parameters"`
	ReturnType     *KType            `json:"returnType"`
	Suspend        *bool             `json:"suspend"`
	TypeParameters []*KTypeParameter `json:"typeParameters"`
	Visibility     string            `json:"visibility,omitempty"`
}

type KParameter struct {
	Annotations []Annotation `json:"annotations"`
	Index       *int32       `json:"index"`
	Kind        *string      `json:"kind"`
	Name        string       `json:"name,omitempty"`
	Optional    *bool        `json:"optional"`
	Type        *KType       `json:"type"`
	Vararg      *bool        `json:"vararg"`
}

type KType struct {
	Annotations    []Annotation       `json:"annotations"`
	Arguments      []*KTypeProjection `json:"arguments"`
	Classifier     KClassifier        `json:"classifier,omitempty"`
	MarkedNullable *bool              `json:"markedNullable"`
}

type KTypeParameter struct {
	Name        *string  `json:"name"`
	Reified     *bool    `json:"reified"`
	UpperBounds []*KType `json:"upperBounds"`
	Variance    *string  `json:"variance"`
}

type KTypeProjection struct {
	Type     *KType `json:"type,omitempty"`
	Variance string `json:"variance,omitempty"`
}

type Language interface{}

type MonitorConfiguration struct {
	ExternalLogs []*ExternalLog `json:"externalLogs"`
	ID           *int64         `json:"id"`
}

type Participant struct {
	Contest              *Contest        `json:"contest,omitempty"`
	ContestStarted       bool            `json:"contestStarted,omitempty"`
	CreationTime         []int64         `json:"creationTime"`
	Deleted              *bool           `json:"deleted"`
	DisplayedName        string          `json:"displayedName,omitempty"`
	DisplayedNameChecked string          `json:"displayedNameChecked,omitempty"`
	ID                   int64           `json:"id,omitempty"`
	Login                string          `json:"login,omitempty"`
	Name                 string          `json:"name,omitempty"`
	NameAutoGenerated    bool            `json:"nameAutoGenerated,omitempty"`
	NotDeleted           *bool           `json:"notDeleted"`
	StartTime            strfmt.DateTime `json:"startTime,omitempty"`
	Team                 *Team           `json:"team,omitempty"`
	TeamParticipation    bool            `json:"teamParticipation,omitempty"`
	Teammates            []*User         `json:"teammates"`
	User                 *User           `json:"user,omitempty"`
	UserID               int64           `json:"userId,omitempty"`
	UserIds              []int64         `json:"userIds"`
}

type Principal struct {
	CheckedID            int64            `json:"checkedId,omitempty"`
	CheckedRawExternalID string           `json:"checkedRawExternalId,omitempty"`
	CreatedAt            *strfmt.DateTime `json:"createdAt"`
	Deleted              *bool            `json:"deleted"`
	ID                   *IDPrincipal     `json:"id"`
	ModifiedAt           *strfmt.DateTime `json:"modifiedAt"`
	NotDeleted           *bool            `json:"notDeleted"`
	RawExternalID        string           `json:"rawExternalId,omitempty"`
	State                *State           `json:"state"`
	Title                *PrincipalTitle  `json:"title"`
}

type PrincipalTitle struct {
	Title *string `json:"title"`
}

type Problem interface{}

type Resource struct {
	CheckedID            int64            `json:"checkedId,omitempty"`
	CheckedRawExternalID string           `json:"checkedRawExternalId,omitempty"`
	CreatedAt            *strfmt.DateTime `json:"createdAt"`
	Deleted              *bool            `json:"deleted"`
	ID                   *IDResource      `json:"id"`
	ModifiedAt           *strfmt.DateTime `json:"modifiedAt"`
	NotDeleted           *bool            `json:"notDeleted"`
	RawExternalID        string           `json:"rawExternalId,omitempty"`
	State                *State           `json:"state"`
	Title                *ResourceTitle   `json:"title"`
}

type ResourceTitle struct {
	Title *string `json:"title"`
}

type ResponseEntity struct {
	Body            interface{} `json:"body,omitempty"`
	StatusCode      string      `json:"statusCode,omitempty"`
	StatusCodeValue int32       `json:"statusCodeValue,omitempty"`
}

type Role struct {
	CheckedID            int64            `json:"checkedId,omitempty"`
	CheckedRawExternalID string           `json:"checkedRawExternalId,omitempty"`
	CreatedAt            *strfmt.DateTime `json:"createdAt"`
	Deleted              *bool            `json:"deleted"`
	ID                   *IDRole          `json:"id"`
	ModifiedAt           *strfmt.DateTime `json:"modifiedAt"`
	NotDeleted           *bool            `json:"notDeleted"`
	PrincipalID          *IDPrincipal     `json:"principalId"`
	RawExternalID        string           `json:"rawExternalId,omitempty"`
	ResourceID           *IDResource      `json:"resourceId"`
	Roles                []*IDRoleMeta    `json:"roles"`
	State                *State           `json:"state"`
}

type RoleMeta struct {
	CheckedID            int64       `json:"checkedId,omitempty"`
	CheckedRawExternalID string      `json:"checkedRawExternalId,omitempty"`
	Content              *JSONNode   `json:"content,omitempty"`
	ID                   *IDRoleMeta `json:"id"`
	RawExternalID        string      `json:"rawExternalId,omitempty"`
}

type Row struct {
	ParticipantInfo *Participant     `json:"participantInfo,omitempty"`
	PlaceFrom       int32            `json:"placeFrom,omitempty"`
	PlaceTo         int32            `json:"placeTo,omitempty"`
	ProblemResults  []*ProblemResult `json:"problemResults"`
	Score           string           `json:"score,omitempty"`
}

type RowTitle struct {
	Name  string `json:"name,omitempty"`
	Title string `json:"title,omitempty"`
}

type RunReport struct {
	RunID int64 `json:"runId,omitempty"`
}

type SimplifiedFullRunReport struct {
	CompileLog           *string                          `json:"compileLog"`
	Compiler             *string                          `json:"compiler"`
	ContestID            *int64                           `json:"contestId"`
	ContestName          *string                          `json:"contestName"`
	FinalScore           *string                          `json:"finalScore"`
	IP                   string                           `json:"ip,omitempty"`
	MaxMemoryUsage       *int64                           `json:"maxMemoryUsage"`
	MaxTimeUsage         *int64                           `json:"maxTimeUsage"`
	Participant          *Participant                     `json:"participant"`
	PostprocessorMessage string                           `json:"postprocessorMessage,omitempty"`
	PrecompileChecks     []*AdditionalSolutionCheckResult `json:"precompileChecks"`
	PreliminaryScore     *string                          `json:"preliminaryScore"`
	ProblemAlias         *string                          `json:"problemAlias"`
	ProblemID            *string                          `json:"problemId"`
	RunID                *int64                           `json:"runId"`
	Status               string                           `json:"status,omitempty"`
	SubmissionTime       *string                          `json:"submissionTime"`
	TestFileType         *string                          `json:"testFileType"`
	Tests                []*TestLog                       `json:"tests"`
	TimeFromStart        *int64                           `json:"timeFromStart"`
	Verdict              *string                          `json:"verdict"`
}

type State struct {
	Name *string `json:"name"`
}

type Team struct {
	ID   *int64 `json:"id"`
	Name string `json:"name,omitempty"`
}

type User struct {
	AccessLevel          *AccessLevel `json:"accessLevel,omitempty"`
	CreationTime         []int64      `json:"creationTime"`
	Creator              *User        `json:"creator,omitempty"`
	CreatorID            []int64      `json:"creatorId"`
	DisplayedName        string       `json:"displayedName,omitempty"`
	Guest                bool         `json:"guest,omitempty"`
	ID                   *int64       `json:"id"`
	Login                string       `json:"login,omitempty"`
	PasswordHash         string       `json:"passwordHash,omitempty"`
	Salt                 string       `json:"salt,omitempty"`
	StudentsGroup        *UserGroup   `json:"studentsGroup,omitempty"`
	StudentsInvitesGroup *UserGroup   `json:"studentsInvitesGroup,omitempty"`
	UID                  string       `json:"uid,omitempty"`
	UIDAsLong            int64        `json:"uidAsLong,omitempty"`
	UserType             string       `json:"userType,omitempty"`
}

type UserGroup struct {
	Global     bool              `json:"global,omitempty"`
	ID         *int64            `json:"id"`
	Name       string            `json:"name,omitempty"`
	Owner      *User             `json:"owner,omitempty"`
	System     bool              `json:"system,omitempty"`
	UserAccess map[string]string `json:"userAccess,omitempty"`
	Users      []*User           `json:"users"`
}

type Group struct {
	ID          *int64 `json:"id,omitempty"`
	MemberCount int64  `json:"memberCount,omitempty"`
	Name        string `json:"name,omitempty"`
}
