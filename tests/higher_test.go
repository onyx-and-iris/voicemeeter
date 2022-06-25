package voicemeeter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrip0Mute(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Strip[0].SetMute(true)
	t.Run("Should return true when SetMute(true)", func(t *testing.T) {
		assert.Equal(t, vmRem.Strip[0].GetMute(), true)
	})

	vmRem.Strip[0].SetMute(false)
	t.Run("Should return false when SetMute(false)", func(t *testing.T) {
		assert.Equal(t, vmRem.Strip[0].GetMute(), false)
	})
}

func TestStrip2Limit(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Strip[2].SetLimit(-8)
	t.Run("Should return -8 when SetLimit(-8)", func(t *testing.T) {
		assert.Equal(t, vmRem.Strip[2].GetLimit(), -8)
	})

	vmRem.Strip[2].SetLimit(-32)
	t.Run("Should return -32 when SetLimit(-8)", func(t *testing.T) {
		assert.Equal(t, vmRem.Strip[2].GetLimit(), -32)
	})

}

func TestStrip4Label(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Strip[4].SetLabel("test0")
	t.Run("Should return test0 when SetLimit('test0')", func(t *testing.T) {
		assert.Equal(t, vmRem.Strip[4].GetLabel(), "test0")
	})

	vmRem.Strip[4].SetLabel("test1")
	t.Run("Should return test1 when SetLimit('test1')", func(t *testing.T) {
		assert.Equal(t, vmRem.Strip[4].GetLabel(), "test1")
	})
}

func TestStrip5Gain(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Strip[4].SetGain(-20.8)
	t.Run("Should return -20.8 when SetGain(-20.8)", func(t *testing.T) {
		assert.Equal(t, vmRem.Strip[4].GetGain(), -20.8)
	})

	vmRem.Strip[4].SetGain(-3.6)
	t.Run("Should return -3.6 when SetGain(-3.6)", func(t *testing.T) {
		assert.Equal(t, vmRem.Strip[4].GetGain(), -3.6)
	})
}

func TestStrip3Mc(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Strip[3].SetMc(true)
	t.Run("Should return true when SetMc(true)", func(t *testing.T) {
		assert.Equal(t, vmRem.Strip[3].GetMc(), true)
	})

	vmRem.Strip[3].SetMc(false)
	t.Run("Should return false when SetMc(false)", func(t *testing.T) {
		assert.Equal(t, vmRem.Strip[3].GetMc(), false)
	})
}

func TestBus3Eq(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Bus[3].SetEq(true)
	t.Run("Should return true when SetEq(true)", func(t *testing.T) {
		assert.Equal(t, vmRem.Bus[3].GetEq(), true)
	})

	vmRem.Bus[3].SetEq(false)
	t.Run("Should return false when SetEq(false)", func(t *testing.T) {
		assert.Equal(t, vmRem.Bus[3].GetEq(), false)
	})
}

func TestBus4Label(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Bus[4].SetLabel("test0")
	t.Run("Should return test0 when SetEq('test0')", func(t *testing.T) {
		assert.Equal(t, vmRem.Bus[4].GetLabel(), "test0")
	})

	vmRem.Bus[4].SetLabel("test1")
	t.Run("Should return test1 when SetEq('test1')", func(t *testing.T) {
		assert.Equal(t, vmRem.Bus[4].GetLabel(), "test1")
	})
}

func TestVbanInStream0On(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Vban.InStream[0].SetOn(true)
	t.Run("Should return true when SetOn(true)", func(t *testing.T) {
		assert.Equal(t, vmRem.Vban.InStream[0].GetOn(), true)
	})

	vmRem.Vban.InStream[0].SetOn(false)
	t.Run("Should return false when SetOn(false)", func(t *testing.T) {
		assert.Equal(t, vmRem.Vban.InStream[0].GetOn(), false)
	})
}

func TestVbanOutStream6On(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Vban.OutStream[6].SetOn(true)
	t.Run("Should return true when SetOn(true)", func(t *testing.T) {
		assert.Equal(t, vmRem.Vban.OutStream[6].GetOn(), true)
	})

	vmRem.Vban.OutStream[6].SetOn(false)
	t.Run("Should return false when SetOn(false)", func(t *testing.T) {
		assert.Equal(t, vmRem.Vban.OutStream[6].GetOn(), false)
	})
}
