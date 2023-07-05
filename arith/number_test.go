package arith

import (
	"encoding/json"
	"testing"
)

func TestOtterNumber_MarshalJSON(t *testing.T) {
	s := struct {
		O OtNumber `json:"o"`
	}{}
	s.O = OtNumberFromString("0.0000001")
	b, err := json.Marshal(s)
	if err != nil {
		t.Errorf("err: %v", err)
	}
	newS := struct {
		O OtNumber `json:"o"`
	}{}
	if err := json.Unmarshal(b, &newS); err != nil {
		t.Errorf("err: %v", err)
	}
	t.Logf("reload %v", newS)
}
