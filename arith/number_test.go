package arith

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestOtterNumber_JSON(t *testing.T) {
	type args struct {
		strNum string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "small",
			args: args{strNum: "0.00000000000000000000000012231214241"},
		},
		{
			name: "mid",
			args: args{strNum: "1.15"},
		},
		{
			name: "large",
			args: args{strNum: "1223121424135252135241.616546878431321321321356"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := struct {
				O OtNumber `json:"o"`
			}{}
			s.O = OtNumberFromString(tt.args.strNum)
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
			nB, err := json.Marshal(newS)
			if !reflect.DeepEqual(nB, b) {
				t.Errorf("json marshal %s unmarshal %s", string(nB), string(b))
			}
		})
	}
}

func TestFitFloat64(t *testing.T) {
	type args struct {
		r   float64
		exp int32
	}
	tests := []struct {
		name string
		args args
		want Number
	}{
		{
			name: "small",
			args: args{
				r:   0.000000000001,
				exp: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fitFloat64(tt.args.r, tt.args.exp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fitFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}
