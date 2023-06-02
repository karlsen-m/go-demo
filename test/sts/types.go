package sts

type R struct {
	Name            string
	Policy          Policy
	DurationSeconds uint64
}
type Policy struct {
	Version   string      `json:"version"`
	Statement []Statement `json:"statement"`
}
type Statement struct {
	Action   []string `json:"action"`
	Resource []string `json:"resource"`
	Effect   string   `json:"effect"`
}
