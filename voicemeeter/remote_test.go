package voicemeeter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBasicRemote(t *testing.T) {
	//t.Skip("skipping test")
	__rem := GetRemote("basic")
	t.Run("Should return a remote basic type", func(t *testing.T) {
		assert.NotNil(t, __rem)
	})
	t.Run("Should equal 'Voicemeeter Basic'", func(t *testing.T) {
		assert.Equal(t, "Voicemeeter Basic", __rem.String())
	})
	t.Run("Should strip length equal 3", func(t *testing.T) {
		assert.Equal(t, 3, len(__rem.Strip))
	})
	t.Run("Should bus length equal 2", func(t *testing.T) {
		assert.Equal(t, 2, len(__rem.Bus))
	})
}

func TestGetBananaRemote(t *testing.T) {
	//t.Skip("skipping test")
	__rem := GetRemote("banana")
	t.Run("Should return a remote banana type", func(t *testing.T) {
		assert.NotNil(t, __rem)
	})
	t.Run("Should equal 'Voicemeeter Banana'", func(t *testing.T) {
		assert.Equal(t, "Voicemeeter Banana", __rem.String())
	})
	t.Run("Should strip length equal 5", func(t *testing.T) {
		assert.Equal(t, 5, len(__rem.Strip))
	})
	t.Run("Should bus length equal 5", func(t *testing.T) {
		assert.Equal(t, 5, len(__rem.Bus))
	})
}

func TestGetPotatoRemote(t *testing.T) {
	//t.Skip("skipping test")
	__rem := GetRemote("potato")
	t.Run("Should return a remote basic type", func(t *testing.T) {
		assert.NotNil(t, __rem)
	})
	t.Run("Should equal 'Voicemeeter Potato'", func(t *testing.T) {
		assert.Equal(t, "Voicemeeter Potato", __rem.String())
	})
	t.Run("Should strip length equal 8", func(t *testing.T) {
		assert.Equal(t, 8, len(__rem.Strip))
	})
	t.Run("Should bus length equal 8", func(t *testing.T) {
		assert.Equal(t, 8, len(__rem.Bus))
	})
}
