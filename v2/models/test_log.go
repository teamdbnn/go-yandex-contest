package models

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
