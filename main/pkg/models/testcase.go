package models

type Testcase struct {
	Key       string
	Execute   func(interface{}) error
	Action    string      `json:"action"`
	Arguments interface{} `json:"arguments"`
}

/*func (tc *Testcase) Execute() error {
	return tc.Execute(tc.Arguments)
}
*/

type TestcaseFile struct {
	Testcases map[string]Testcase `json:"testcases"`
}
