package arith

import (
	"math"
	"strconv"
)

type Float64 float64

func (f Float64) Neg() Number {
	return -f
}
func (f Float64) Abs() Number {
	return Float64(math.Abs(float64(f)))
}
func (f Float64) Add(n Number) Number {
	switch n.(type) {
	case OtNumber:
		return f.Add(n.(OtNumber).n)
	case Float64:
		return f + n.(Float64)
	case expFloat64:
		return n.Add(f)
	}
	panic("invalid type")
}

func (f Float64) Sub(n Number) Number {
	return f.Add(n.Neg())
}

func (f Float64) Mul(n Number) Number {
	switch n.(type) {
	case OtNumber:
		return f.Mul(n.(OtNumber).n)
	case Float64:
		n2 := n.(Float64)
		r := float64(f * n2)
		return fitFloat64(r, 0)
	case expFloat64:
		return n.Mul(f)
	}
	panic("invalid type")
}

func (f Float64) Div(n Number) Number {
	switch n.(type) {
	case OtNumber:
		return f.Div(n.(OtNumber).n)
	case Float64:
		n2 := n.(Float64)
		r := float64(f / n2)
		return fitFloat64(r, 0)
	case expFloat64:
		n2 := n.(expFloat64)
		r := float64(f) / n2.val
		return fitFloat64(r, -n2.exp)
	}
	panic("invalid type")
}

func (f Float64) String() string {
	return strconv.FormatFloat(float64(f), 'f', -1, 64)
}

func (f Float64) FormatFloat(precison int32) string {
	return strconv.FormatFloat(float64(f), 'f', int(precison), 64)
}

func (f Float64) Float64() float64 {
	return float64(f)
}

func (f Float64) Cmp(n Number) int {
	switch n.(type) {
	case OtNumber:
		return f.Cmp(n.(OtNumber).n)
	case Float64:
		if f < n.(Float64) {
			return -1
		}
		if f > n.(Float64) {
			return 1
		}
		return 0
	case expFloat64:
		return -n.Cmp(f)
	}
	panic("invalid type")
}

func (f Float64) Max(numbers ...Number) Number {
	var maxNumber Number = f
	for _, n := range numbers {
		if maxNumber.Cmp(n) > 0 {
			maxNumber = n
		}
	}
	return maxNumber
}

func (f Float64) Min(numbers ...Number) Number {
	var minNumber Number = f
	for _, n := range numbers {
		if minNumber.Cmp(n) < 0 {
			minNumber = n
		}
	}
	return minNumber
}

func (f Float64) Pow(number Number) Number {
	return Float64(math.Pow(f.Float64(), number.Float64()))
}

func (f Float64) Atan() Number {
	return Float64(math.Atan(float64(f)))
}
