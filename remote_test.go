package voicemeeter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBasicRemote(t *testing.T) {
	//t.Skip("skipping test")
	__rem := NewRemote("basic")
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
	t.Run("Should return a valid command pointer", func(t *testing.T) {
		assert.NotNil(t, __rem.Command)
	})
	t.Run("Should return a valid vban pointer", func(t *testing.T) {
		assert.NotNil(t, __rem.Vban)
	})
	t.Run("Should return nil recorder pointer", func(t *testing.T) {
		assert.Nil(t, __rem.Recorder)
	})
}

func TestGetBananaRemote(t *testing.T) {
	//t.Skip("skipping test")
	__rem := NewRemote("banana")
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
	t.Run("Should return a valid command pointer", func(t *testing.T) {
		assert.NotNil(t, __rem.Command)
	})
	t.Run("Should return a valid vban pointer", func(t *testing.T) {
		assert.NotNil(t, __rem.Vban)
	})
	t.Run("Should return a valid recorder", func(t *testing.T) {
		assert.NotNil(t, __rem.Recorder)
	})
}

func TestGetPotatoRemote(t *testing.T) {
	//t.Skip("skipping test")
	__rem := NewRemote("potato")
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
	t.Run("Should return a valid command pointer", func(t *testing.T) {
		assert.NotNil(t, __rem.Command)
	})
	t.Run("Should return a valid vban pointer", func(t *testing.T) {
		assert.NotNil(t, __rem.Vban)
	})
	t.Run("Should return a valid recorder", func(t *testing.T) {
		assert.NotNil(t, __rem.Recorder)
	})
}
