package voicemeeter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVMPath(t *testing.T) {
	//t.Skip("skipping test")
	_, err := dllPath()
	t.Run("Should return err as nil", func(t *testing.T) {
		assert.Nil(t, err)
	})
}
