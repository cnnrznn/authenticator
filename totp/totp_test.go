package totp

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStep(t *testing.T) {
	step, _ := calcStep(time.Unix(90, 0))
	assert.Equal(t,
		uint64(3),
		step,
	)

	step, _ = calcStep(time.Unix(91, 0))
	assert.Equal(t,
		uint64(3),
		step,
	)

	step, _ = calcStep(time.Unix(119, 0))
	assert.Equal(t,
		uint64(3),
		step,
	)

	step, _ = calcStep(time.Unix(120, 0))
	assert.Equal(t,
		uint64(4),
		step,
	)
}

func FuzzStep(f *testing.F) {
	f.Fuzz(func(t *testing.T, i int64) {
		tx := time.Unix(i, 0)
		s, _ := calcStep(tx)
		expected := uint64(i / 30)
		assert.Equal(t, expected, s)
	})
}

func TestTOTP(t *testing.T) {
	s, _, err := Generate("JBSWY3DPEHPK3PXP", time.Unix(0, 0))

	assert.NoError(t, err)
	assert.Equal(t, "282-760", s)
}
