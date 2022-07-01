package voicemeeter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOutputs(t *testing.T) {
	//t.Skip("skipping test")
	__o := newOutputs("strip[0]", 0)
	t.Run("Should return an output type", func(t *testing.T) {
		assert.NotNil(t, __o)
	})
}
