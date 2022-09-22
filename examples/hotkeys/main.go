package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
	"github.com/onyx-and-iris/voicemeeter"
)

func main() {
	if err := keyboard.Open(); err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	vm, err := vmConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer vm.Logout()

	fmt.Println("Press ESC to quit")
Loop:
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch char {
		case '0':
			fmt.Printf("Logged into Voicemeeter %s, version %s\n", vm.Type(), vm.Version())
		case '1':
			vm.Strip[0].SetMute(!vm.Strip[0].GetMute())
		case '2':
			if vm.Strip[3].GetGain() == -12.8 {
				vm.Strip[3].FadeBy(-8.3, 500)
			} else {
				vm.Strip[3].FadeTo(-12.8, 500)
			}
		case '3':
			vm.Strip[5].AppMute("Spotify", true)
		default:
			if key == keyboard.KeyEsc {
				break Loop
			}
		}
	}
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
