package voicemeeter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrip0Mute(t *testing.T) {
	//t.Skip("skipping test")
	vm.Strip[0].SetMute(true)
	sync()
	t.Run("Should return true when SetMute(true)", func(t *testing.T) {
		assert.True(t, vm.Strip[0].GetMute())
	})

	vm.Strip[0].SetMute(false)
	sync()
	t.Run("Should return false when SetMute(false)", func(t *testing.T) {
		assert.False(t, vm.Strip[0].GetMute())
	})
}

func TestStrip3A1(t *testing.T) {
	//t.Skip("skipping test")
	vm.Strip[3].SetA1(true)
	sync()
	t.Run("Should return true when SetA1(true)", func(t *testing.T) {
		assert.True(t, vm.Strip[3].GetA1())
	})

	vm.Strip[3].SetA1(false)
	sync()
	t.Run("Should return false when SetA1(false)", func(t *testing.T) {
		assert.False(t, vm.Strip[3].GetA1())
	})
}

func TestStrip2Limit(t *testing.T) {
	//t.Skip("skipping test")
	vm.Strip[2].SetLimit(-8)
	sync()
	t.Run("Should return -8 when SetLimit(-8)", func(t *testing.T) {
		assert.Equal(t, vm.Strip[2].GetLimit(), -8)
	})

	vm.Strip[2].SetLimit(-32)
	sync()
	t.Run("Should return -32 when SetLimit(-8)", func(t *testing.T) {
		assert.Equal(t, vm.Strip[2].GetLimit(), -32)
	})

}

func TestStrip4Label(t *testing.T) {
	//t.Skip("skipping test")
	vm.Strip[4].SetLabel("test0")
	sync()
	t.Run("Should return test0 when SetLimit('test0')", func(t *testing.T) {
		assert.Equal(t, "test0", vm.Strip[4].GetLabel())
	})

	vm.Strip[4].SetLabel("test1")
	sync()
	t.Run("Should return test1 when SetLimit('test1')", func(t *testing.T) {
		assert.Equal(t, "test1", vm.Strip[4].GetLabel())
	})
}

func TestStrip5Gain(t *testing.T) {
	//t.Skip("skipping test")
	vm.Strip[4].SetGain(-20.8)
	sync()
	t.Run("Should return -20.8 when SetGain(-20.8)", func(t *testing.T) {
		assert.Equal(t, vm.Strip[4].GetGain(), -20.8)
	})

	vm.Strip[4].SetGain(-3.6)
	sync()
	t.Run("Should return -3.6 when SetGain(-3.6)", func(t *testing.T) {
		assert.Equal(t, vm.Strip[4].GetGain(), -3.6)
	})
}

func TestStrip3Comp(t *testing.T) {
	//t.Skip("skipping test")
	vm.Strip[4].SetComp(8.1)
	sync()
	t.Run("Should return 8.1 when SetGain(8.1)", func(t *testing.T) {
		assert.Equal(t, vm.Strip[4].GetComp(), 8.1)
	})

	vm.Strip[4].SetComp(1.6)
	sync()
	t.Run("Should return 1.6 when SetGain(1.6)", func(t *testing.T) {
		assert.Equal(t, vm.Strip[4].GetComp(), 1.6)
	})
}

func TestStrip5Mc(t *testing.T) {
	//t.Skip("skipping test")
	vm.Strip[5].SetMc(true)
	sync()
	t.Run("Should return true when SetMc(true)", func(t *testing.T) {
		assert.True(t, vm.Strip[5].GetMc())
	})

	vm.Strip[5].SetMc(false)
	sync()
	t.Run("Should return false when SetMc(false)", func(t *testing.T) {
		assert.False(t, vm.Strip[5].GetMc())
	})
}

