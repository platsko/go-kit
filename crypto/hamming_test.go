// Copyright © 2020-2021 The EVEN Solutions Developers Team

package crypto_test

import (
	"testing"

	"github.com/platsko/go-kit/bytes"
	"github.com/platsko/go-kit/crypto"
)

func Benchmark_Hamming(tb *testing.B) {
	vx := bytes.RandBytes(256)
	vy := bytes.RandBytes(256)
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		if _, err := crypto.Hamming(vx, vy); err != nil {
			tb.Fatal(err)
		}
	}
}

func Test_Hamming(t *testing.T) {
	t.Parallel()

	tests := [7]struct {
		name    string
		vx      []byte
		vy      []byte
		want    int
		wantErr bool
	}{
		{
			name: "OK",
			vx:   []byte{1},
			vy:   []byte{15},
			want: 3,
		},
		{
			name:    "different_Size_ERR",
			vx:      make([]byte, 1),
			vy:      make([]byte, 2),
			wantErr: true,
		},
		{
			name:    "zero_Size_VX_ERR",
			vx:      make([]byte, 0),
			vy:      make([]byte, 1),
			wantErr: true,
		},
		{
			name:    "zero_Size_VY_ERR",
			vx:      make([]byte, 1),
			vy:      make([]byte, 0),
			wantErr: true,
		},
		{
			name:    "nil_VX_ERR",
			vx:      nil,
			vy:      make([]byte, 1),
			wantErr: true,
		},
		{
			name:    "nil_VY_ERR",
			vx:      make([]byte, 1),
			vy:      nil,
			wantErr: true,
		},
		{
			// Identity axiom:
			// (Hamming(vx, vy) == 0) && (vx == vy)
			name: "Identity_Axiom_OK",
			vx:   []byte{15},
			vy:   []byte{15},
			want: 0,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := crypto.Hamming(test.vx, test.vy)
			if (err != nil) != test.wantErr {
				t.Errorf("Hamming() error: %v | want: %v", err, test.wantErr)
				return
			}
			if got != test.want {
				t.Errorf("Hamming() got: %v | want: %v", got, test.want)
			}
		})
	}
}

// Symmetry axiom:
// Hamming(vx, vy) == Hamming(vy, vx).
func Test_SymmetryAxiom(t *testing.T) {
	t.Parallel()

	tests := [1]struct {
		name string
		vx   []byte
		vy   []byte
	}{
		{
			name: "SymmetryAxiom_OK",
			vx:   []byte{1},
			vy:   []byte{15},
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			distXY, _ := crypto.Hamming(test.vx, test.vy)
			distYX, _ := crypto.Hamming(test.vy, test.vx)
			if !(distXY == distYX) {
				t.Errorf("Hamming() does not correspond to the symmetry axiom")
			}
		})
	}
}

// Triangular axiom:
// Hamming(vx, vz) <= Hamming(vx, vy) + Hamming(vy, vz).
func Test_TriangularAxiom(t *testing.T) {
	t.Parallel()

	tests := [1]struct {
		name string
		vx   []byte
		vy   []byte
		vz   []byte
	}{
		{
			name: "TriangularAxiom_OK",
			vx:   []byte{1},
			vy:   []byte{15},
			vz:   []byte{3},
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			distXY, _ := crypto.Hamming(test.vx, test.vy)
			distXZ, _ := crypto.Hamming(test.vx, test.vz)
			distYZ, _ := crypto.Hamming(test.vy, test.vz)
			if !(distXY+distYZ >= distXZ) {
				t.Errorf("Hamming() does not correspond to the triangular axiom")
			}
		})
	}
}