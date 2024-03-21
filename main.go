package num

import (
	"fmt"
	"math"
	"strconv"
	"sync"

	"golang.org/x/exp/constraints"
)

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
	m.Range(func(key, value interface{}) bool {
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
