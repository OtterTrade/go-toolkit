package arith

import (
	"errors"
	"math"

	"go.mongodb.org/mongo-driver/bson/bsontype"
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

// OtNumber wrapper different Number type, valid digit number is 16 because use float64 type
type OtNumber struct {
	Number
}

func OtNumberFromString(s string) OtNumber {
	o := &OtNumber{}
	err := o.UnmarshalJSON([]byte(s))
	if err != nil {
		panic(err)
	}
	return *o
}

// MarshalJSON JSON marshal
func (o OtNumber) MarshalJSON() ([]byte, error) {
	return []byte(o.String()), nil
}

func (o *OtNumber) UnmarshalJSON(bytes []byte) error {
	v := expFloat64{}
	if err := v.UnmarshalJSON(bytes); err != nil {
		return err
	}
	if v.exp == 0 {
		o.Number = Float64(v.val)
	} else {
		o.Number = v
	}
	return nil
}

// MarshalBSONValue implement OtNumber mongodb marshal
func (o *OtNumber) MarshalBSONValue() (bsontype.Type, []byte, error) {
	d, err := o.MarshalJSON()
	return bsontype.String, d, err
}

// UnmarshalBSONValue implement OtNumber mongodb unmarshal
func (o *OtNumber) UnmarshalBSONValue(ty bsontype.Type, data []byte) error {
	if ty == bsontype.String {
		return o.UnmarshalJSON(data)
	}
	return errors.New("OtNumber Unmarshal must be type string")
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
