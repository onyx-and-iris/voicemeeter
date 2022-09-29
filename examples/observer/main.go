package main

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/onyx-and-iris/voicemeeter"
)

// observer represents a single receiver of updates
type observer struct {
	vm *voicemeeter.Remote
}

// newObserver add ldirty events to the eventlist and returns an observer type
func newObserver(vm *voicemeeter.Remote) *observer {
	vm.EventAdd("ldirty")
	return &observer{vm}
}

// register registers this observer to receive updates
func (o observer) register() {
	o.vm.Register(o)
}

// deregister deregisters this observer to receive updates
func (o observer) deregister() {
	o.vm.Deregister(o)
}

// OnUpdate satisfies the observer interface defined in publisher.go
// for each event type an action is triggered when the event occurs.
func (o observer) OnUpdate(subject string) {
	if subject == "pdirty" {
		fmt.Println("pdirty!")
	} else if subject == "mdirty" {
		fmt.Println("mdirty!")
	} else if subject == "midi" {
		var current = o.vm.Midi.Current()
		var val = o.vm.Midi.Get(current)
		fmt.Printf("Value of midi button %d: %d\n", current, val)
	} else if subject == "ldirty" {
		fmt.Printf("%v %v %v %v %v %v %v %v\n",
			o.vm.Bus[0].Levels().IsDirty(),
			o.vm.Bus[1].Levels().IsDirty(),
			o.vm.Bus[2].Levels().IsDirty(),
			o.vm.Bus[3].Levels().IsDirty(),
			o.vm.Bus[4].Levels().IsDirty(),
			o.vm.Bus[5].Levels().IsDirty(),
			o.vm.Bus[6].Levels().IsDirty(),
			o.vm.Bus[7].Levels().IsDirty(),
		)
	}
}

func init() {
	log.SetLevel(log.InfoLevel)
}

// main connects to Voiceemeter, registers observer for updates
// runs updates for 30 seconds and then deregisters observer.
func main() {
	vm, err := vmConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer vm.Logout()

	o := newObserver(vm)
	o.register()
	time.Sleep(30 * time.Second)
	o.deregister()
}

// vmConnect connects to Voicemeeter potato and logs into the API
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
