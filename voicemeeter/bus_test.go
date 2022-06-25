package voicemeeter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPhysBus(t *testing.T) {
	//t.Skip("skipping test")
	__bus := newPhysicalBus(0)
	t.Run("Should return a physical bus type", func(t *testing.T) {
		assert.NotNil(t, __bus)
	})
	t.Run("Should return 'PhysicalBus0'", func(t *testing.T) {
		assert.Equal(t, "PhysicalBus0", __bus.String())
	})
}

func TestGetVirtBus(t *testing.T) {
	//t.Skip("skipping test")
	__bus := newVirtualBus(4)
	t.Run("Should return a basic kind", func(t *testing.T) {
		assert.NotNil(t, __bus)
	})
	t.Run("Should return 'VirtualBus4'", func(t *testing.T) {
		assert.Equal(t, "VirtualBus4", __bus.String())
	})
}
