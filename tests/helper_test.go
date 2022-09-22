package voicemeeter_test

import (
	"log"
	"testing"

	"github.com/onyx-and-iris/voicemeeter"
)

var (
	vm, err = voicemeeter.NewRemote("potato", 30)
)

func TestMain(m *testing.M) {
	if err != nil {
		log.Fatal(err)
	}

	err = vm.Login()
	if err != nil {
		log.Fatal(err)
	}
	defer vm.Logout()

	m.Run()
}
