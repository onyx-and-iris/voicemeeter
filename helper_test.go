package voicemeeter

import (
	"os"
	"testing"
	"time"
)

var (
	vm = NewRemote("potato")
)

func TestMain(m *testing.M) {
	vm.Login()
	code := m.Run()
	vm.Logout()
	os.Exit(code)
}

func sync() {
	time.Sleep(30 * time.Millisecond)
	for vm.Pdirty() || vm.Mdirty() {
	}
}
