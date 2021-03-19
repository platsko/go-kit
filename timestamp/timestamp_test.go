// Copyright © 2020-2021 The EVEN Solutions Developers Team

package timestamp_test

import (
	"reflect"
	"testing"
	"time"

	json "github.com/json-iterator/go"
	"google.golang.org/protobuf/proto"

	. "github.com/platsko/go-kit/timestamp"
	"github.com/platsko/go-kit/timestamp/proto/pb"
)

const (
	testTimestampLayout = "1970-12-31T23:23:59.123456789Z+0000UTC"
)

func Benchmark_Now(tb *testing.B) {
	for i := 0; i < tb.N; i++ {
		_ = Now()
	}
}

func Benchmark_DecodeTimestamp(tb *testing.B) {
	ts := Now()
	pbuf, _ := ts.Encode()
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		if _, err := DecodeTimestamp(pbuf); err != nil {
			tb.Fatal(err)
		}
	}
}

func Benchmark_Timestamp_Decode(tb *testing.B) {
	ts := Now()
	pbuf, _ := ts.Encode()
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		ts := Timestamp{}
		if err := ts.Decode(pbuf); err != nil {
			tb.Fatal(err)
		}
	}
}

func Benchmark_Timestamp_Encode(tb *testing.B) {
	ts := Now()
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		if _, err := ts.Encode(); err != nil {
			tb.Fatal(err)
		}
	}
}

func Benchmark_Timestamp_Marshal(tb *testing.B) {
	ts := Now()
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		if _, err := ts.Marshal(); err != nil {
			tb.Fatal(err)
		}
	}
}

func Benchmark_Timestamp_MarshalJSON(tb *testing.B) {
	ts := Now()
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		if _, err := ts.MarshalJSON(); err != nil {
			tb.Fatal(err)
		}
	}
}

func Benchmark_Timestamp_Parse(tb *testing.B) {
	for i := 0; i < tb.N; i++ {
		ts := Timestamp{}
		if err := ts.Parse(testTimestampLayout); err != nil {
			tb.Fatal(err)
		}
	}
}

func Benchmark_Timestamp_Pretty(tb *testing.B) {
	ts := Now()
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		_ = ts.Pretty()
	}
}

func Benchmark_Timestamp_String(tb *testing.B) {
	ts := Now()
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		_ = ts.String()
	}
}

func Benchmark_Timestamp_UnixNanoStr(tb *testing.B) {
	ts := Now()
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		_ = ts.UnixNanoStr()
	}
}

func Benchmark_Timestamp_Unmarshal(tb *testing.B) {
	ts := Now()
	blob, _ := ts.Marshal()
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		ts := Timestamp{}
		if err := ts.Unmarshal(blob); err != nil {
			tb.Fatal(err)
		}
	}
}

func Benchmark_Timestamp_UnmarshalJSON(tb *testing.B) {
	ts := Now()
	blob, _ := ts.MarshalJSON()
	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		ts := Timestamp{}
		if err := ts.UnmarshalJSON(blob); err != nil {
			tb.Fatal(err)
		}
	}
}

func Test_Now(t *testing.T) {
	t.Parallel()

	zone, offset := time.Now().UTC().Zone()

	tests := [1]struct {
		name       string
		wantZone   string
		wantOffset int
	}{
		{
			name:       "UTC",
			wantZone:   zone,
			wantOffset: offset,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			zone, offset := Now().Zone()
			if zone != test.wantZone {
				t.Errorf("Now() zone: %#v | want: %#v", zone, test.wantZone)
			}
			if offset != test.wantOffset {
				t.Errorf("Now() offset: %#v | want: %#v", offset, test.wantOffset)
			}
		})
	}
}

