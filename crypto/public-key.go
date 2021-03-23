// Copyright © 2020-2021 The EVEN Solutions Developers Team

package crypto

import (
	"encoding/base64"
	"encoding/hex"

	json "github.com/json-iterator/go"
	"github.com/libp2p/go-libp2p-core/crypto"
	"google.golang.org/protobuf/proto"

	"github.com/platsko/go-kit/crypto/proto/pb"
	"github.com/platsko/go-kit/equal"
	"github.com/platsko/go-kit/errors"
)

type (
	// PublicKey represents public key interface.
	PublicKey interface {
		// Embedded equaler interface.
		equal.Equaler

		// Algo returns the public key Algo.
		Algo() Algo

		// Base64 encodes a public key to base64 string.
		Base64() (string, error)

		// Decode sets decoded data from protobuf message.
		Decode(*pb.PublicKey) error

		// Encode converts data to protobuf message.
		Encode() (*pb.PublicKey, error)

		// Equals checks whether two public keys are the same.
		Equals(PublicKey) bool

		// Hash224 calculates SHA224 checksum
		// over SHA256 checksum over public key value.
		Hash224() (Hash224, error)

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
		Unmarshal([]byte) error

		// UnmarshalJSON implements unmarshaler interface for types
		// that can unmarshal a JSON description of themselves.
		UnmarshalJSON([]byte) error

		// Verify verifies signable object.
		Verify(Signable) (bool, error)
	}

	// publicKey implements PublicKey interface.
	publicKey struct {
		ki crypto.PubKey
	}
)

var (
	// Make sure publicKey implements PublicKey interface.
	_ PublicKey = (*publicKey)(nil)
)

// NewPublicKey returns PublicKey interface.
func NewPublicKey(ki crypto.PubKey) PublicKey {
	return &publicKey{ki: ki}
}

// DecodePublicKey decodes a protobuf encoded message.
func DecodePublicKey(pbuf *pb.PublicKey) (PublicKey, error) {
	pbKey := publicKey{}
	if err := pbKey.Decode(pbuf); err != nil {
		return nil, err
	}

	return &pbKey, nil
}

// Algo implements PublicKey.Algo method of interface.
func (c *publicKey) Algo() Algo {
	if c.ki == nil {
		return UNKNOWN
	}

	return Algo(c.ki.Type())
}

// Base64 implements PublicKey.Base64 method of interface.
func (c *publicKey) Base64() (string, error) {
	b, err := c.Raw()
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

// Decode implements PublicKey.Decode method of interface.
func (c *publicKey) Decode(pbuf *pb.PublicKey) error {
	ki, err := crypto.UnmarshalPublicKey(pbuf.Blob)
	if err != nil {
		return err
	}

	c.ki = ki

	return nil
}

// Encode implements PublicKey.Encode method of interface.
func (c *publicKey) Encode() (*pb.PublicKey, error) {
	if c.ki == nil {
		return nil, ErrPublicKeyCannotBeNil()
	}

	blob, err := crypto.MarshalPublicKey(c.ki)
	if err != nil {
		return nil, err
	}

	pbuf := pb.PublicKey{Blob: blob}

	return &pbuf, nil
}

// Equals implements PublicKey.Equals method of interface.
func (c *publicKey) Equals(pbKey PublicKey) bool {
	return equal.BasicEqual(c, pbKey)
}

// Hash224 implements PublicKey.Hash224 method of interface.
func (c *publicKey) Hash224() (Hash224, error) {
	b, err := c.Raw()
	if err != nil {
		return Hash224{}, err
	}

	return NewHash224(b), nil
}

// Marshal implements PublicKey.Marshal method of interface.
func (c *publicKey) Marshal() ([]byte, error) {
	pbuf, err := c.Encode()
	if err != nil {
		return nil, err
	}

	return proto.Marshal(pbuf)
}

// MarshalJSON implements PublicKey.MarshalJSON method of interface.
func (c *publicKey) MarshalJSON() ([]byte, error) {
	pbuf, err := c.Encode()
	if err != nil {
		return nil, err
	}

	return json.Marshal(pbuf)
}

// Raw implements method of equal.Equaler interface.
func (c *publicKey) Raw() ([]byte, error) {
	if c.ki == nil {
		return nil, errors.ErrNilPointerValue()
	}

	return c.ki.Raw()
}

// String implements PublicKey.String method of interface.
func (c *publicKey) String() string {
	b, err := c.Raw()
	if err != nil {
		return "<nil>"
	}

	return hex.EncodeToString(b)
}

// Unmarshal implements PublicKey.Unmarshal method of interface.
func (c *publicKey) Unmarshal(b []byte) error {
	pbuf := new(pb.PublicKey)
	if err := proto.Unmarshal(b, pbuf); err != nil {
		return err
	}

	return c.Decode(pbuf)
}

// UnmarshalJSON implements PublicKey.UnmarshalJSON method of interface.
func (c *publicKey) UnmarshalJSON(data []byte) error {
	pbuf := pb.PublicKey{}
	if err := json.Unmarshal(data, &pbuf); err != nil {
		return err
	}

	return c.Decode(&pbuf)
}

// Verify implements PublicKey.Verify method of interface.
func (c *publicKey) Verify(signable Signable) (bool, error) {
	if c.ki == nil {
		return false, ErrPublicKeyCannotBeNil()
	}

	if signable == nil {
		return false, ErrSignableCannotBeNil()
	}

	sign := signable.GetSignature()
	if sign == nil {
		return false, ErrSignatureCannotBeNil()
	}

	blob, err := sign.Raw()
	if err != nil {
		return false, err
	}

	hash, err := signable.Hash()
	if err != nil {
		return false, err
	}

	return c.ki.Verify(hash[:], blob)
}
