package models

type OnlineCompiler struct {
	Output     string `json:"output"`
	StatusCode int    `json:"statusCode"`
	Memory     string `json:"memory"`
	CpuTime    string `json:"cpuTime"`
}
