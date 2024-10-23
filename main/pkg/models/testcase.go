package models

type Testcase struct {
	Key    string
	Action string         `json:"action"`
	Values ActionExecutor `json:"arguments"`
}

type TestcaseFile struct {
	Testcases map[string]Testcase `json:"testcases"`
}

type Response struct {
	Response map[string]interface{} `json:"responses"`
}
