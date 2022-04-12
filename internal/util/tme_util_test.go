package util

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeToString(t *testing.T) {
	t.Run("should return with timezone", func(t *testing.T) {
		utcTime := time.UnixMilli(1438167001716).UTC()
		result := TimeToString(utcTime)
		assert.Equal(t, "Wed Jul 29 10:50:01 2015 UTC", result)
	})

	t.Run("should return invalid time format message", func(t *testing.T) {
		result := TimeToString(1234)
		assert.Equal(t, "invalid time 1234", result)
	})
}
