package num

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/conbanwa/logs"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
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
	assert.Equalf(t, fmt.Sprint(a), fmt.Sprint(f), "ParseFloat64(%v)", a)
}

func BenchmarkFloatToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FloatToString(12.232321, 0.001)
	}
}

func TestFloatToString(t *testing.T) {
	assert.Equal(t, "1", FloatToString(1.10231000, 1))
	assert.Equal(t, "0.102", FloatToString(0.10231000, 0.001))
	assert.Equal(t, "189.61", FloatToString(189.61020000, 0.01))
	assert.NotEqual(t, "1.10231000", FloatToString(1.10231000, 1e-8))
	assert.Equal(t, 0.129999, FloatToFixed(0.1299999, 1e-6))
	logs.Assert(FloatToFixed(1.10231000, 0.000000001) != 1.10231000, FloatToFixed(1.10231000, 0.000000001))
	logs.Assert(0.102-FloatToFixed(0.10231000, 0.001) < 0.00000000000001, FloatToFixed(0.10231000, 0.001))
	assert.Equal(t, FloatToString(1.10231000, 1), "1")
	assert.Equal(t, FloatToString(1.10231000, 0), "1.10231")
	assert.Equal(t, FloatToString(0.10231000, 0.001), "0.102")
	assert.Equal(t, FloatToString(189.61020000, 0.01), "189.61")
	assert.Equal(t, FloatToString(1.10231000, 1e-7), "1.1023100")
	assert.Equal(t, FloatToString(0.1299999, 0.0001), "0.1299")
	assert.Equal(t, FloatToString(6.7597, 0.01), "6.75")
	assert.Equal(t, FloatToFixed(1.10231000, 1), 1.0)
	assert.Equal(t, FloatToFixed(1.10231000, 0), 1.10231)
	assert.Equal(t, FloatToFixed(189.61020000, 0.01), 189.61)
	logs.ErrorIfNotSame(FloatToFixed(0.10231000, 0.001), 0.102)
	logs.ErrorIfNotSame(FloatToString(1.10231000, 1e-8), "1.10231")
	logs.ErrorIfNotSame(FloatToFixed(1.10231000, 0.000000001), 1.10231)
	logs.ErrorIfNotSame(FloatToFixed(0.1299999, 0.0001), 0.1299)
	logs.ErrorIfNotSame(FloatToFixed(6.7597, 0.01), 6.75)
	type args struct {
		v    float64
		step float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "0",
			args: args{v: 1111.55, step: 0.1},
			want: "1111.5",
		},
		{
			name: "1",
			args: args{v: 3341.055, step: 0.01},
			want: "3341.05",
		},
		{
			name: "2",
			args: args{v: 61.0555, step: 0.001},
			want: "61.055",
		},
		{
			name: "3",
			args: args{v: 5551.0555, step: 10},
			want: "5550",
		},
		{
			name: "4",
			args: args{v: 441.0555, step: 100},
			want: "400",
		},
		{
			name: "5",
			args: args{v: 2.9999999, step: 0.1},
			want: "2.9",
		},
		{
			name: "6",
			args: args{v: 1.9999, step: 1},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloatToString(tt.args.v, tt.args.step); got != tt.want {
				t.Errorf("FloatToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloatToFixed(t *testing.T) {
	type args struct {
		v    float64
		step float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "0",
			args: args{v: 1.0555, step: 0.1},
			want: 1,
		},
		{
			name: "1",
			args: args{v: 1.0555, step: 0.01},
			want: 1.05,
		},
		{
			name: "2",
			args: args{v: 1.0555, step: 0.001},
			want: 1.055,
		},
		{
			name: "3",
			args: args{v: 5551.0555, step: 10},
			want: 5550,
		},
		{
			name: "4",
			args: args{v: 441.0555, step: 100},
			want: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, FloatToFixed(tt.args.v, tt.args.step), "FloatToFixed(%v, %v)", tt.args.v, tt.args.step)
		})
	}
}

func TestParseInteger(t *testing.T) {
	t.Log(ParseInteger[int64](-1))
	type args struct {
		v any
	}
	type testCase[T interface{ constraints.Integer }] struct {
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
			assert.Equalf(t, tt.want, ParseInteger[integer](v), "ParseInteger(%v)", v)
			parseInt(t, v)
		})
	}
}

func parseInt(t *testing.T, v any) {
	var arr []string
	var into = []any{v, ParseInteger[int32](v), ParseInteger[int](v), ParseInteger[int64](v), v.(int)}
	for _, in := range into {
		arr = append(arr, fmt.Sprint(in))
	}
	if !logs.Uniform(arr...) {
		t.Error(ParseInteger[int32](v), ParseInteger[int](v), v.(int) < 0, ParseInteger[int64](v), v.(int))
	}
}
func parseUint(t *testing.T, v any) {
	var arr []string
	var into = []any{v, ParseInteger[int32](v), ParseInteger[int](v), ParseInteger[int64](v), v.(int)}
	for _, in := range into {
		arr = append(arr, fmt.Sprint(in))
	}
	if !logs.Uniform(arr...) {
		t.Error(ParseInteger[int32](v), ParseInteger[int](v), v.(int) < 0, ParseInteger[int64](v), v.(int))
	}
}

func FuzzParseInt(f *testing.F) {
	var zanies = uint(18446744073709551115)
	testLargeUint(f, uint(9223372036854775807+1e9))
	testLargeUint(f, zanies)
	f.Add(zanies)
	f.Fuzz(func(t *testing.T, a uint) {
		parseUint(t, int(a))
	})
}

func testLargeUint(f *testing.F, zanies uint) {
	f.Log(zanies, int(zanies), ParseInteger[int](zanies))
	assert.Equal(f, int(zanies), ParseInteger[int](zanies))
}
