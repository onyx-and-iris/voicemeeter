package voicemeeter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDevice(t *testing.T) {
	//t.Skip("skipping test")
	__dev := newDevice()
	t.Run("Should return a button type", func(t *testing.T) {
		assert.NotNil(t, __dev)
	})
}