func TestStrip2GainLayer3(t *testing.T) {
	//t.Skip("skipping test")
	vm.Strip[2].GainLayer()[3].Set(-18.3)
	sync()
	t.Run("Should return -18.3 when SetMc(true)", func(t *testing.T) {
		assert.Equal(t, vm.Strip[2].GainLayer()[3].Get(), -18.3)
	})

	vm.Strip[2].GainLayer()[3].Set(-25.6)
	sync()
	t.Run("Should return -25.6 when SetMc(true)", func(t *testing.T) {
		assert.Equal(t, vm.Strip[2].GainLayer()[3].Get(), -25.6)
	})
}

func TestBus3Eq(t *testing.T) {
	//t.Skip("skipping test")
	vm.Bus[3].SetEq(true)
	sync()
	t.Run("Should return true when SetEq(true)", func(t *testing.T) {
		assert.True(t, vm.Bus[3].GetEq())
	})

	vm.Bus[3].SetEq(false)
	sync()
	t.Run("Should return false when SetEq(false)", func(t *testing.T) {
		assert.False(t, vm.Bus[3].GetEq())
	})
}

func TestBus4Label(t *testing.T) {
	//t.Skip("skipping test")
	vm.Bus[4].SetLabel("test0")
	sync()
	t.Run("Should return test0 when SetEq('test0')", func(t *testing.T) {
		assert.Equal(t, "test0", vm.Bus[4].GetLabel())
	})

	vm.Bus[4].SetLabel("test1")
	sync()
	t.Run("Should return test1 when SetEq('test1')", func(t *testing.T) {
		assert.Equal(t, "test1", vm.Bus[4].GetLabel())
	})
}

func TestBus3ModeAmix(t *testing.T) {
	//t.Skip("skipping test")
	vm.Bus[3].Mode().SetAmix(true)
	sync()
	t.Run("Should return true when Mode().SetAmix(true)", func(t *testing.T) {
		assert.True(t, vm.Bus[3].Mode().GetAmix())
	})
}

func TestVbanInStream0On(t *testing.T) {
	//t.Skip("skipping test")
	vm.Vban.InStream[0].SetOn(true)
	sync()
	t.Run("Should return true when SetOn(true)", func(t *testing.T) {
		assert.True(t, vm.Vban.InStream[0].GetOn())
	})

	vm.Vban.InStream[0].SetOn(false)
	sync()
	t.Run("Should return false when SetOn(false)", func(t *testing.T) {
		assert.False(t, vm.Vban.InStream[0].GetOn())
	})
}

func TestVbanOutStream6On(t *testing.T) {
	//t.Skip("skipping test")
	vm.Vban.OutStream[6].SetOn(true)
	sync()
	t.Run("Should return true when SetOn(true)", func(t *testing.T) {
		assert.True(t, vm.Vban.OutStream[6].GetOn())
	})

	vm.Vban.OutStream[6].SetOn(false)
	sync()
	t.Run("Should return false when SetOn(false)", func(t *testing.T) {
		assert.False(t, vm.Vban.OutStream[6].GetOn())
	})
}

func TestVbanOutStream3Name(t *testing.T) {
	t.Skip("skipping test")
	vm.Vban.OutStream[3].SetName("test0")
	sync()
	t.Run("Should return test0 when SetName('test0')", func(t *testing.T) {
		assert.Equal(t, "test0", vm.Vban.OutStream[3].GetName())
	})

	vm.Vban.OutStream[3].SetName("test1")
	sync()
	t.Run("Should return test1 when SetName('test1')", func(t *testing.T) {
		assert.Equal(t, "test1", vm.Vban.OutStream[3].GetName())
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
		vm.Vban.InStream[4].SetBit(16)
	})
}

func TestVbanOutStream4Bit(t *testing.T) {
	//t.Skip("skipping test")
	vm.Vban.OutStream[4].SetBit(16)
	sync()
	t.Run("Should return 16 when SetBit(16)", func(t *testing.T) {
		assert.Equal(t, vm.Vban.OutStream[4].GetBit(), 16)
	})

	vm.Vban.OutStream[4].SetBit(24)
	sync()
	t.Run("Should return 24 when SetBit(24)", func(t *testing.T) {
		assert.Equal(t, vm.Vban.OutStream[4].GetBit(), 24)
	})
}