func Test_DecodeTimestamp(t *testing.T) {
	t.Parallel()

	ts := Now()
	blob, _ := ts.MarshalBinary()

	tests := [4]struct {
		name    string
		pbuf    *pb.Timestamp
		want    Timestamp
		wantErr bool
	}{
		{
			name: "OK",
			pbuf: &pb.Timestamp{Blob: blob},
			want: ts,
		},
		{
			name: "version_Unsupported_ERR",
			pbuf: func() *pb.Timestamp {
				blob := make([]byte, 16)
				// byte on index=0 position encodes version
				blob[0] = 15 // 15 is unsupported version
				return &pb.Timestamp{Blob: blob}
			}(),
			wantErr: true,
		},
		{
			name:    "empty_BLOB_ERR",
			pbuf:    &pb.Timestamp{Blob: make([]byte, 0)},
			wantErr: true,
		},
		{
			name:    "invalid_Len_ERR",
			pbuf:    &pb.Timestamp{Blob: make([]byte, 3)},
			wantErr: true,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := DecodeTimestamp(test.pbuf)
			if (err != nil) != test.wantErr {
				t.Errorf("DecodeTimestamp() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("DecodeTimestamp() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_Timestamp_Decode(t *testing.T) {
	t.Parallel()

	ts := Now()
	blob, _ := ts.MarshalBinary()

	tests := [4]struct {
		name    string
		pbuf    *pb.Timestamp
		want    Timestamp
		wantErr bool
	}{
		{
			name: "OK",
			pbuf: &pb.Timestamp{Blob: blob},
			want: ts,
		},
		{
			name: "version_Unsupported_ERR",
			pbuf: func() *pb.Timestamp {
				blob := make([]byte, 16)
				// byte on index=0 position encodes version
				blob[0] = 15 // 15 is unsupported version
				return &pb.Timestamp{Blob: blob}
			}(),
			wantErr: true,
		},
		{
			name:    "empty_BLOB_ERR",
			pbuf:    &pb.Timestamp{Blob: make([]byte, 0)},
			wantErr: true,
		},
		{
			name:    "invalid_Len_ERR",
			pbuf:    &pb.Timestamp{Blob: make([]byte, 3)},
			wantErr: true,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := Timestamp{}
			err := got.Decode(test.pbuf)
			if (err != nil) != test.wantErr {
				t.Errorf("Decode() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Decode() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_Timestamp_Encode(t *testing.T) {
	t.Parallel()

	ts := Now()
	blob, _ := ts.Time.MarshalBinary()

	tests := [3]struct {
		name    string
		time    Timestamp
		want    *pb.Timestamp
		wantErr bool
	}{
		{
			name: "OK",
			time: ts,
			want: &pb.Timestamp{Blob: blob},
		},
		{
			name:    "fractional_Zone_Offset_ERR",
			time:    Timestamp{Time: ts.In(time.FixedZone("fractional zone offset", -1))},
			wantErr: true,
		},
		{
			name:    "unexpected_Zone_Offset_ERR",
			time:    Timestamp{Time: ts.In(time.FixedZone("unexpected zone offset", -60))},
			wantErr: true,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := test.time.Encode()
			if (err != nil) != test.wantErr {
				t.Errorf("Encode() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Encode() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_Timestamp_Marshal(t *testing.T) {
	t.Parallel()

	ts := Now()
	pbuf, _ := ts.Encode()
	want, _ := proto.Marshal(pbuf)

	tests := [3]struct {
		name    string
		time    Timestamp
		want    []byte
		wantErr bool
	}{
		{
			name: "OK",
			time: ts,
			want: want,
		},
		{
			name:    "fractional_Zone_Offset_ERR",
			time:    Timestamp{Time: ts.In(time.FixedZone("fractional zone offset", -1))},
			wantErr: true,
		},
		{
			name:    "unexpected_Zone_Offset_ERR",
			time:    Timestamp{Time: ts.In(time.FixedZone("unexpected zone offset", -60))},
			wantErr: true,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := test.time.Marshal()
			if (err != nil) != test.wantErr {
				t.Errorf("Marshal() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Marshal() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_Timestamp_MarshalJSON(t *testing.T) {
	t.Parallel()

	ts := Now()
	pbuf, _ := ts.Encode()
	want, _ := json.Marshal(pbuf)

	tests := [3]struct {
		name    string
		time    Timestamp
		want    []byte
		wantErr bool
	}{
		{
			name: "OK",
			time: ts,
			want: want,
		},
		{
			name:    "fractional_Zone_Offset_ERR",
			time:    Timestamp{Time: ts.In(time.FixedZone("fractional zone offset", -1))},
			wantErr: true,
		},
		{
			name:    "unexpected_Zone_Offset_ERR",
			time:    Timestamp{Time: ts.In(time.FixedZone("unexpected zone offset", -60))},
			wantErr: true,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got, err := test.time.MarshalJSON()
			if (err != nil) != test.wantErr {
				t.Errorf("MarshalJSON() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("MarshalJSON() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_Timestamp_Parse(t *testing.T) {
	t.Parallel()

	tests := [2]struct {
		name    string
		text    string
		want    string
		wantErr bool
	}{
		{
			name: "OK",
			text: testTimestampLayout,
			want: testTimestampLayout,
		},
		{
			name:    "ERR",
			text:    "0000-00-00T00:00:00.000000000Z-0000UTC",
			wantErr: true,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ts := Timestamp{}
			if err := ts.Parse(test.text); (err != nil) != test.wantErr {
				t.Errorf("Parse() error: %v | want: %v", err, test.wantErr)
				return
			}
			if test.wantErr {
				return
			}
			if got := ts.Format(LayoutTimestamp); got != test.want {
				t.Errorf("Parse() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_Timestamp_Pretty(t *testing.T) {
	t.Parallel()

	ts := Timestamp{}
	if err := ts.Parse(testTimestampLayout); err != nil {
		t.Errorf("Parse() error: %v", err)
		return
	}

	tests := [1]struct {
		name string
		time Timestamp
		want string
	}{
		{
			name: "OK",
			time: ts,
			want: "Thu Dec 31 23:23:59 +0000 1970",
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.time.Pretty(); got != test.want {
				t.Errorf("Pretty() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_Timestamp_String(t *testing.T) {
	t.Parallel()

	ts := Timestamp{}
	if err := ts.Parse(testTimestampLayout); err != nil {
		t.Errorf("Parse() error: %v", err)
		return
	}

	tests := [1]struct {
		name string
		time Timestamp
		want string
	}{
		{
			name: "OK",
			time: ts,
			want: testTimestampLayout,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.time.String(); got != test.want {
				t.Errorf("String() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_Timestamp_UnixNanoStr(t *testing.T) {
	t.Parallel()

	ts := Timestamp{}
	if err := ts.Parse(testTimestampLayout); err != nil {
		t.Errorf("Parse() error: %v", err)
		return
	}

	tests := [1]struct {
		name string
		time Timestamp
		want string
	}{
		{
			name: "OK",
			time: ts,
			want: "31533839123456789",
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			if got := test.time.UnixNanoStr(); got != test.want {
				t.Errorf("UnixNanoStr() got: %v | want: %v", got, test.want)
			}
		})
	}
}

func Test_Timestamp_Unmarshal(t *testing.T) {
	t.Parallel()

	ts := Now()
	blob, _ := ts.Marshal()

	tests := [2]struct {
		name    string
		blob    []byte
		want    Timestamp
		wantErr bool
	}{
		{
			name: "OK",
			blob: blob,
			want: ts,
		},
		{
			name:    "invalid_BLOB_ERR",
			blob:    []byte(":"), // invalid data
			wantErr: true,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := Timestamp{}
			if err := got.Unmarshal(test.blob); (err != nil) != test.wantErr {
				t.Errorf("Unmarshal() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Unmarshal() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}

func Test_Timestamp_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	ts := Now()
	blob, _ := ts.MarshalJSON()

	tests := [2]struct {
		name    string
		blob    []byte
		want    Timestamp
		wantErr bool
	}{
		{
			name: "OK",
			blob: blob,
			want: ts,
		},
		{
			name:    "invalid_JSON_ERR",
			blob:    []byte(":"), // invalid json
			wantErr: true,
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			got := Timestamp{}
			if err := got.UnmarshalJSON(test.blob); (err != nil) != test.wantErr {
				t.Errorf("UnmarshalJSON() error: %v | want: %v", err, test.wantErr)
				return
			}
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("UnmarshalJSON() got: %#v | want: %#v", got, test.want)
			}
		})
	}
}
