package voicemeeter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrip0Mute(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Strip[0].SetMute(true)
	sync()
	t.Run("Should return true when SetMute(true)", func(t *testing.T) {
		assert.True(t, vmRem.Strip[0].GetMute())
	})

	vmRem.Strip[0].SetMute(false)
	sync()
	t.Run("Should return false when SetMute(false)", func(t *testing.T) {
		assert.False(t, vmRem.Strip[0].GetMute())
	})
}

func TestStrip2Limit(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Strip[2].SetLimit(-8)
	sync()
	t.Run("Should return -8 when SetLimit(-8)", func(t *testing.T) {
		assert.Equal(t, vmRem.Strip[2].GetLimit(), -8)
	})

	vmRem.Strip[2].SetLimit(-32)
	sync()
	t.Run("Should return -32 when SetLimit(-8)", func(t *testing.T) {
		assert.Equal(t, vmRem.Strip[2].GetLimit(), -32)
	})

}

func TestStrip4Label(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Strip[4].SetLabel("test0")
	sync()
	t.Run("Should return test0 when SetLimit('test0')", func(t *testing.T) {
		assert.Equal(t, "test0", vmRem.Strip[4].GetLabel())
	})

	vmRem.Strip[4].SetLabel("test1")
	sync()
	t.Run("Should return test1 when SetLimit('test1')", func(t *testing.T) {
		assert.Equal(t, "test1", vmRem.Strip[4].GetLabel())
	})
}

func TestStrip5Gain(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Strip[4].SetGain(-20.8)
	sync()
	t.Run("Should return -20.8 when SetGain(-20.8)", func(t *testing.T) {
		assert.Equal(t, vmRem.Strip[4].GetGain(), -20.8)
	})

	vmRem.Strip[4].SetGain(-3.6)
	sync()
	t.Run("Should return -3.6 when SetGain(-3.6)", func(t *testing.T) {
		assert.Equal(t, vmRem.Strip[4].GetGain(), -3.6)
	})
}

func TestStrip3Comp(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Strip[4].SetComp(8.1)
	sync()
	t.Run("Should return 8.1 when SetGain(8.1)", func(t *testing.T) {
		assert.Equal(t, vmRem.Strip[4].GetComp(), 8.1)
	})

	vmRem.Strip[4].SetComp(1.6)
	sync()
	t.Run("Should return 1.6 when SetGain(1.6)", func(t *testing.T) {
		assert.Equal(t, vmRem.Strip[4].GetComp(), 1.6)
	})
}

func TestStrip5Mc(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Strip[5].SetMc(true)
	sync()
	t.Run("Should return true when SetMc(true)", func(t *testing.T) {
		assert.True(t, vmRem.Strip[5].GetMc())
	})

	vmRem.Strip[5].SetMc(false)
	sync()
	t.Run("Should return false when SetMc(false)", func(t *testing.T) {
		assert.False(t, vmRem.Strip[5].GetMc())
	})
}

func TestStrip2GainLayer3(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Strip[2].GainLayer()[3].Set(-18.3)
	sync()
	t.Run("Should return -18.3 when SetMc(true)", func(t *testing.T) {
		assert.Equal(t, vmRem.Strip[2].GainLayer()[3].Get(), -18.3)
	})

	vmRem.Strip[2].GainLayer()[3].Set(-25.6)
	sync()
	t.Run("Should return -25.6 when SetMc(true)", func(t *testing.T) {
		assert.Equal(t, vmRem.Strip[2].GainLayer()[3].Get(), -25.6)
	})
}

func TestBus3Eq(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Bus[3].SetEq(true)
	sync()
	t.Run("Should return true when SetEq(true)", func(t *testing.T) {
		assert.True(t, vmRem.Bus[3].GetEq())
	})

	vmRem.Bus[3].SetEq(false)
	sync()
	t.Run("Should return false when SetEq(false)", func(t *testing.T) {
		assert.False(t, vmRem.Bus[3].GetEq())
	})
}

func TestBus4Label(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Bus[4].SetLabel("test0")
	sync()
	t.Run("Should return test0 when SetEq('test0')", func(t *testing.T) {
		assert.Equal(t, "test0", vmRem.Bus[4].GetLabel())
	})

	vmRem.Bus[4].SetLabel("test1")
	sync()
	t.Run("Should return test1 when SetEq('test1')", func(t *testing.T) {
		assert.Equal(t, "test1", vmRem.Bus[4].GetLabel())
	})
}

func TestBus3ModeAmix(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Bus[3].Mode().SetAmix(true)
	sync()
	t.Run("Should return true when Mode().SetAmix(true)", func(t *testing.T) {
		assert.True(t, vmRem.Bus[3].Mode().GetAmix())
	})
}

func TestVbanInStream0On(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Vban.InStream[0].SetOn(true)
	sync()
	t.Run("Should return true when SetOn(true)", func(t *testing.T) {
		assert.True(t, vmRem.Vban.InStream[0].GetOn())
	})

	vmRem.Vban.InStream[0].SetOn(false)
	sync()
	t.Run("Should return false when SetOn(false)", func(t *testing.T) {
		assert.False(t, vmRem.Vban.InStream[0].GetOn())
	})
}

func TestVbanOutStream6On(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Vban.OutStream[6].SetOn(true)
	sync()
	t.Run("Should return true when SetOn(true)", func(t *testing.T) {
		assert.True(t, vmRem.Vban.OutStream[6].GetOn())
	})

	vmRem.Vban.OutStream[6].SetOn(false)
	sync()
	t.Run("Should return false when SetOn(false)", func(t *testing.T) {
		assert.False(t, vmRem.Vban.OutStream[6].GetOn())
	})
}

func TestVbanOutStream3Name(t *testing.T) {
	t.Skip("skipping test")
	vmRem.Vban.OutStream[3].SetName("test0")
	sync()
	t.Run("Should return test0 when SetName('test0')", func(t *testing.T) {
		assert.Equal(t, "test0", vmRem.Vban.OutStream[3].GetName())
	})

	vmRem.Vban.OutStream[3].SetName("test1")
	sync()
	t.Run("Should return test1 when SetName('test1')", func(t *testing.T) {
		assert.Equal(t, "test1", vmRem.Vban.OutStream[3].GetName())
	})
}

func TestVbanInStream4Bit(t *testing.T) {
	t.Skip("skipping test")
	t.Run("Should panic when instream SetBit(16)", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic")
			}
		}()
		vmRem.Vban.InStream[4].SetBit(16)
	})
}

func TestVbanOutStream4Bit(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Vban.OutStream[4].SetBit(16)
	sync()
	t.Run("Should return 16 when SetBit(16)", func(t *testing.T) {
		assert.Equal(t, vmRem.Vban.OutStream[4].GetBit(), 16)
	})

	vmRem.Vban.OutStream[4].SetBit(24)
	sync()
	t.Run("Should return 24 when SetBit(24)", func(t *testing.T) {
		assert.Equal(t, vmRem.Vban.OutStream[4].GetBit(), 24)
	})
}
