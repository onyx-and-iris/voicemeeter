package voicemeeter_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/onyx-and-iris/voicemeeter-api-go"
)

var (
	vm, err = voicemeeter.NewRemote("potato")
)

func TestMain(m *testing.M) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer vm.Logout()

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
