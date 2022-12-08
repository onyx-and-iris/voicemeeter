package voicemeeter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrip0Mute(t *testing.T) {
	//t.Skip("skipping test")
	vm.Strip[0].SetMute(true)
	t.Run("Should return true when Strip[0].SetMute(true)", func(t *testing.T) {
		assert.True(t, vm.Strip[0].Mute())
	})

	vm.Strip[0].SetMute(false)
	t.Run("Should return false when Strip[0].SetMute(false)", func(t *testing.T) {
		assert.False(t, vm.Strip[0].Mute())
	})
}

func TestStrip3A1(t *testing.T) {
	//t.Skip("skipping test")
	vm.Strip[3].SetA1(true)
	t.Run("Should return true when Strip[3].SetA1(true)", func(t *testing.T) {
		assert.True(t, vm.Strip[3].A1())
	})

	vm.Strip[3].SetA1(false)
	t.Run("Should return false when Strip[3].SetA1(false)", func(t *testing.T) {
		assert.False(t, vm.Strip[3].A1())
	})
}

func TestStrip2Limit(t *testing.T) {
	//t.Skip("skipping test")
	vm.Strip[2].SetLimit(-8)
	t.Run("Should return -8 when Strip[2].SetLimit(-8)", func(t *testing.T) {
		assert.Equal(t, vm.Strip[2].Limit(), -8)
	})

	vm.Strip[2].SetLimit(-32)
	t.Run("Should return -32 when Strip[2].SetLimit(-32)", func(t *testing.T) {
		assert.Equal(t, vm.Strip[2].Limit(), -32)
	})

}

func TestStrip4Label(t *testing.T) {
	//t.Skip("skipping test")
	vm.Strip[4].SetLabel("test0")
	t.Run("Should return test0 when Strip[4].SetLabel('test0')", func(t *testing.T) {
		assert.Equal(t, "test0", vm.Strip[4].Label())
	})

	vm.Strip[4].SetLabel("test1")
	t.Run("Should return test1 when Strip[4].SetLabel('test1')", func(t *testing.T) {
		assert.Equal(t, "test1", vm.Strip[4].Label())
	})
}

func TestStrip5Gain(t *testing.T) {
	//t.Skip("skipping test")
	vm.Strip[4].SetGain(-20.8)
	t.Run("Should return -20.8 when Strip[4].SetGain(-20.8)", func(t *testing.T) {
		assert.Equal(t, vm.Strip[4].Gain(), -20.8)
	})

	vm.Strip[4].SetGain(-3.6)
	t.Run("Should return -3.6 when Strip[4].SetGain(-3.6)", func(t *testing.T) {
		assert.Equal(t, vm.Strip[4].Gain(), -3.6)
	})
}

func TestStrip3CompKnob(t *testing.T) {
	//t.Skip("skipping test")
	vm.Strip[4].Comp().SetKnob(8.1)
	t.Run("Should return 8.1 when Strip[4].SetComp(8.1)", func(t *testing.T) {
		assert.Equal(t, vm.Strip[4].Comp().Knob(), 8.1)
	})

	vm.Strip[4].Comp().SetKnob(1.6)
	t.Run("Should return 1.6 when Strip[4].SetComp(1.6)", func(t *testing.T) {
		assert.Equal(t, vm.Strip[4].Comp().Knob(), 1.6)
	})
}

func TestStrip0CompGainIn(t *testing.T) {
	//t.Skip("skipping test")
	vm.Strip[0].Comp().SetGainIn(3.4)
	t.Run("Should return 3.4 when Strip[0].Comp().SetGainIn(3.4)", func(t *testing.T) {
		assert.Equal(t, vm.Strip[0].Comp().GainIn(), 3.4)
	})

	vm.Strip[0].Comp().SetGainIn(-19.3)
	t.Run("Should return -19.3 when Strip[0].Comp().SetGainIn(-19.3)", func(t *testing.T) {
		assert.Equal(t, vm.Strip[0].Comp().GainIn(), -19.3)
	})
}

func TestStrip3GateKnob(t *testing.T) {
	//t.Skip("skipping test")
	vm.Strip[4].Gate().SetKnob(8.1)
	t.Run("Should return 8.1 when Strip[4].SetComp(8.1)", func(t *testing.T) {
		assert.Equal(t, vm.Strip[4].Gate().Knob(), 8.1)
	})

	vm.Strip[4].Gate().SetKnob(1.6)
	t.Run("Should return 1.6 when Strip[4].SetComp(1.6)", func(t *testing.T) {
		assert.Equal(t, vm.Strip[4].Gate().Knob(), 1.6)
	})
}

func TestStrip0GateAttack(t *testing.T) {
	//t.Skip("skipping test")
	vm.Strip[0].Comp().SetAttack(3.4)
	t.Run("Should return 3.4 when Strip[0].Comp().SetAttack(3.4)", func(t *testing.T) {
		assert.Equal(t, vm.Strip[0].Comp().Attack(), 3.4)
	})

	vm.Strip[0].Comp().SetAttack(190.3)
	t.Run("Should return -19.3 when Strip[0].Comp().SetAttack(-19.3)", func(t *testing.T) {
		assert.Equal(t, vm.Strip[0].Comp().Attack(), 190.3)
	})
}

