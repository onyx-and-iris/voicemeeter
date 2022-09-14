package voicemeeter

import (
	"log"
	"os"
	"testing"
)

var (
	vm, err = NewRemote("potato", 30)
)

func TestMain(m *testing.M) {
	if err != nil {
		log.Fatal(err)
	}

	vm.Login()
	code := m.Run()
	vm.Logout()
	os.Exit(code)
}
