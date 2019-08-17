package model

type Compile struct {
	Language string `json:"language"`
	Body     string `json:"body"`
}

type CompileResult struct {
	Language     string `json:"language"`
	ErrorMessage string `json:"error_message"`
	Output       string `json:"output"`
}
