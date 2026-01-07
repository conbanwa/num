package num

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkFloatToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FloatToString(12.232321, 0.001)
	}
}
func BenchmarkSprint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AfterDot(125635465465464566546.5555, 2)
	}
}

func AfterDot(f float64, precision int) string {
	s := fmt.Sprint(f)
	for i, r := range s {
		if r == '.' {
			if precision == 0 {
				return s[:i]
			}
			return s[:min(i+precision+1, len(s)-1)]
		}
	}
	return s
}

func TestFloatToString(t *testing.T) {
	assert.Equal(t, "1", FloatToString(1.10231000, 1))
	assert.Equal(t, "0.102", FloatToString(0.10231000, 0.001))
	assert.Equal(t, "189.61", FloatToString(189.61020000, 0.01))
	assert.NotEqual(t, "1.10231000", FloatToString(1.10231000, 1e-8))
	assert.Equal(t, 0.129999, FloatToFixed(0.1299999, 1e-6))
	assert.Equal(t, FloatToString(0.10231000, 0.001), "0.102")
	assert.Equal(t, FloatToString(1.10231000, 1), "1")
	assert.Equal(t, FloatToString(1.10231000, 0), "1.10231")
	assert.Equal(t, FloatToString(189.61020000, 0.01), "189.61")
	assert.Equal(t, FloatToString(1.10231000, 1e-7), "1.1023100")
	assert.Equal(t, FloatToString(0.1299999, 0.0001), "0.1299")
	assert.Equal(t, FloatToString(6.7597, 0.01), "6.75")
	assert.Equal(t, FloatToFixed(1.10231000, 1), 1.0)
	assert.Equal(t, FloatToFixed(1.10231000, 0), 1.10231)
	assert.Equal(t, FloatToFixed(189.61020000, 0.01), 189.61)
	// logs.ErrorIfNotSame(FloatToFixed(0.10231000, 0.001), 0.102)
	// logs.ErrorIfNotSame(FloatToString(1.10231000, 1e-8), "1.10231")
	// logs.ErrorIfNotSame(FloatToFixed(1.10231000, 0.000000001), 1.10231)
	// logs.ErrorIfNotSame(FloatToFixed(0.1299999, 0.0001), 0.1299)
	// logs.ErrorIfNotSame(FloatToFixed(1.10231000, 0.000000001), 1.10231000, FloatToFixed(1.10231000, 0.000000001))
	// logs.ErrorIfNotSame(FloatToFixed(0.10231000, 0.001), 0.102, FloatToFixed(0.10231000, 0.001))
	assert.Equal(t, FloatToFixed(6.7597, 0.01), 6.75)
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
