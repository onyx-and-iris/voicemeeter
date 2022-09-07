package main

import (
	"fmt"
	"log"
	"time"

	"github.com/onyx-and-iris/voicemeeter-api-go"

	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/events"
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

func main() {
	vm, err := voicemeeter.NewRemote("potato")
	if err != nil {
		log.Fatal(err)
	}

	err = vm.Login()
	if err != nil {
		log.Fatal(err)
	}
	defer vm.Logout()

	obs, err := goobs.New("localhost:4455", goobs.WithPassword("mystrongpass"))
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
