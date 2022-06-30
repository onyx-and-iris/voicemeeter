package voicemeeter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPhysStrip(t *testing.T) {
	//t.Skip("skipping test")
	__strip := newPhysicalStrip(0, newPotatoKind())
	t.Run("Should return a physical strip type", func(t *testing.T) {
		assert.NotNil(t, __strip)
	})
	t.Run("Should return 'PhysicalStrip0'", func(t *testing.T) {
		assert.Equal(t, "PhysicalStrip0", __strip.String())
	})
}

func TestGetVirtStrip(t *testing.T) {
	//t.Skip("skipping test")
	__strip := newVirtualStrip(4, newPotatoKind())
	t.Run("Should return a basic kind", func(t *testing.T) {
		assert.NotNil(t, __strip)
	})
	t.Run("Should return 'VirtualStrip4'", func(t *testing.T) {
		assert.Equal(t, "VirtualStrip4", __strip.String())
	})
}
