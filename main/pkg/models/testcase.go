package models

import "encoding/json"

type Testcase struct {
	Action    string          `json:"action"`
	Arguments json.RawMessage `json:"arguments"`
}

type TestcaseFile struct {
	Testcases map[string]Testcase `json:"testcases"`
}
