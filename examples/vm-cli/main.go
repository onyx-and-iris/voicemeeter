package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/onyx-and-iris/voicemeeter/v2"
)

const (
	FLOAT = iota
	STRING
)

type result struct {
	kind        int
	stringParam string
	floatParam  float64
}

type verbosePrinter struct {
	verbose bool
}

func newVerbosePrinter() *verbosePrinter {
	return &verbosePrinter{}
}

func (v *verbosePrinter) printf(format string, a ...interface{}) {
	if v.verbose {
		fmt.Printf(format, a...)
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
		loglevel    int
		help        bool
	)

	flag.BoolVar(&help, "help", false, "print the help dialogue")
	flag.BoolVar(&help, "h", false, "print the help dialogue (shorthand)")
	flag.StringVar(&kind, "kind", "banana", "kind of voicemeeter")
	flag.StringVar(&kind, "k", "banana", "kind of voicemeeter (shorthand)")
	flag.IntVar(&delay, "delay", 20, "delay between commands")
	flag.IntVar(&delay, "d", 20, "delay between commands (shorthand)")
	flag.BoolVar(&interactive, "interactive", false, "toggle interactive mode")
	flag.BoolVar(&interactive, "i", false, "toggle interactive mode (shorthand)")
	flag.IntVar(&loglevel, "loglevel", int(log.WarnLevel), "set the log level")
	flag.IntVar(&loglevel, "l", int(log.WarnLevel), "set the log level (shorthand)")
	flag.BoolVar(&vPrinter.verbose, "verbose", false, "toggle console output")
	flag.BoolVar(&vPrinter.verbose, "v", false, "toggle console output (shorthand)")
	flag.Parse()

	if help {
		help_dialogue()
		return
	}

	if loglevel >= int(log.PanicLevel) && loglevel <= int(log.TraceLevel) {
		log.SetLevel(log.Level(loglevel))
	}

	vm, err := vmConnect(kind, delay)
	if err != nil {
		log.Fatal(err)
	}
	defer vm.Logout()

	if interactive {
		interactiveMode(vm)
		return
	}

	args := flag.Args()
	if len(args) == 0 {
		help_dialogue()
		return
	}

	for _, arg := range args {
		if err := parse(vm, arg); err != nil {
			log.Error(err.Error())
		}
	}
}

func help_dialogue() {
	fmt.Printf("usage: ./vm-cli [-h] [-i] [-k] [-l] [-d] [-v]\n" +
		"Where:\n" +
		"\th: Print the help dialogue\n" +
		"\ti: Enable interactive mode\n" +
		"\tk: The kind of Voicemeeter GUI to launch, defaults to Banana\n" +
		"\tl: Log level 0 up to 6, (defaults to 3, Warn Level)\n" +
		"\td: Set the delay between commands (defaults to 20ms)\n" +
		"\tv: Enable extra console output (toggle and set messages).")
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

func interactiveMode(vm *voicemeeter.Remote) error {
	fmt.Println("Interactive mode enabled. Enter 'Q' to exit.")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf(">> ")
	for scanner.Scan() {
		input := scanner.Text()
		if strings.ToUpper(input) == "Q" {
			break
		}

		for _, cmd := range strings.Split(input, " ") {
			if err := parse(vm, cmd); err != nil {
				log.Error(err.Error())
			}
		}
		fmt.Printf(">> ")
	}
	if scanner.Err() != nil {
		return scanner.Err()
	}
	return nil
}

func parse(vm *voicemeeter.Remote, cmd string) error {
	if cmd[0] == '!' {
		if err := toggleCmd(vm, cmd[1:]); err != nil {
			return err
		}
	} else if strings.Contains(cmd, "=") {
		if err := setCmd(vm, cmd); err != nil {
			return err
		}
	} else {
		r := result{kind: FLOAT}
		if err := getCmd(vm, cmd, &r); err != nil {
			return err
		}
		switch r.kind {
		case FLOAT:
			fmt.Printf("%s: %.2f\n", cmd, r.floatParam)
		case STRING:
			fmt.Printf("%s: %s\n", cmd, r.stringParam)
		}
	}
	return nil
}

func toggleCmd(vm *voicemeeter.Remote, cmd string) error {
	r := result{kind: FLOAT}
	if err := getCmd(vm, cmd, &r); err != nil {
		return err
	}
	if r.kind == FLOAT && (r.floatParam == 0 || r.floatParam == 1) {
		vPrinter.printf("Toggling %s\n", cmd)
		vm.SetFloat(cmd, 1-r.floatParam)
	} else {
		log.Warnf("%s does not appear to be a boolean parameter", cmd)
	}
	return nil
}

func setCmd(vm *voicemeeter.Remote, cmd string) error {
	if err := vm.SendText(cmd); err != nil {
		err = fmt.Errorf("unable to set %s", cmd)
		return err
	}
	vPrinter.printf("Setting %s\n", cmd)
	return nil
}

func getCmd(vm *voicemeeter.Remote, cmd string, r *result) error {
	if val, err := vm.GetFloat(cmd); err == nil {
		r.floatParam = val
	} else if val, err := vm.GetString(cmd); err == nil {
		r.kind = STRING
		r.stringParam = val
	} else {
		err := fmt.Errorf("unknown parameter '%s'", cmd)
		return err
	}
	return nil
}
