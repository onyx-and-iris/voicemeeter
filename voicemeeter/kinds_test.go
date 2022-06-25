package voicemeeter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetKindBasic(t *testing.T) {
	//t.Skip("skipping test")
	__kind := newBasicKind()
	t.Run("Should return a basic kind", func(t *testing.T) {
		assert.NotNil(t, __kind)
	})
	t.Run("Should equal 'Basic'", func(t *testing.T) {
		assert.Equal(t, "Basic", __kind.String())
	})
}

func TestGetKindBanana(t *testing.T) {
	//t.Skip("skipping test")
	__kind := newBananaKind()
	t.Run("Should return a banana kind", func(t *testing.T) {
		assert.NotNil(t, __kind)
	})
	t.Run("Should return 'Banana'", func(t *testing.T) {
		assert.Equal(t, "Banana", __kind.String())
	})
}

func TestGetKindPotato(t *testing.T) {
	//t.Skip("skipping test")
	__kind := newPotatoKind()
	t.Run("Should return a potato kind", func(t *testing.T) {
		assert.NotNil(t, __kind)
	})
	t.Run("Should return 'Potato'", func(t *testing.T) {
		assert.Equal(t, "Potato", __kind.String())
	})
}
