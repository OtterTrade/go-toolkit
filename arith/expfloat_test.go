package arith

import (
	"github.com/shopspring/decimal"
	"reflect"
	"testing"
)

func Test_expFloat64_Cmp(t *testing.T) {
	type fields struct {
		s string
	}
	type args struct {
		n Number
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "small",
			fields: fields{
				s: "0.00000000000000000000012345",
			},
			args: args{
				n: OtNumberFromString("0.00000000000000000000987654"),
			},
			want: -1,
		},
		{
			name: "big",
			fields: fields{
				s: "0.00000000000000000000012345",
			},
			args: args{
				n: OtNumberFromString("-0.00000000000000000000987654"),
			},
			want: 1,
		},
		{
			name: "equal",
			fields: fields{
				s: "0.0000000000000000000000000000000000000000000012345",
			},
			args: args{
				n: OtNumberFromString("0.0000000000000000000000000000000000000000000012345"),
			},
			want: 0,
		},
		{
			name: "small_float",
			fields: fields{
				s: "0.21342",
			},
			args: args{
				n: OtNumberFromString("12134124312524614143241"),
			},
			want: -1,
		},
		{
			name: "equal_float",
			fields: fields{
				s: "0.21342",
			},
			args: args{
				n: OtNumberFromString("0.21342"),
			},
			want: 0,
		},
		{
			name: "big_float",
			fields: fields{
				s: "0.21342",
			},
			args: args{
				n: OtNumberFromString("-67860.21342"),
			},
			want: 1,
		},
		{
			name: "big_float2",
			fields: fields{
				s: "0.21342",
			},
			args: args{
				n: OtNumberFromString("-6.21342"),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &expFloat64{}
			err := f.UnmarshalJSON([]byte(tt.fields.s))
			if err != nil {
				t.Errorf("unmarshal err: %v", err)
			}
			if got := f.Cmp(tt.args.n); got != tt.want {
				t.Errorf("Cmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expFloat64_Add(t *testing.T) {
	type fields struct {
		s string
	}
	type args struct {
		s string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "normal",
			fields: fields{
				s: "0.002314",
			},
			args: args{
				s: "-6.21342",
			},
		},
		{
			name: "both_small",
			fields: fields{
				s: "0.0000000000000002",
			},
			args: args{
				s: "-0.000000000000000000212",
			},
		},
		{
			name: "boarder_precision",
			fields: fields{
				s: "-0.000000000000001",
			},
			args: args{
				s: "1.0000021342",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a1, _ := decimal.NewFromString(tt.fields.s)
			a2, _ := decimal.NewFromString(tt.args.s)
			tt.want = a1.Add(a2).String()
			f := &expFloat64{}
			err := f.UnmarshalJSON([]byte(tt.fields.s))
			if err != nil {
				t.Errorf("unmarshal err: %v", err)
			}
			if got := f.Add(OtNumberFromString(tt.args.s)).String(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
