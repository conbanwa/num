package num

import (
	"fmt"
	"math"
	"strconv"
	"sync"

	"golang.org/x/exp/constraints"
)

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
		err = fmt.Errorf("not a number")
	}
	return
}

func ToFloat64(a any) float64 {
	f, err := ParseFloat64(a)
	if err != nil {
		panic(err)
		return 0
	}
	return f
}

func ParseInteger[T constraints.Integer](v any) T {
	if v == nil {
		return 0
	}
	var i int64
	switch v.(type) {
	case int:
		return T(v.(int))
	case float64:
		return T(v.(float64))
	case string:
		i, _ = strconv.ParseInt(v.(string), 0, 64)
	case nil:
		return 0
	default:
		i, _ = strconv.ParseInt(fmt.Sprint(v), 0, 64)
	}
	return T(i)
}

func FloatToString(v float64, step float64) string {
	return strconv.FormatFloat(FloatToFixed(v, step), 'f', int(-math.Log10(step)+0.5), 64)
}

func FloatToFixed(v float64, step float64) float64 {
	if step == 0 {
		return v
	}
	return step * math.Floor(v/step)
}

func SyncMapLen(m *sync.Map) (l int) {
	m.Range(func(key, value any) bool {
		l++
		return true
	})
	return
}

func SyncMapValueDefault0[T comparable](sm *sync.Map, k string) (res T) {
	return SyncMapValue[T](sm, k, res)
}

func SyncMapValue[T any](sm *sync.Map, k string, defaults T) (res T) {
	value, ok := sm.Load(k)
	if !ok {
		return defaults
	}
	res, ok = value.(T)
	if !ok {
		return defaults
	}
	return
}
