package voicemeeter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVbanInStream(t *testing.T) {
	//t.Skip("skipping test")
	__vbi := newVbanInStream(0)
	t.Run("Should return a vban instream type", func(t *testing.T) {
		assert.NotNil(t, __vbi)
	})
}

func TestGetVbanOutStream(t *testing.T) {
	//t.Skip("skipping test")
	__vbo := newVbanOutStream(0)
	t.Run("Should return a vban outstream type", func(t *testing.T) {
		assert.NotNil(t, __vbo)
	})
}
