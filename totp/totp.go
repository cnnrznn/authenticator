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

func Generate(secret string, t time.Time) (string, int, error) {
	counter := make([]byte, 8)
	step, remaining := calcStep(t)
	binary.BigEndian.PutUint64(counter, step)

	key, err := base32.StdEncoding.WithPadding(base32.NoPadding).DecodeString(secret)
	if err != nil {
		return "", 0, err
	}

	token, err := hotp(counter, key)
	if err != nil {
		return "", 0, err
	}

	return token, remaining, nil
}

func calcStep(t time.Time) (uint64, int) {
	period := int64(30)
	t0 := t.Unix() / period
	t1 := t0*period + period
	return uint64(t0), int(t1 - t.Unix())
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
