package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/onyx-and-iris/voicemeeter-api-go"
)

func main() {
	kindId := flag.String("kind", "banana", "kind of voicemeeter")
	delay := flag.Int("delay", 15, "delay between commands")
	flag.Parse()

	vm, err := vmConnect(kindId, delay)
	if err != nil {
		log.Fatal(err)
	}
	defer vm.Logout()

	err = run_commands(vm)
	if err != nil {
		fmt.Println(err)
	}
}

func vmConnect(kindId *string, delay *int) (*voicemeeter.Remote, error) {
	vm, err := voicemeeter.NewRemote(*kindId, *delay)
	if err != nil {
		return nil, err
	}

	err = vm.Login()
	if err != nil {
		return nil, err
	}

	return vm, nil
}

func run_commands(vm *voicemeeter.Remote) error {
	for _, arg := range flag.Args() {
		if arg[0] == '!' {
			val, err := vm.GetFloat(arg[1:])
			if err != nil {
				err = fmt.Errorf("unable to toggle %s", arg[1:])
				return err
			}
			vm.SetFloat(arg[1:], 1-val)
			fmt.Println("Toggling", arg[1:])
		} else {
			if strings.Contains(arg, "=") {
				fmt.Println("Running command", arg)
				err := vm.SendText(arg)
				if err != nil {
					err = fmt.Errorf("unable to set %s", arg)
					return err
				}
			} else {
				val_f, err := vm.GetFloat(arg)
				if err != nil {
					val_s, err := vm.GetString(arg)
					if err != nil {
						err = fmt.Errorf("unable to get %s", arg)
						return err
					}
					fmt.Println("Value of", arg, "is:", val_s)
				} else {
					fmt.Println("Value of", arg, "is:", val_f)
				}
			}
		}
	}
	return nil
}
