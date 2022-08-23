package voicemeeter

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var (
	vm, err = NewRemote("potato")
)

func TestMain(m *testing.M) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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
