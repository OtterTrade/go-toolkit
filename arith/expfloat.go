package arith

import (
	"github.com/shopspring/decimal"
	"math"
)

const ExpUnit = 4

// expFloat64 number = val * 10 ^ exp, inner implement, not used by user
type expFloat64 struct {
	val float64
	exp int32
}

func (f expFloat64) Cmp(n Number) int {
	switch n.(type) {
	case OtNumber:
		return f.Cmp(n.(OtNumber).n)
	case Float64:
		n2 := n.Float64() * math.Pow(10, -float64(f.exp))
		if f.val < n2 {
			return -1
		}
		if f.val > n2 {
			return 1
		}
		return 0
	case expFloat64:
		n2 := f.Sub(n).(expFloat64).val
		if n2 < 0 {
			return -1
		}
		if n2 > 0 {
			return 1
		}
		return 0
	}
	panic("invalid type")
}

func (f expFloat64) FormatFloat(precison int32) string {
	return decimal.NewFromFloat(f.val).Shift(f.exp).StringFixed(precison)
}

func (f expFloat64) String() string {
	return decimal.NewFromFloat(f.val).Shift(f.exp).String()
}

func (f expFloat64) Float64() float64 {
	return math.Pow(10, float64(f.exp)) * f.val
}

func (f expFloat64) MarshalJSON() ([]byte, error) {
	s := decimal.NewFromFloat(f.val).Shift(f.exp).String()
	return []byte(s), nil
}

func (f *expFloat64) UnmarshalJSON(bytes []byte) error {
	s := string(bytes)
	v, err := decimal.NewFromString(s)
	if err != nil {
		return err
	}
	if v.Exponent() > -ExpUnit {
		f.exp = 0
		f.val = v.InexactFloat64()
		return nil
	}
	f.exp = (v.Exponent() / ExpUnit) * ExpUnit
	f.val = v.Shift(-f.exp).InexactFloat64()
	return nil
}

func (f expFloat64) Add(n Number) Number {
	switch n.(type) {
	case OtNumber:
		return f.Add(n.(OtNumber).n)
	case Float64:
		return f.Add(expFloat64{exp: 0, val: float64(n.(Float64))})
	case expFloat64:
		n2 := n.(expFloat64)
		alignExp := f.exp
		if n2.exp < f.exp {
			alignExp = n2.exp
		}
		v1 := f.val
		if alignExp < f.exp {
			v1 *= math.Pow(10, float64(-alignExp+f.exp))
		}
		v2 := n2.val
		if alignExp < n2.exp {
			v2 *= math.Pow(10, float64(-alignExp+n2.exp))
		}
		z := expFloat64{
			exp: alignExp,
			val: v1 + v2,
		}
		return z
	}
	panic("invalid type")
}

func (f expFloat64) Sub(n Number) Number {
	return f.Add(n.Neg())
}

func (f expFloat64) Neg() Number {
	return expFloat64{val: -f.val, exp: f.exp}
}

func (f expFloat64) Abs() Number {
	return expFloat64{val: math.Abs(f.val), exp: f.exp}
}

func (f expFloat64) Mul(n Number) Number {
	switch n.(type) {
	case OtNumber:
		return f.Mul(n.(OtNumber).n)
	case Float64:
		n2 := n.(Float64)
		r := float64(n2) * f.val
		return FitFloat64(r, f.exp)
	case expFloat64:
		n2 := n.(expFloat64)
		r := n2.val * f.val
		return FitFloat64(r, f.exp+n2.exp)
	}
	panic("invalid type")
}

func (f expFloat64) Div(n Number) Number {
	switch n.(type) {
	case OtNumber:
		return f.Div(n.(OtNumber).n)
	case Float64:
		n2 := n.(Float64)
		r := f.val / float64(n2)
		return FitFloat64(r, f.exp)
	case expFloat64:
		n2 := n.(expFloat64)
		r := f.val / n2.val
		return FitFloat64(r, f.exp-n2.exp)
	}
	panic("invalid type")
}

func (f expFloat64) Max(numbers ...Number) Number {
	var maxNumber Number = f
	for _, n := range numbers {
		if maxNumber.Cmp(n) > 0 {
			maxNumber = n
		}
	}
	return maxNumber
}

func (f expFloat64) Min(numbers ...Number) Number {
	var minNumber Number = f
	for _, n := range numbers {
		if minNumber.Cmp(n) < 0 {
			minNumber = n
		}
	}
	return minNumber
}

func (f expFloat64) Pow(number Number) Number {
	exponential := number.Float64()
	floatExp := exponential * float64(f.exp)
	f.exp = int32(floatExp) / ExpUnit * ExpUnit
	f.val = math.Pow(f.val, exponential) * math.Pow(10, floatExp-float64(f.exp))
	if f.exp == 0 {
		return Float64(f.val)
	}
	return f
}

func (f expFloat64) Atan() Number {
	return Float64(math.Atan(f.Float64()))
}
