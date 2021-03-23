// Copyright © 2020-2021 The EVEN Solutions Developers Team

package bytes_test

import (
	"testing"

	. "github.com/platsko/go-kit/bytes"
)

func Benchmark_Equal(b *testing.B) {
	const size = 1024
	vx, vy := make([]byte, size), RandBytes(size)
	copy(vx, vy)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Equal(vx, vy)
	}
}

func Test_Equal(t *testing.T) {
	t.Parallel()

	const size = 1024
	vx, vy := make([]byte, size), RandBytes(size)
	copy(vx, vy)

	tests := [7]struct {
		name string
		vx   []byte
		vy   []byte
		want bool
	}{
		{
			name: "TRUE",
			vx:   vx,
			vy:   vy,
			want: true,
		},
		{
			name: "nil_VX_&_VY_TRUE",
			vx:   nil,
			vy:   nil,
			want: true,
		},
		{
			name: "zeros_VX_&_VY_TRUE",
			vx:   []byte{},
			vy:   []byte{},
			want: true,
		},
		{
			name: "FALSE",
			vx:   vx,
			vy:   RandBytes(size),
		},
		{
			name: "nil_VX_FALSE",
			vx:   nil,
			vy:   vy,
		},
		{
			name: "nil_VY_FALSE",
			vx:   vx,
			vy:   nil,
		},
		{
			name: "len_diff_FALSE",
			vx:   []byte{},
			vy:   []byte{255},
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := Equal(test.vx, test.vy); got != test.want {
				t.Errorf("Equal() got: %v | want: %v", got, test.want)
			}
		})
	}
}
