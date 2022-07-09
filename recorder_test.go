package voicemeeter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRecorder(t *testing.T) {
	//t.Skip("skipping test")
	__rec := newRecorder()
	t.Run("Should return a button type", func(t *testing.T) {
		assert.NotNil(t, __rec)
	})
}
