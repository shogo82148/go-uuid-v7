package gouuidv7

import (
	"cmp"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"time"
)

// A UUID is a Universally Unique Identifier as specified in RFC 9562.
//
// UUIDs are comparable, such as with the == operator.
type UUID [16]byte

// String returns the string representation of u.
//
// It uses the lowercase hex-and-dash representation defined in RFC 9562.
func (u UUID) String() string {
	b, _ := u.MarshalText()
	return string(b)
}

// MarshalText implements the [encoding.TextMarshaler] interface.
// The encoding is the same as returned by [UUID.String]
func (u UUID) MarshalText() ([]byte, error) {
	return u.AppendText(make([]byte, 0, 36))
}

// AppendText implements the [encoding.TextAppender] interface.
// The encoding is the same as returned by [UUID.String]
func (u UUID) AppendText(b []byte) ([]byte, error) {
	off := len(b)
	b = append(b, "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"...)
	dst := b[off:]
	hex.Encode(dst[0:8], u[0:4])
	hex.Encode(dst[9:13], u[4:6])
	hex.Encode(dst[14:18], u[6:8])
	hex.Encode(dst[19:23], u[8:10])
	hex.Encode(dst[24:36], u[10:16])
	return b, nil
}

// Compare compares the UUID u with v.
// If u is before v, it returns -1.
// If u is after v, it returns +1.
// If they are the same, it returns 0.
//
// Compare uses the big-endian byte order defined in
// [Section 6.11 of RFC 9562] for sorting.
//
// [Section 6.11 of RFC 9562]: https://www.rfc-editor.org/rfc/rfc9562#section-6.11
func (u UUID) Compare(v UUID) int {
	for i := range u {
		if c := cmp.Compare(u[i], v[i]); c != 0 {
			return c
		}
	}
	return 0
}

// NewV7 returns a new version 7 UUID.
//
// Version 7 UUIDs contain a timestamp in the most significant 48 bits,
// and at least 62 bits of random data.
//
// NewV7 always returns UUIDs which sort in increasing order,
// except when the system clock moves backwards.
func NewV7() UUID {
	now := time.Now()
	secs := uint64(now.Unix())
	nanos := uint64(now.Nanosecond())
	msecs := secs*1000 + nanos/1_000_000

	var u UUID
	binary.BigEndian.PutUint64(u[:], msecs<<16)
	rand.Read(u[6:])
	u.setVersion(7)
	u.setVariant(0x10)
	return u
}

func (u *UUID) setVersion(version byte) {
	u[6] = (u[6] & 0b0000_1111) | (version << 4)
}

func (u *UUID) setVariant(variant byte) {
	u[8] = (u[8] & 0b0011_1111) | (variant << 6)
}
