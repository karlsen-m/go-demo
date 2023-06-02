package sts

import "testing"

func TestGetToken(t *testing.T) {
	s, err := NewSts("", "")
	if err != nil {
		panic(err)
	}
	credential := s.GetFederationCredentials(R{
		Name: "ocr",
		Policy: Policy{
			Version: "2.0",
			Statement: []Statement{
				{
					Action:   []string{"ocr:*"},
					Resource: []string{"*"},
					Effect:   "allow",
				},
			},
		},
		DurationSeconds: 30,
	})
	t.Log(credential)
	return
}
