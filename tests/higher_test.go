package voicemeeter_test

import (
	"os"
	"testing"

	"github.com/onyx-and-iris/voicemeeter-api-go/voicemeeter"
)

var (
	vmRem = voicemeeter.NewRemote("banana")
)

func TestMain(m *testing.M) {
	vmRem.Login()
	code := m.Run()
	vmRem.Logout()
	os.Exit(code)
}

func TestStrip0Mute(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Strip[0].SetMute(true)
	if vmRem.Strip[0].GetMute() != true {
		t.Error("TestStrip0Mute did not match true")
	}
}

func TestStrip2Limit(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Strip[2].SetLimit(-8)
	if vmRem.Strip[2].GetLimit() != -8 {
		t.Error("TestStrip3Limit did not match -8")
	}
}

func TestStrip4Label(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Strip[4].SetLabel("test0")
	if vmRem.Strip[4].GetLabel() != "test0" {
		t.Error("TestStrip4Label did not match test0")
	}
}

func TestStrip5Gain(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Strip[4].SetGain(-20.8)
	if vmRem.Strip[4].GetGain() != -20.8 {
		t.Error("TestStrip5Gain did not match -20.8")
	}
}

func TestStrip3Mc(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Strip[3].SetMc(true)
	if vmRem.Strip[3].GetMc() != true {
		t.Error("TestStrip3Mc did not match true")
	}
}

func TestBus3Eq(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Bus[0].SetEq(true)
	if vmRem.Bus[0].GetEq() != true {
		t.Error("TestBus3Eq did not match true")
	}
}

func TestBus4Label(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Bus[4].SetLabel("test0")
	if vmRem.Bus[4].GetLabel() != "test0" {
		t.Error("TestBus4Label did not match test0")
	}
}

func TestVbanInStream0On(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Vban.InStream[0].SetOn(true)
	if vmRem.Vban.InStream[0].GetOn() != true {
		t.Error("TestVbanInStream0On did not match true")
	}
}

func TestVbanOutStream6On(t *testing.T) {
	//t.Skip("skipping test")
	vmRem.Vban.OutStream[6].SetOn(true)
	if vmRem.Vban.OutStream[6].GetOn() != true {
		t.Error("TestVbanOutStream6On did not match true")
	}
}
