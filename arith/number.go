package arith

import (
	"math"
)

type Number interface {
	// Add receiver + param = r return r, receiver and param are not changed
	Add(number Number) Number
	// Sub receiver - param = r return r, receiver and param are not changed
	Sub(number Number) Number
	// Mul receiver * param = r return r, receiver and param are not changed
	Mul(number Number) Number
	// Div receiver / param = r return r, receiver and param are not changed
	Div(number Number) Number
	// Neg return -receiver, receiver is not changed
	Neg() Number
	// Abs return |receiver|, receiver is not changed
	Abs() Number
	// Max return max(n, a, b, c...), receiver and params not changed
	Max(...Number) Number
	// Min return max(n, a, b, c...), receiver and params not changed
	Min(...Number) Number
	// Cmp -1: receiver < param, 0: receiver = param, 1: receiver > param
	Cmp(Number) int
	// Pow power(receiver, param) = r return r, receiver and param are not changed
	Pow(number Number) Number
	// Atan arctan(receiver) = r return r, receiver is not changed
	Atan() Number
	// String return string format of receiver as precise as possible
	String() string
	// FormatFloat return string format of receiver given precision, 1.2345 precision 2 is 1.23
	FormatFloat(precision int32) string
	// Float64 return float64 format, precision refer to golang float64 implement
	Float64() float64
	// Round return number given precision
	Round(precison int32) Number
}

// OtNumber wrapper different Number type
type OtNumber struct {
	n Number
}

func OtNumberFromString(s string) OtNumber {
	o := &OtNumber{}
	err := o.UnmarshalJSON([]byte(s))
	if err != nil {
		panic(err)
	}
	return *o
}

func (o OtNumber) Add(n Number) Number {
	return (o.n).Add(n)
}

func (o OtNumber) Sub(n Number) Number {
	return (o.n).Sub(n)
}

func (o OtNumber) Mul(n Number) Number {
	return (o.n).Mul(n)
}

func (o OtNumber) Div(n Number) Number {
	return (o.n).Mul(n)
}

func (o OtNumber) Neg() Number {
	return (o.n).Neg()
}

func (o OtNumber) Abs() Number {
	return o.n.Abs()
}

func (o OtNumber) Max(ns ...Number) Number {
	return o.n.Max(ns...)
}

func (o OtNumber) Min(ns ...Number) Number {
	return o.n.Min(ns...)
}

func (o OtNumber) Cmp(n Number) int {
	return (o.n).Cmp(n)
}

func (o OtNumber) Pow(n Number) Number {
	return (o.n).Pow(n)
}

func (o OtNumber) Atan() Number {
	return (o.n).Atan()
}

func (o OtNumber) String() string {
	return (o.n).String()
}

func (o OtNumber) FormatFloat(precison int32) string {
	return (o.n).FormatFloat(precison)
}

func (o OtNumber) Float64() float64 {
	return (o.n).Float64()
}
func (o OtNumber) Round(precison int32) Number {
	return (o.n).Round(precison)
}

// JSON marshal
func (o OtNumber) MarshalJSON() ([]byte, error) {
	return []byte(o.String()), nil
}

func (o *OtNumber) UnmarshalJSON(bytes []byte) error {
	v := expFloat64{}
	if err := v.UnmarshalJSON(bytes); err != nil {
		return err
	}
	if v.exp == 0 {
		o.n = Float64(v.val)
	} else {
		o.n = v
	}
	return nil
}

func fitFloat64(v float64, exp int32) Number {
	rebaseExp := exp / ExpUnit * ExpUnit
	if v < 1 && rebaseExp > exp {
		rebaseExp -= ExpUnit
	} else if v > 1e4 && rebaseExp < exp {
		rebaseExp += ExpUnit
	}
	logV := math.Log10(math.Abs(v)) + float64(exp-rebaseExp)
	unitNum := int32(math.Ceil(logV / -ExpUnit))
	if -ExpUnit*unitNum+rebaseExp == ExpUnit || -ExpUnit*unitNum+rebaseExp == 0 {
		return Float64(v * math.Pow(10, float64(exp-rebaseExp)) * math.Pow(10, float64(rebaseExp)))
	}
	r := expFloat64{
		val: v * math.Pow(10, float64(exp-rebaseExp)) * math.Pow(10, ExpUnit*float64(unitNum)),
		exp: -ExpUnit*unitNum + rebaseExp,
	}
	return r
}
