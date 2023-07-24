package arith

import (
	"encoding/json"
	"math/rand"
	"reflect"
	"testing"

	"github.com/shopspring/decimal"
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

func Test_fitFloat64(t *testing.T) {
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
			name: "1",
			args: args{
				r:   0.00001,
				exp: -10,
			},
			want: expFloat64{
				val: 10,
				exp: -16,
			},
		},
		{
			name: "2",
			args: args{
				r:   1e8,
				exp: -10,
			},
			want: expFloat64{
				val: 100,
				exp: -4,
			},
		},
		{
			name: "3",
			args: args{
				r:   1e8,
				exp: 10,
			},
			want: expFloat64{
				val: 100,
				exp: 16,
			},
		},
		{
			name: "4",
			args: args{
				r:   1e21,
				exp: 3,
			},
			want: expFloat64{
				val: 1,
				exp: 24,
			},
		},
		{
			name: "5",
			args: args{
				r:   1e21 * 21.21241,
				exp: -21,
			},
			want: Float64(21.21241),
		},
		{
			name: "6",
			args: args{
				r:   123444.1231,
				exp: 0,
			},
			want: Float64(123444.1231),
		},
		{
			name: "7",
			args: args{
				r:   123444.1231,
				exp: 1,
			},
			want: Float64(1234441.231),
		},
		{
			name: "8",
			args: args{
				r:   123444.1231,
				exp: -1,
			},
			want: Float64(12344.41231),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fitFloat64(tt.args.r, tt.args.exp); !reflect.DeepEqual(got, tt.want) || !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("fitFloat64() = %v, want %v", got.String(), tt.want.String())
			}
		})
	}
}

func TestOtNumber_MarshalJSON(t *testing.T) {
	type fields struct {
		s string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "1",
			fields: fields{
				s: "0.0000000000000000241",
			},
			want:    []byte("0.0000000000000000241"),
			wantErr: false,
		},
		{
			name: "2",
			fields: fields{
				s: "-0.0000000000000000241",
			},
			want:    []byte("-0.0000000000000000241"),
			wantErr: false,
		},
		{
			name: "3",
			fields: fields{
				s: "242352.212",
			},
			want:    []byte("242352.212"),
			wantErr: false,
		},
		{
			name: "4",
			fields: fields{
				s: "-2352.2132",
			},
			want:    []byte("-2352.2132"),
			wantErr: false,
		},
		{
			name: "max_precision",
			fields: fields{
				s: "9876543210123456",
			},
			want:    []byte("9876543210123456"),
			wantErr: false,
		},
		{
			name: "16digit_precision",
			fields: fields{
				s: "98765432101234567",
			},
			want:    []byte("98765432101234570"),
			wantErr: false,
		},
		{
			name: "16digit_precision",
			fields: fields{
				s: "987654321012345678",
			},
			want:    []byte("987654321012345700"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := OtNumberFromString(tt.fields.s)
			got, err := o.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

// Test_RandomCalc use decimal (DivisionPrecision has set to 100) as ground truth, test precision
func Test_RandomCalc(t *testing.T) {
	generateRandom := func() string {
		return decimal.NewFromFloat(rand.NormFloat64()).Shift(int32(rand.Intn(64) - 32)).String()
	}
	diffs := make([]Number, 0)
	for i := 0; i < 100; i++ {
		v1 := generateRandom()
		v2 := generateRandom()
		ot1 := OtNumberFromString(v1)
		ot2 := OtNumberFromString(v2)
		d1 := decimalFromString(v1)
		d2 := decimalFromString(v2)
		r1 := ot1.Add(ot2)
		r2 := d1.Add(d2)
		diff := r1.Div(r2).Sub(Float64(1)).Abs()
		diffs = append(diffs, diff)
		if diff.Cmp(Float64(0.000001)) > 0 {
			t.Errorf("%s + %s = %s, %s diff: %sprecision out of range", v1, v2, r1, r2, diff)
		}
		r1 = ot1.Mul(ot2)
		r2 = d1.Mul(d2)
		diff = r1.Div(r2).Sub(Float64(1)).Abs()
		diffs = append(diffs, diff)
		if diff.Cmp(Float64(0.000001)) > 0 {
			t.Errorf("%s * %s = %s, %s diff: %sprecision out of range", v1, v2, r1, r2, diff)
		}
		if ot2.Cmp(Float64(0)) == 0 || r2.Cmp(Float64(0)) == 0 {
			continue
		}
		r1 = ot1.Div(ot2)
		r2 = d1.Div(d2)
		diff = r1.Div(r2).Sub(Float64(1)).Abs()
		diffs = append(diffs, diff)
		if diff.Cmp(Float64(0.000001)) > 0 {
			t.Errorf("%s / %s = %s, %s diff: %sprecision out of range", v1, v2, r1, r2, diff)
		}
	}
	maxDiff := diffs[0].Max(diffs...).Float64()
	t.Logf("max diff: %v", maxDiff)
}

// Test_Marshal_Unmarshal test struct marshal/unmarshal
func Test_Marshal_Unmarshal(t *testing.T) {
	type m struct {
		Num OtNumber `json:"num"`
	}

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
			args: args{strNum: "1234567.890123456"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &m{}
			v.Num = OtNumberFromString(tt.args.strNum)
			d, err := json.Marshal(v)
			if err != nil {
				t.Errorf("json marshal %v err %v", v, err)
			}
			r := &m{}
			err = json.Unmarshal(d, r)
			if err != nil {
				t.Errorf("json unmarshal %v err %v", string(d), err)
			}
			if !reflect.DeepEqual(v, r) {
				t.Errorf("json marshal %v unmarshal %v", v, r)
			}
		})
	}

}
