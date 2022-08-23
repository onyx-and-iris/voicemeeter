package voicemeeter

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"

	"golang.org/x/sys/windows/registry"
)

// dllPath returns the Voicemeeter installation path as a string
func dllPath() (string, error) {
	if runtime.GOOS != "windows" {
		return "", errors.New("only Windows OS supported")
	}

	var regkey string
	if strings.Contains(runtime.GOARCH, "64") {
		regkey = `SOFTWARE\WOW6432Node\Microsoft\Windows\CurrentVersion\Uninstall`
	} else {
		regkey = `SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall`
	}
	var vmkey = `\VB:Voicemeeter {17359A74-1236-5467}`

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, regkey+vmkey, registry.QUERY_VALUE)
	if err != nil {
		return "", errors.New("unable to access registry")
	}
	defer k.Close()

	path, _, err := k.GetStringValue(`UninstallString`)
	if err != nil {
		return "", errors.New("unable to read Voicemeeter path from registry")
	}

	var dllName string
	if strings.Contains(runtime.GOARCH, "64") {
		dllName = `VoicemeeterRemote64.dll`
	} else {
		dllName = `VoicemeeterRemote.dll`
	}
	return fmt.Sprintf("%v\\%s", filepath.Dir(path), dllName), nil
}

// getDllPath is a helper function for error handling
func getDllPath() string {
	path, err := dllPath()
	if err != nil {
		log.Fatal(err)
	}
	return path
}
