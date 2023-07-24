package arith

import (
	"fmt"
	"math"

	"github.com/shopspring/decimal"
)

func init() {
	decimal.DivisionPrecision = 100
}

type decimalNumber struct {
	decimal.Decimal
}

func decimalFromString(s string) decimalNumber {
	v, err := decimal.NewFromString(s)
	if err != nil {
		panic(err)
	}
	r := decimalNumber{}
	r.Decimal = v
	return r
}

func (d decimalNumber) Mul(number Number) Number {
	var param decimal.Decimal
	switch number.(type) {
	case Float64:
		param = decimal.NewFromFloat(number.Float64())
	case expFloat64:
		var err error
		param, err = decimal.NewFromString(number.String())
		if err != nil {
			panic(fmt.Sprintf("Decimal parse %s failed", number.String()))
		}
	case decimalNumber:
		param = number.(decimalNumber).Decimal
	}
	r := decimalNumber{}
	r.Decimal = d.Decimal.Mul(param)
	return r
}

func (d decimalNumber) Div(number Number) Number {
	var param decimal.Decimal
	switch number.(type) {
	case Float64:
		param = decimal.NewFromFloat(number.Float64())
	case expFloat64:
		var err error
		param, err = decimal.NewFromString(number.String())
		if err != nil {
			panic(fmt.Sprintf("Decimal parse %s failed", number.String()))
		}
	case decimalNumber:
		param = number.(decimalNumber).Decimal
	}
	r := decimalNumber{}
	r.Decimal = d.Decimal.Div(param)
	return r
}

func (d decimalNumber) Neg() Number {
	r := decimalNumber{}
	r.Decimal = d.Decimal.Neg()
	return r
}

func (d decimalNumber) Abs() Number {
	r := decimalNumber{}
	r.Decimal = d.Decimal.Abs()
	return r
}

func (d decimalNumber) Max(number ...Number) Number {
	var r Number = d
	for _, n := range number {
		if r.Cmp(n) < 0 {
			r = n
		}
	}
	return r
}

func (d decimalNumber) Min(number ...Number) Number {
	var r Number = d
	for _, n := range number {
		if r.Cmp(n) > 0 {
			r = n
		}
	}
	return r
}

func (d decimalNumber) Cmp(number Number) int {
	v := d.Sub(number)
	switch v.(type) {
	case Float64:
		f := v.Float64()
		if f < 0 {
			return -1
		}
		if f > 0 {
			return 1
		}
		return 0
	case expFloat64:
		f := v.(expFloat64).val
		if f < 0 {
			return -1
		}
		if f > 0 {
			return 1
		}
		return 0
	case decimalNumber:
		return v.(decimalNumber).Decimal.Sign()
	}
	panic("invalid type")
}

func (d decimalNumber) Pow(number Number) Number {
	expo := number.Float64()
	expoInt := int32(number.Float64())
	leftExpo := expo - float64(expoInt)
	r := decimalNumber{}
	r.Decimal = d.Decimal.Pow(decimal.NewFromInt32(expoInt))
	return r.Mul(Float64(math.Pow(d.Float64(), leftExpo)))
}

func (d decimalNumber) Atan() Number {
	r := decimalNumber{}
	r.Decimal = d.Decimal.Atan()
	return r
}

func (d decimalNumber) FormatFloat(precision int32) string {
	return d.Decimal.Round(precision).String()
}

func (d decimalNumber) Float64() float64 {
	f, _ := d.Decimal.Float64()
	return f
}

func (d decimalNumber) Round(precison int32) Number {
	r := decimalNumber{}
	r.Decimal = d.Decimal.Round(precison)
	return r
}

func (d decimalNumber) Add(number Number) Number {
	var param decimal.Decimal
	switch number.(type) {
	case Float64:
		param = decimal.NewFromFloat(number.Float64())
	case expFloat64:
		var err error
		param, err = decimal.NewFromString(number.String())
		if err != nil {
			panic(fmt.Sprintf("Decimal parse %s failed", number.String()))
		}
	case decimalNumber:
		param = number.(decimalNumber).Decimal
	}
	r := decimalNumber{}
	r.Decimal = d.Decimal.Add(param)
	return r
}

func (d decimalNumber) Sub(n Number) Number {
	return d.Add(n.Neg())
}
