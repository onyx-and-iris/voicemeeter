package voicemeeter_test

import (
	"bytes"
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/onyx-and-iris/voicemeeter/v2"
)

var (
	logstring bytes.Buffer
	vm, err   = voicemeeter.NewRemote("potato", 30)
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

	log.SetOutput(&logstring)

	m.Run()
}
