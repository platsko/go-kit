// Copyright © 2020-2021 The EVEN Solutions Developers Team

package timestamp

import (
	"fmt"
	"time"

	json "github.com/json-iterator/go"
	"google.golang.org/protobuf/proto"

	"github.com/evenlab/go-kit/timestamp/proto/pb"
)

const (
	LayoutTimestamp = "2006-01-02T15:04:05.999999999Z-0700MST"
)

type (
	// Timestamp wraps time.
	Timestamp struct {
		time.Time
	}
)

// Now returns Timestamp initialized with current UTC time.
func Now() Timestamp {
	ts := time.Now().UTC()

	return Timestamp{Time: ts}
}

// DecodeTimestamp decodes a protobuf encoded message.
func DecodeTimestamp(pbuf *pb.Timestamp) (ts Timestamp, err error) {
	err = ts.Decode(pbuf)

	return ts, err
}

// Decode converts protobuf message into object.
func (t *Timestamp) Decode(pbuf *pb.Timestamp) error {
	return t.Time.UnmarshalBinary(pbuf.Blob)
}

// Encode converts object into protobuf message.
func (t *Timestamp) Encode() (*pb.Timestamp, error) {
	blob, err := t.MarshalBinary()
	if err != nil {
		return nil, err
	}

	return &pb.Timestamp{Blob: blob}, nil
}

// Marshal implements marshaler interface for types
// that can marshal themselves into bytes.
func (t *Timestamp) Marshal() ([]byte, error) {
	pbuf, err := t.Encode()
	if err != nil {
		return nil, err
	}

	return proto.Marshal(pbuf)
}

// MarshalJSON implements marshaler interface for types
// that can marshal themselves into valid JSON.
func (t *Timestamp) MarshalJSON() ([]byte, error) {
	pbuf, err := t.Encode()
	if err != nil {
		return nil, err
	}

	return json.Marshal(pbuf)
}

// Parse sets Timestamp form LayoutTimestamp string.
func (t *Timestamp) Parse(value string) error {
	ts, err := time.Parse(LayoutTimestamp, value)
	if err != nil {
		return err
	}

	t.Time = ts

	return nil
}

// Pretty returns human-readable string.
func (t *Timestamp) Pretty() string {
	return t.Time.Format(time.RubyDate)
}

// String implements stringer interface.
func (t *Timestamp) String() string {
	return t.Time.Format(LayoutTimestamp)
}

// UnixNanoStr returns string representation as string of UnixNano.
func (t *Timestamp) UnixNanoStr() string {
	return fmt.Sprintf("%v", t.Time.UnixNano())
}

// Unmarshal implements unmarshaler interface for types
// that can unmarshal bytes of themselves.
func (t *Timestamp) Unmarshal(b []byte) error {
	pbuf := pb.Timestamp{}
	if err := proto.Unmarshal(b, &pbuf); err != nil {
		return err
	}

	return t.Decode(&pbuf)
}

// UnmarshalJSON implements unmarshaler interface for types
// that can unmarshal a JSON description of themselves.
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	pbuf := pb.Timestamp{}
	if err := json.Unmarshal(b, &pbuf); err != nil {
		return err
	}

	return t.Decode(&pbuf)
}
