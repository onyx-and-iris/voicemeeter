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

type (
	verbosePrinter struct {
		verbose bool
	}
)

func newVerbosePrinter() *verbosePrinter {
	return &verbosePrinter{}
}

func (v *verbosePrinter) printf(format string, a ...interface{}) {
	if v.verbose {
		fmt.Printf(format+"\n", a...)
	}
}

var (
	vPrinter *verbosePrinter
)

func init() {
	vPrinter = newVerbosePrinter()
}

func main() {
	var (
		kind        string
		delay       int
		interactive bool
	)

	flag.StringVar(&kind, "kind", "banana", "kind of voicemeeter")
	flag.StringVar(&kind, "k", "banana", "kind of voicemeeter (shorthand)")
	flag.IntVar(&delay, "delay", 20, "delay between commands")
	flag.IntVar(&delay, "d", 20, "delay between commands (shorthand)")
	flag.BoolVar(&vPrinter.verbose, "verbose", false, "toggle console output")
	flag.BoolVar(&vPrinter.verbose, "v", false, "toggle console output (shorthand)")
	flag.BoolVar(&interactive, "interactive", false, "toggle interactive mode")
	flag.BoolVar(&interactive, "i", false, "toggle interactive mode (shorthand)")
	flag.Parse()

	vm, err := vmConnect(kind, delay)
	if err != nil {
		log.Fatal(err)
	}
	defer vm.Logout()

	err = runCommands(vm, interactive)
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

func runCommands(vm *voicemeeter.Remote, interactive bool) error {
	if interactive {
		return interactiveMode(vm)
	}
	args := flag.Args()
	if len(args) == 0 {
		err := fmt.Errorf("must provide some commands to run")
		return err
	}
	for _, arg := range args {
		err := parse(vm, arg)
		if err != nil {
			vPrinter.printf(err.Error())
		}
	}
	return nil
}

func interactiveMode(vm *voicemeeter.Remote) error {
	vPrinter.printf("running in interactive mode... waiting for input")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "q" || input == "quit" || input == "" {
			return nil
		}
		for _, cmd := range strings.Split(input, " ") {
			err := parse(vm, cmd)
			if err != nil {
				vPrinter.printf(err.Error())
			}
		}
	}
	if scanner.Err() != nil {
		return scanner.Err()
	}
	return nil
}

func parse(vm *voicemeeter.Remote, cmd string) error {
	if cmd[0] == '!' {
		err := toggleCmd(vm, cmd[1:])
		if err != nil {
			return err
		}
	} else {
		if strings.Contains(cmd, "=") {
			err := setCmd(vm, cmd)
			if err != nil {
				return err
			}
		} else {
			err := getCmd(vm, cmd)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func toggleCmd(vm *voicemeeter.Remote, cmd string) error {
	val, err := vm.GetFloat(cmd)
	if err != nil {
		err = fmt.Errorf("unable to toggle %s", cmd)
		return err
	}
	vm.SetFloat(cmd, 1-val)
	vPrinter.printf("Toggling %s", cmd)
	return nil
}

func setCmd(vm *voicemeeter.Remote, cmd string) error {
	vPrinter.printf("Running command %s", cmd)
	err := vm.SendText(cmd)
	if err != nil {
		err = fmt.Errorf("unable to set %s", cmd)
		return err
	}
	return nil
}

func getCmd(vm *voicemeeter.Remote, cmd string) error {
	valF, err := vm.GetFloat(cmd)
	if err != nil {
		valS, err := vm.GetString(cmd)
		if err != nil {
			err = fmt.Errorf("unable to get %s", cmd)
			return err
		}
		vPrinter.printf("Value of %s is: %s", cmd, valS)
	} else {
		vPrinter.printf("Value of %s is: %v", cmd, valF)
	}
	return nil
}
