package voicemeeter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCommand(t *testing.T) {
	//t.Skip("skipping test")
	__c := newCommand()
	t.Run("Should return a command type", func(t *testing.T) {
		assert.NotNil(t, __c)
	})
}
