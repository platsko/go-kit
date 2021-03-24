// Copyright © 2020-2021 The EVEN Solutions Developers Team

package crypto

import (
	"encoding/hex"

	json "github.com/json-iterator/go"
	"google.golang.org/protobuf/proto"

	"github.com/evenlab/go-kit/crypto/proto/pb"
	"github.com/evenlab/go-kit/equal"
	"github.com/evenlab/go-kit/errors"
)

type (
	// Signature represents signature interface.
	Signature interface {
		// Embedded equaler interface.
		equal.Equaler

		// Decode sets decoded data from protobuf message.
		Decode(*pb.Signature)

		// Encode converts data to protobuf message.
		Encode() *pb.Signature

		// Equals checks whether two signatures are the same.
		Equals(Signature) bool

		// Marshal implements marshaler interface for types
		// that can marshal themselves into bytes.
		Marshal() ([]byte, error)

		// MarshalJSON implements marshaler interface for types
		// that can marshal themselves into valid JSON.
		MarshalJSON() ([]byte, error)

		// String implements stringer interface for types
		// that can converts themselves to string format.
		String() string

		// Unmarshal implements unmarshaler interface for types
		// that can unmarshal bytes of themselves.
		Unmarshal(b []byte) error

		// UnmarshalJSON implements unmarshaler interface for types
		// that can unmarshal a JSON description of themselves.
		UnmarshalJSON([]byte) error
	}

	// signature implements signature interface.
	signature struct {
		blob []byte
	}
)

var (
	// Make sure signature implements Signature interface.
	_ Signature = (*signature)(nil)
)

// NewSignature returns Signature interface.
func NewSignature(blob []byte) Signature {
	return &signature{blob: blob}
}

// DecodeSignature decodes a protobuf encoded message.
func DecodeSignature(pbuf *pb.Signature) Signature {
	return &signature{blob: pbuf.Blob}
}

// Decode implements Signature.Decode method of interface.
func (c *signature) Decode(pbuf *pb.Signature) {
	c.blob = pbuf.Blob
}

// Encode implements Signature.Encode method of interface.
func (c *signature) Encode() *pb.Signature {
	return &pb.Signature{Blob: c.blob}
}

// Equals implements Equaler.Equals method of interface.
func (c *signature) Equals(sign Signature) bool {
	return equal.BasicEqual(c, sign)
}

// Marshal implements Signature.Marshal method of interface.
func (c *signature) Marshal() ([]byte, error) {
	pbuf := c.Encode()

	return proto.Marshal(pbuf)
}

// MarshalJSON implements Signature.MarshalJSON method of interface.
func (c *signature) MarshalJSON() ([]byte, error) {
	pbuf := c.Encode()

	return json.Marshal(pbuf)
}

// Raw implements method of equal.Equaler interface.
func (c *signature) Raw() ([]byte, error) {
	if c.blob == nil {
		return nil, errors.ErrNilPointerValue()
	}

	b := make([]byte, len(c.blob))
	copy(b, c.blob)

	return b, nil
}

// String implements Signature.String method of interface.
func (c *signature) String() string {
	if c.blob == nil {
		return "<nil>"
	}

	return hex.EncodeToString(c.blob)
}

// Unmarshal implements Signature.Unmarshal method of interface.
func (c *signature) Unmarshal(b []byte) error {
	pbuf := pb.Signature{}
	err := proto.Unmarshal(b, &pbuf)
	if err != nil {
		return err
	}

	c.Decode(&pbuf)

	return nil
}

// UnmarshalJSON implements Signature.UnmarshalJSON method of interface.
func (c *signature) UnmarshalJSON(b []byte) error {
	pbuf := pb.Signature{}
	if err := json.Unmarshal(b, &pbuf); err != nil {
		return err
	}

	c.Decode(&pbuf)

	return nil
}
