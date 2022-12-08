package main

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/onyx-and-iris/voicemeeter/v2"

	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/events"

	"github.com/BurntSushi/toml"
)

func onStart(vm *voicemeeter.Remote) {
	vm.Strip[0].SetMute(true)
	vm.Strip[1].SetB1(true)
	vm.Strip[2].SetB1(true)
}

func onBrb(vm *voicemeeter.Remote) {
	vm.Strip[7].FadeTo(0, 500)
	vm.Bus[0].SetMute(true)
}

func onLive(vm *voicemeeter.Remote) {
	vm.Strip[0].SetMute(false)
	vm.Strip[7].FadeTo(-6, 500)
	vm.Strip[7].SetA3(true)
	vm.Vban.InStream[0].SetOn(true)
}

func onEnd(vm *voicemeeter.Remote) {
	vm.Strip[0].SetMute(true)
	vm.Strip[1].SetMute(true)
	vm.Strip[1].SetB1(false)
	vm.Strip[2].SetMute(true)
	vm.Strip[2].SetB1(false)
	vm.Vban.InStream[0].SetOn(false)
}

func init() {
	log.SetLevel(log.InfoLevel)
}

func main() {
	vm, err := vmConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer vm.Logout()

	obs, err := obsConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer obs.Disconnect()

	version, _ := obs.General.GetVersion()
	fmt.Printf("OBS Studio version: %s\n", version.ObsVersion)
	fmt.Printf("Websocket server version: %s\n", version.ObsWebSocketVersion)

	go obs.Listen(func(event any) {
		switch e := event.(type) {
		case *events.CurrentProgramSceneChanged:
			fmt.Printf("Switched to scene %s\n", e.SceneName)
			switch e.SceneName {
			case "START":
				onStart(vm)
			case "BRB":
				onBrb(vm)
			case "LIVE":
				onLive(vm)
			case "END":
				onEnd(vm)
			}
		}
	})

	time.Sleep(30 * time.Second)
}

func vmConnect() (*voicemeeter.Remote, error) {
	vm, err := voicemeeter.NewRemote("potato", 0)
	if err != nil {
		return nil, err
	}

	err = vm.Login()
	if err != nil {
		return nil, err
	}

	return vm, nil
}

func obsConnect() (*goobs.Client, error) {
	type (
		connection struct {
			Host     string
			Port     int
			Password string
		}

		config struct {
			Connection map[string]connection
		}
	)

	f := "config.toml"
	if _, err := os.Stat(f); err != nil {
		err := fmt.Errorf("unable to locate %s", f)
		return nil, err
	}

	var c config
	_, err := toml.DecodeFile(f, &c.Connection)
	if err != nil {
		return nil, err
	}
	conn := c.Connection["connection"]

	obs, err := goobs.New(fmt.Sprintf("%s:%d", conn.Host, conn.Port), goobs.WithPassword(conn.Password))
	if err != nil {
		return nil, err
	}
	return obs, nil
}
