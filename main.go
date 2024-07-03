package num

import (
	"fmt"
	"math"
	"strconv"
)

type Integer interface {
	Signed | Unsigned
}
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func ToInt[T Integer](v any) T {
	n, err := ParseInt[T](v)
	Panic(err)
	return n
}

func ToFloat64(a any) float64 {
	f, err := ParseFloat64(a)
	Panic(err)
	return f
}

func ParseInt[T Integer](v any) (n T, err error) {
	if v == nil {
		err = fmt.Errorf("object is nil")
		return
	}
	var i int64
	switch v.(type) {
	case int:
		return T(v.(int)), nil
	case float64:
		return T(v.(float64)), nil
	case string:
		i, err = strconv.ParseInt(v.(string), 0, 64)
	case nil:
		err = fmt.Errorf("object is nil")
	default:
		i, err = strconv.ParseInt(fmt.Sprint(v), 0, 64)
	}
	return T(i), nil
}

func ParseFloat64(a any) (f float64, err error) {
	switch a := a.(type) {
	case int:
		f = float64(a)
	case int8:
		f = float64(a)
	case int16:
		f = float64(a)
	case int32:
		f = float64(a)
	case int64:
		f = float64(a)
	case uint:
		f = float64(a)
	case uint8:
		f = float64(a)
	case uint16:
		f = float64(a)
	case uint32:
		f = float64(a)
	case uint64:
		f = float64(a)
	case float32:
		f = float64(a)
	case float64:
		f = a
	case string:
		f, err = strconv.ParseFloat(a, 64)
	default:
		err = fmt.Errorf(fmt.Sprint(a) + " is not a number")
	}
	return
}

func FloatToString(v float64, step float64) string {
	p := int(-math.Log10(step) + 0.5) //sometimes Log10 returns .999
	return strconv.FormatFloat(FloatToFixed(v, step), 'f', p, 64)
}

func FloatToFixed(v float64, step float64) float64 {
	if step == 0 {
		return v
	}
	return step * math.Floor(v/step)
}

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
