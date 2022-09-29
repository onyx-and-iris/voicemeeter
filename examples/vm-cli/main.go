package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/onyx-and-iris/voicemeeter"
)

var (
	kind    string
	delay   int
	verbose bool
)

func main() {
	flag.StringVar(&kind, "kind", "banana", "kind of voicemeeter")
	flag.StringVar(&kind, "k", "banana", "kind of voicemeeter (shorthand)")
	flag.IntVar(&delay, "delay", 20, "delay between commands")
	flag.IntVar(&delay, "d", 20, "delay between commands (shorthand)")
	flag.BoolVar(&verbose, "verbose", false, "toggle console output")
	flag.BoolVar(&verbose, "v", false, "toggle console output (shorthand)")
	flag.Parse()

	vm, err := vmConnect(kind, delay)
	if err != nil {
		log.Fatal(err)
	}
	defer vm.Logout()

	err = runCommands(vm, verbose)
	if err != nil {
		fmt.Println(err)
	}
}

func vmConnect(kind string, delay int) (*voicemeeter.Remote, error) {
	vm, err := voicemeeter.NewRemote(kind, delay)
	if err != nil {
		return nil, err
	}

	err = vm.Login()
	if err != nil {
		return nil, err
	}

	return vm, nil
}

func runCommands(vm *voicemeeter.Remote, verbose bool) error {
	for _, arg := range flag.Args() {
		if arg[0] == '!' {
			val, err := vm.GetFloat(arg[1:])
			if err != nil {
				err = fmt.Errorf("unable to toggle %s", arg[1:])
				return err
			}
			vm.SetFloat(arg[1:], 1-val)
			if verbose {
				fmt.Println("Toggling", arg[1:])
			}
		} else {
			if strings.Contains(arg, "=") {
				if verbose {
					fmt.Println("Running command", arg)
				}
				err := vm.SendText(arg)
				if err != nil {
					err = fmt.Errorf("unable to set %s", arg)
					return err
				}
			} else {
				valF, err := vm.GetFloat(arg)
				if err != nil {
					valS, err := vm.GetString(arg)
					if err != nil {
						err = fmt.Errorf("unable to get %s", arg)
						return err
					}
					if verbose {
						fmt.Println("Value of", arg, "is:", valS)
					}
				} else {
					if verbose {
						fmt.Println("Value of", arg, "is:", valF)
					}
				}
			}
		}
	}
	return nil
}
