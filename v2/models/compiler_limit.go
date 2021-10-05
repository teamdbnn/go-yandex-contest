package models

type CompilerLimit struct {
	CompilerName  string `json:"compilerName,omitempty"`
	IdlenessLimit int64  `json:"idlenessLimit,omitempty"`
	MemoryLimit   int64  `json:"memoryLimit,omitempty"`
	OutputLimit   int64  `json:"outputLimit,omitempty"`
	TimeLimit     int64  `json:"timeLimit,omitempty"`
}
