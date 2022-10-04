package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/onyx-and-iris/voicemeeter"
)

func main() {
	var (
		kind        string
		delay       int
		verbose     bool
		interactive bool
	)

	flag.StringVar(&kind, "kind", "banana", "kind of voicemeeter")
	flag.StringVar(&kind, "k", "banana", "kind of voicemeeter (shorthand)")
	flag.IntVar(&delay, "delay", 20, "delay between commands")
	flag.IntVar(&delay, "d", 20, "delay between commands (shorthand)")
	flag.BoolVar(&verbose, "verbose", false, "toggle console output")
	flag.BoolVar(&verbose, "v", false, "toggle console output (shorthand)")
	flag.BoolVar(&interactive, "interactive", false, "toggle interactive mode")
	flag.BoolVar(&interactive, "i", false, "toggle interactive mode (shorthand)")
	flag.Parse()

	vm, err := vmConnect(kind, delay)
	if err != nil {
		log.Fatal(err)
	}
	defer vm.Logout()

	err = runCommands(vm, verbose, interactive)
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

func runCommands(vm *voicemeeter.Remote, verbose, interactive bool) error {
	if interactive {
		return interactiveMode(vm, verbose)
	}
	for _, arg := range flag.Args() {
		err := parse(vm, arg, verbose)
		if err != nil {
			return err
		}
	}
	return nil
}

func interactiveMode(vm *voicemeeter.Remote, verbose bool) error {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "q" || input == "quit" || input == "" {
			return nil
		}
		for _, cmd := range strings.Split(input, " ") {
			err := parse(vm, cmd, verbose)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func parse(vm *voicemeeter.Remote, cmd string, verbose bool) error {
	if cmd[0] == '!' {
		err := toggleCmd(vm, cmd[1:], verbose)
		if err != nil {
			return err
		}
	} else {
		if strings.Contains(cmd, "=") {
			err := setCmd(vm, cmd, verbose)
			if err != nil {
				return err
			}
		} else {
			err := getCmd(vm, cmd, verbose)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func toggleCmd(vm *voicemeeter.Remote, cmd string, verbose bool) error {
	val, err := vm.GetFloat(cmd)
	if err != nil {
		err = fmt.Errorf("unable to toggle %s", cmd)
		return err
	}
	vm.SetFloat(cmd[1:], 1-val)
	if verbose {
		fmt.Println("Toggling", cmd)
	}
	return nil
}

func setCmd(vm *voicemeeter.Remote, cmd string, verbose bool) error {
	if verbose {
		fmt.Println("Running command", cmd)
	}
	err := vm.SendText(cmd)
	if err != nil {
		err = fmt.Errorf("unable to set %s", cmd)
		return err
	}
	return nil
}

func getCmd(vm *voicemeeter.Remote, cmd string, verbose bool) error {
	valF, err := vm.GetFloat(cmd)
	if err != nil {
		valS, err := vm.GetString(cmd)
		if err != nil {
			err = fmt.Errorf("unable to get %s", cmd)
			return err
		}
		if verbose {
			fmt.Println("Value of", cmd, "is:", valS)
		}
	} else {
		if verbose {
			fmt.Println("Value of", cmd, "is:", valF)
		}
	}
	return nil
}
