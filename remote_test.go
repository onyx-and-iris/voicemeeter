package voicemeeter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBasicRemote(t *testing.T) {
	//t.Skip("skipping test")
	__rem, _ := NewRemote("basic", 0)
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
	__rem, _ := NewRemote("banana", 0)
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
	__rem, _ := NewRemote("potato", 0)
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

func TestSetAndGetFloatParameter(t *testing.T) {
	//t.Skip("skipping test")
	var param = "strip[0].mute"
	var exp = float64(1)
	vm.SetFloat(param, 1)
	t.Run("Should get a float parameter", func(t *testing.T) {
		val, _ := vm.GetFloat(param)
		assert.Equal(t, exp, val)
	})
}

func TestSetAndGetStringParameter(t *testing.T) {
	//t.Skip("skipping test")
	var param = "strip[0].label"
	var exp = "test0"
	vm.SetString(param, exp)
	t.Run("Should get a string parameter", func(t *testing.T) {
		val, _ := vm.GetString(param)
		assert.Equal(t, exp, val)
	})
}
