package voicemeeter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetButton(t *testing.T) {
	//t.Skip("skipping test")
	__mb := newButton(0)
	t.Run("Should return a button type", func(t *testing.T) {
		assert.NotNil(t, __mb)
	})
}
