package arith

import (
	"math"
)

type Number interface {
	Add(number Number) Number
	Sub(number Number) Number
	Mul(number Number) Number
	Div(number Number) Number
	Neg() Number
	Abs() Number
	Max(...Number) Number
	Min(...Number) Number

	// Cmp -1 this < param, 0 this = param, 1 this > param
	Cmp(Number) int
	Pow(number Number) Number
	Atan() Number

	String() string
	FormatFloat(precison int32) string
	Float64() float64
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

func fitFloat64(r float64, exp int32) Number {
	unitNum := int32(math.Log10(r) / math.Log10(1e-4))

	if unitNum >= 1 {
		if -ExpUnit*unitNum+exp == 0 {
			return Float64(r * math.Pow(10, ExpUnit*float64(unitNum)))
		}
		return expFloat64{
			val: r * math.Pow(10, ExpUnit*float64(unitNum)),
			exp: -ExpUnit*unitNum + exp,
		}
	} else if unitNum < -2 {
		if ExpUnit*unitNum+exp == 0 {
			return Float64(r / math.Pow(10, ExpUnit*float64(unitNum)))
		}
		return expFloat64{
			val: r / math.Pow(10, ExpUnit*float64(unitNum)),
			exp: ExpUnit*unitNum + exp,
		}
	}
	return Float64(r)
}