func TestStrip5Mc(t *testing.T) {
	//t.Skip("skipping test")
	vm.Strip[5].SetMc(true)
	t.Run("Should return true when Strip[5].SetMc(true)", func(t *testing.T) {
		assert.True(t, vm.Strip[5].Mc())
	})

	vm.Strip[5].SetMc(false)
	t.Run("Should return false when Strip[5].SetMc(false)", func(t *testing.T) {
		assert.False(t, vm.Strip[5].Mc())
	})
}

func TestStrip2GainLayer3(t *testing.T) {
	//t.Skip("skipping test")
	vm.Strip[2].GainLayer()[3].Set(-18.3)
	t.Run("Should return -18.3 when Strip[2].GainLayer()[3].Set(-18.3)", func(t *testing.T) {
		assert.Equal(t, vm.Strip[2].GainLayer()[3].Get(), -18.3)
	})

	vm.Strip[2].GainLayer()[3].Set(-25.6)
	t.Run("Should return -25.6 when Strip[2].GainLayer()[3].Set(-25.6)", func(t *testing.T) {
		assert.Equal(t, vm.Strip[2].GainLayer()[3].Get(), -25.6)
	})
}

func TestBus3EqOn(t *testing.T) {
	//t.Skip("skipping test")
	vm.Bus[3].Eq().SetOn(true)
	t.Run("Should return true when Bus[3].Eq().SetOn(true)", func(t *testing.T) {
		assert.True(t, vm.Bus[3].Eq().On())
	})

	vm.Bus[3].Eq().SetOn(false)
	t.Run("Should return false when Bus[3].SetEq(false)", func(t *testing.T) {
		assert.False(t, vm.Bus[3].Eq().On())
	})
}

func TestBus4Label(t *testing.T) {
	//t.Skip("skipping test")
	vm.Bus[4].SetLabel("test0")
	t.Run("Should return test0 when Bus[4].SetLabel('test0')", func(t *testing.T) {
		assert.Equal(t, "test0", vm.Bus[4].Label())
	})

	vm.Bus[4].SetLabel("test1")
	t.Run("Should return test1 when Bus[4].SetLabel('test1')", func(t *testing.T) {
		assert.Equal(t, "test1", vm.Bus[4].Label())
	})
}

func TestBus3ModeAmix(t *testing.T) {
	//t.Skip("skipping test")
	vm.Bus[3].Mode().SetAmix(true)
	t.Run("Should return true when Bus[3].Mode().SetAmix(true)", func(t *testing.T) {
		assert.True(t, vm.Bus[3].Mode().Amix())
	})
}

func TestVbanInStream0On(t *testing.T) {
	//t.Skip("skipping test")
	vm.Vban.InStream[0].SetOn(true)
	t.Run("Should return true when Vban.InStream[0].SetOn(true)", func(t *testing.T) {
		assert.True(t, vm.Vban.InStream[0].On())
	})

	vm.Vban.InStream[0].SetOn(false)
	t.Run("Should return false when Vban.InStream[0].SetOn(false)", func(t *testing.T) {
		assert.False(t, vm.Vban.InStream[0].On())
	})
}

func TestVbanOutStream6On(t *testing.T) {
	//t.Skip("skipping test")
	vm.Vban.OutStream[6].SetOn(true)
	t.Run("Should return true when Vban.OutStream[6].SetOn(true)", func(t *testing.T) {
		assert.True(t, vm.Vban.OutStream[6].On())
	})

	vm.Vban.OutStream[6].SetOn(false)
	t.Run("Should return false when Vban.OutStream[6].SetOn(false)", func(t *testing.T) {
		assert.False(t, vm.Vban.OutStream[6].On())
	})
}

func TestVbanOutStream3Name(t *testing.T) {
	//t.Skip("skipping test")
	vm.Vban.OutStream[3].SetName("test0")
	t.Run("Should return test0 when Vban.OutStream[3].SetName('test0')", func(t *testing.T) {
		assert.Equal(t, "test0", vm.Vban.OutStream[3].Name())
	})

	vm.Vban.OutStream[3].SetName("test1")
	t.Run("Should return test1 when Vban.OutStream[3].SetName('test1')", func(t *testing.T) {
		assert.Equal(t, "test1", vm.Vban.OutStream[3].Name())
	})
}

func TestVbanInStream4Bit(t *testing.T) {
	//t.Skip("skipping test")
	vm.Vban.InStream[4].SetBit(16)
	t.Run("Should log 'bit is readonly for vban instreams' when instream Vban.InStream[4].SetBit(16)", func(t *testing.T) {
		assert.Contains(t, logstring.String(), "bit is readonly for vban instreams")
	})
}

func TestVbanOutStream4Bit(t *testing.T) {
	//t.Skip("skipping test")
	vm.Vban.OutStream[4].SetBit(16)
	t.Run("Should return 16 when Vban.OutStream[4].SetBit(16)", func(t *testing.T) {
		assert.Equal(t, vm.Vban.OutStream[4].Bit(), 16)
	})

	vm.Vban.OutStream[4].SetBit(24)
	t.Run("Should return 24 when Vban.OutStream[4].SetBit(24)", func(t *testing.T) {
		assert.Equal(t, vm.Vban.OutStream[4].Bit(), 24)
	})
}
