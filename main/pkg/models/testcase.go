package models

import "encoding/json"

type Testcase struct {
	//Key       string
	Action    string          `json:"action"`
	Arguments json.RawMessage `json:"arguments"`
	//Values    ActionExecutor
}

type TestcaseFile struct {
	Testcases map[string]Testcase `json:"testcases"`
}

type Response struct {
	Response map[string]interface{} `json:"responses"`
}
