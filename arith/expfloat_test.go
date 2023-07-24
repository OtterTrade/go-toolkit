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

func Test_expFloat64_FormatFloat(t *testing.T) {
	type fields struct {
		s string
	}
	type args struct {
		precison int32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
		{
			name: "",
			fields: fields{
				s: "0.00000000000001",
			},
			args: args{
				precison: 2,
			},
			want: "0.00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &expFloat64{}
			err := f.UnmarshalJSON([]byte(tt.fields.s))
			if err != nil {
				t.Errorf("unmarshal err: %v", err)
			}
			if got := f.FormatFloat(tt.args.precison); got != tt.want {
				t.Errorf("FormatFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expFloat64_Abs(t *testing.T) {
	type fields struct {
		val float64
		exp int32
	}
	tests := []struct {
		name   string
		fields fields
		want   Number
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := expFloat64{
				val: tt.fields.val,
				exp: tt.fields.exp,
			}
			if got := f.Abs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expFloat64_Atan(t *testing.T) {
	type fields struct {
		val float64
		exp int32
	}
	tests := []struct {
		name   string
		fields fields
		want   Number
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := expFloat64{
				val: tt.fields.val,
				exp: tt.fields.exp,
			}
			if got := f.Atan(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Atan() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_expFloat64_Div(t *testing.T) {
	type fields struct {
		val float64
		exp int32
	}
	type args struct {
		n Number
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Number
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := expFloat64{
				val: tt.fields.val,
				exp: tt.fields.exp,
			}
			if got := f.Div(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expFloat64_Float64(t *testing.T) {
	type fields struct {
		val float64
		exp int32
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := expFloat64{
				val: tt.fields.val,
				exp: tt.fields.exp,
			}
			if got := f.Float64(); got != tt.want {
				t.Errorf("Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expFloat64_MarshalJSON(t *testing.T) {
	type fields struct {
		val float64
		exp int32
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := expFloat64{
				val: tt.fields.val,
				exp: tt.fields.exp,
			}
			got, err := f.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expFloat64_Max(t *testing.T) {
	type fields struct {
		val float64
		exp int32
	}
	type args struct {
		numbers []Number
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Number
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := expFloat64{
				val: tt.fields.val,
				exp: tt.fields.exp,
			}
			if got := f.Max(tt.args.numbers...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expFloat64_Min(t *testing.T) {
	type fields struct {
		val float64
		exp int32
	}
	type args struct {
		numbers []Number
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Number
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := expFloat64{
				val: tt.fields.val,
				exp: tt.fields.exp,
			}
			if got := f.Min(tt.args.numbers...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expFloat64_Mul(t *testing.T) {
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
			name: "1",
			fields: fields{
				s: "43000.14",
			},
			args: args{
				s: "3",
			},
		},
		{
			name: "2",
			fields: fields{
				s: "0.0000000000000002",
			},
			args: args{
				s: "-0.000000000000000000212",
			},
		},
		{
			name: "3",
			fields: fields{
				s: "-0.000000000000001",
			},
			args: args{
				s: "1.21342",
			},
		},
		{
			name: "4",
			fields: fields{
				s: "124314148.1",
			},
			args: args{
				s: "36995.",
			},
		},
		{
			name: "5",
			fields: fields{
				s: "-12141.2",
			},
			args: args{
				s: "21.2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a1, _ := decimal.NewFromString(tt.fields.s)
			a2, _ := decimal.NewFromString(tt.args.s)
			tt.want = a1.Mul(a2).String()
			f := &expFloat64{}
			err := f.UnmarshalJSON([]byte(tt.fields.s))
			if err != nil {
				t.Errorf("unmarshal err: %v", err)
			}
			if got := f.Mul(OtNumberFromString(tt.args.s)).String(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expFloat64_Neg(t *testing.T) {
	type fields struct {
		val float64
		exp int32
	}
	tests := []struct {
		name   string
		fields fields
		want   Number
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := expFloat64{
				val: tt.fields.val,
				exp: tt.fields.exp,
			}
			if got := f.Neg(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Neg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expFloat64_Pow(t *testing.T) {
	type fields struct {
		val float64
		exp int32
	}
	type args struct {
		number Number
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Number
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := expFloat64{
				val: tt.fields.val,
				exp: tt.fields.exp,
			}
			if got := f.Pow(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expFloat64_Round(t *testing.T) {
	type fields struct {
		val float64
		exp int32
	}
	type args struct {
		precision int32
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Number
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := expFloat64{
				val: tt.fields.val,
				exp: tt.fields.exp,
			}
			if got := f.Round(tt.args.precision); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Round() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expFloat64_String(t *testing.T) {
	type fields struct {
		val float64
		exp int32
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := expFloat64{
				val: tt.fields.val,
				exp: tt.fields.exp,
			}
			if got := f.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expFloat64_Sub(t *testing.T) {
	type fields struct {
		val float64
		exp int32
	}
	type args struct {
		n Number
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Number
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := expFloat64{
				val: tt.fields.val,
				exp: tt.fields.exp,
			}
			if got := f.Sub(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_expFloat64_UnmarshalJSON(t *testing.T) {
	type fields struct {
		val float64
		exp int32
	}
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &expFloat64{
				val: tt.fields.val,
				exp: tt.fields.exp,
			}
			if err := f.UnmarshalJSON(tt.args.bytes); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
