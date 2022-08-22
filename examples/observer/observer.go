package main

import (
	"fmt"
	"time"

	"github.com/onyx-and-iris/voicemeeter-api-go"
)

type observer struct {
	vm *voicemeeter.Remote
}

func (o observer) Register() {
	o.vm.Register(o)
}

func (o observer) Deregister() {
	o.vm.Register(o)
}

func (o observer) OnUpdate(subject string) {
	if subject == "pdirty" {
		fmt.Println("pdirty!")
	}
	if subject == "mdirty" {
		fmt.Println("mdirty!")
	}
	if subject == "midi" {
		var current = o.vm.Midi.Current()
		var val = o.vm.Midi.Get(current)
		fmt.Printf("Value of midi button %d: %d\n", current, val)
	}
	if subject == "ldirty" {
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

func main() {
	vm := voicemeeter.NewRemote("potato")
	vm.Login()
	// enable level updates (disabled by default)
	vm.EventAdd("ldirty")

	o := observer{vm}
	o.Register()
	time.Sleep(30 * time.Second)
	o.Deregister()

	vm.Logout()
}
