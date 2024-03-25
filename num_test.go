package num

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/conbanwa/compare"
	"github.com/stretchr/testify/assert"
)

func FuzzParseFloat64(f *testing.F) {
	var (
		a int
		b uint
		c int8
		d int16
		e int32
		j float32 = float32(math.NaN())
		k float64 = math.NaN()
		g string  = "NaN"
	)
	f.Add(a, b, c, d, e, j, k, g)
	f.Fuzz(func(t *testing.T, a int, b uint, c int8, d int16, e int32, j float32, k float64, g string) {
		parseFloat(t, a)
		parseFloat(t, b)
		parseFloat(t, c)
		parseFloat(t, d)
		parseFloat(t, e)
		parseFloat(t, j)
		parseFloat(t, k)
		parseFloat(t, g)
	})
}
func TestParseFloat64(t *testing.T) {
	tests := []struct {
		name  string
		a     any
		wantF float64
	}{
		{
			name:  "1",
			a:     1,
			wantF: 1,
		}, {
			name:  "2",
			a:     "1",
			wantF: 1,
		}, {
			name:  "3",
			a:     math.NaN(),
			wantF: math.NaN(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parseFloat(t, tt.a)
		})
	}
}
func parseFloat(t *testing.T, a any) {
	f, err := ParseFloat64(a)
	if err != nil {
		return
	}
	switch a.(type) {
	case string:
		sf, err := strconv.ParseFloat(a.(string), 64)
		if math.IsNaN(f) && math.IsNaN(sf) {
			return
		}
		assert.Equal(t, sf, f, err)
		return
	case float32:
		if math.IsNaN(f) && math.IsNaN(float64(a.(float32))) {
			return
		}
		assert.Equal(t, float32(f), a)
		return
	}
	assert.Equalf(t, fmt.Sprint(a), fmt.Sprint(f), "ToFloat64(%v)", a)
}

func TestParseInteger(t *testing.T) {
	t.Log(ToInt[int64](-1))
	type args struct {
		v any
	}
	type testCase[T interface{ Integer }] struct {
		name string
		args args
		want T
	}
	type integer int64
	tests := []testCase[integer]{
		{name: "1", args: args{v: 23}, want: 23},
		{name: "2", args: args{v: -23}, want: -23},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := tt.args.v
			assert.Equalf(t, tt.want, ToInt[integer](v), "ToInt(%v)", v)
			parseInt(t, v)
		})
	}
}

func parseInt(t *testing.T, v any) {
	var arr []string
	var into = []any{v, ToInt[int32](v), ToInt[int](v), ToInt[int64](v), v.(int)}
	for _, in := range into {
		arr = append(arr, fmt.Sprint(in))
	}
	if !compare.AreEqual(arr...) {
		t.Error(ToInt[int32](v), ToInt[int](v), v.(int) < 0, ToInt[int64](v), v.(int))
	}
}
func parseUint(t *testing.T, v any) {
	var arr []string
	var into = []any{v, ToInt[int32](v), ToInt[int](v), ToInt[int64](v), v.(int)}
	for _, in := range into {
		arr = append(arr, fmt.Sprint(in))
	}
	if !compare.AreEqual(arr...) {
		t.Error(ToInt[int32](v), ToInt[int](v), v.(int) < 0, ToInt[int64](v), v.(int))
	}
}

func FuzzParseInt(f *testing.F) {
	var zanies = uint(18446744073709551115)
	//testLargeUint(f, uint(9223372036854775807+1e9))
	//testLargeUint(f, zanies)
	f.Add(zanies)
	f.Fuzz(func(t *testing.T, a uint) {
		parseUint(t, int(a))
	})
}

func testLargeUint(f *testing.F, zanies uint) {
	f.Log(zanies, int(zanies), ToInt[int](zanies))
	assert.Equal(f, int(zanies), ToInt[int](zanies))
}
