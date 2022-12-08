package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/onyx-and-iris/voicemeeter/v2"
)

// observer represents a single receiver of updates
type observer struct {
	vm     *voicemeeter.Remote
	events chan string
}

// newObserver returns an observer type
func newObserver(vm *voicemeeter.Remote) *observer {
	o := &observer{vm, make(chan string)}
	return o
}

// OnUpdate satisfies the observer interface defined in publisher.go
// for each event type an action is triggered when the event occurs.
func (o observer) Listen() {
	o.vm.Register(o.events)

	for s := range o.events {
		switch s {
		case "pdirty", "mdirty":
			fmt.Println(s)
		case "midi":
			var current = o.vm.Midi.Current()
			var val = o.vm.Midi.Get(current)
			fmt.Printf("Value of midi button %d: %d\n", current, val)
		case "ldirty":
			for _, bus := range o.vm.Bus {
				if bus.Levels().IsDirty() {
					fmt.Println(bus, bus.Levels().All())
				}
			}
		}
	}
}

func init() {
	log.SetLevel(log.InfoLevel)
}

func runObserver(vm *voicemeeter.Remote) {
	o := newObserver(vm)
	o.Listen()
}

// main connects to Voiceemeter, registers observer for updates
// runs updates for 30 seconds and then deregisters observer.
func main() {
	vm, err := vmConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer vm.Logout()

	runObserver(vm)
}

// vmConnect connects to Voicemeeter potato and logs into the API
func vmConnect() (*voicemeeter.Remote, error) {
	vm, err := voicemeeter.NewRemote("basic", 0)
	if err != nil {
		return nil, err
	}

	err = vm.Login()
	if err != nil {
		return nil, err
	}
	vm.EventAdd("ldirty")

	return vm, nil
}
