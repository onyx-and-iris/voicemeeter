package voicemeeter_test

import (
	"log"
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
		log.Fatal(err)
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
