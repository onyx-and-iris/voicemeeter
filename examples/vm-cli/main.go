package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/onyx-and-iris/voicemeeter-api-go"
)

func main() {
	vm, err := vmConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer vm.Logout()

	for _, arg := range os.Args[1:] {
		if arg[0] == '!' {
			vm.SetFloat(arg[1:], 1-vm.GetFloat(arg[1:]))
			fmt.Println("Toggling", arg[1:])
		} else {
			if strings.Contains(arg, "=") {
				fmt.Println("Running command", arg)
				vm.SendText(arg)
			} else {
				s := strings.Split(arg, ".")
				if strings.Contains(s[1], "label") {
					val := vm.GetString(arg)
					fmt.Println("Value of", arg, "is:", val)
				} else {
					val := vm.GetFloat(arg)
					fmt.Println("Value of", arg, "is:", val)
				}
			}
		}
	}
}

func vmConnect() (*voicemeeter.Remote, error) {
	vm, err := voicemeeter.NewRemote("banana", 15)
	if err != nil {
		return nil, err
	}

	err = vm.Login()
	if err != nil {
		return nil, err
	}

	return vm, nil
}
