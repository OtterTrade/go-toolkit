package arith

import (
	"math"

	"github.com/shopspring/decimal"
)

const ExpUnit = 4

var SmallBoarder = 0.0001
var LargeBoarder = 1e15

// expFloat64 number = val * 10 ^ exp, inner implement, not used by user, precision is also limited by float64
type expFloat64 struct {
	val float64
	exp int32
}

func (f expFloat64) Cmp(n Number) int {
	switch n.(type) {
	case decimalNumber:
		return -n.Cmp(f)
	case OtNumber:
		return f.Cmp(n.(OtNumber).Number)
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
	rawV := v.InexactFloat64()
	if rawV != 0 && (math.Abs(rawV) < SmallBoarder || math.Abs(rawV) > LargeBoarder) {
		digitNum := math.Log10(math.Abs(rawV))

		f.exp = int32(digitNum) / ExpUnit * ExpUnit
		if digitNum > 2*ExpUnit {
			f.exp -= 2 * ExpUnit
		}
		f.val = v.Shift(-f.exp).InexactFloat64()
		return nil
	}
	f.exp = 0
	f.val = rawV
	return nil
}

func (f expFloat64) Add(n Number) Number {
	switch n.(type) {
	case decimalNumber:
		return n.Add(f)
	case OtNumber:
		return f.Add(n.(OtNumber).Number)
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
		r := expFloat64{
			exp: alignExp,
			val: v1 + v2,
		}
		return r
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
	case decimalNumber:
		return n.Mul(f)
	case OtNumber:
		return f.Mul(n.(OtNumber).Number)
	case Float64:
		n2 := n.(Float64)
		r := float64(n2) * f.val
		return fitFloat64(r, f.exp)
	case expFloat64:
		n2 := n.(expFloat64)
		r := n2.val * f.val
		return fitFloat64(r, f.exp+n2.exp)
	}
	panic("invalid type")
}

func (f expFloat64) Div(n Number) Number {
	switch n.(type) {
	case decimalNumber:
		d := decimalNumber{}
		d.Decimal = decimal.NewFromFloat(f.val)
		d.Decimal = d.Decimal.Shift(f.exp)
		return d.Div(n)
	case OtNumber:
		return f.Div(n.(OtNumber).Number)
	case Float64:
		n2 := n.(Float64)
		r := f.val / float64(n2)
		return fitFloat64(r, f.exp)
	case expFloat64:
		n2 := n.(expFloat64)
		r := f.val / n2.val
		return fitFloat64(r, f.exp-n2.exp)
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

func (f expFloat64) Round(precision int32) Number {
	exp := f.exp + precision
	v := math.Round(math.Pow(10, float64(exp)) * f.val)
	// val * 10 ** exp -> val * 10 ** (exp + precision) * 10 ** (-precision)
	return fitFloat64(v, -precision)
}
