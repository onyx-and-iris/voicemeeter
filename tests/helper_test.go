package voicemeeter_test

import (
	"os"
	"testing"
	"time"

	"github.com/onyx-and-iris/voicemeeter-api-go"
)

var (
	vm = voicemeeter.NewRemote("potato")
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
