package voicemeeter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEq(t *testing.T) {
	//t.Skip("skipping test")
	__e := newEq("strip[0].EQ", 0)
	t.Run("Should return an eQ type", func(t *testing.T) {
		assert.NotNil(t, __e)
	})
}
