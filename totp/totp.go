package totp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"math"
	"time"
)

func Generate(secret string, t time.Time) (string, error) {
	counter := make([]byte, 8)
	binary.BigEndian.PutUint64(counter, step(t))

	key, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(secret)
	if err != nil {
		return "", err
	}

	return hotp(counter, key)
}

func step(t time.Time) uint64 {
	return uint64(t.Unix()) / 30
}

func hotp(counter, secret []byte) (string, error) {
	mac := hmac.New(sha1.New, secret)
	mac.Write(counter)
	sum := mac.Sum(nil)

	// Dynamic truncation
	offset := sum[len(sum)-1] & 0xf
	code := int(sum[offset]&0x7f)<<24 |
		int(sum[offset+1]&0xff)<<16 |
		int(sum[offset+2]&0xff)<<8 |
		int(sum[offset+3]&0xff)

	result := code % int(math.Pow10(6))

	return fmt.Sprintf("%06d", result), nil
}
