package voicemeeter_test

import (
	"os"
	"testing"

	"github.com/onyx-and-iris/voicemeeter-api-go/voicemeeter"
)

var (
	vmRem = voicemeeter.GetRemote("banana")
)

func TestMain(m *testing.M) {
	vmRem.Login()
	code := m.Run()
	vmRem.Logout()
	os.Exit(code)
}
