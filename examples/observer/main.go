package main

import (
	"fmt"
	"log"
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
	o.vm.Deregister(o)
}

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

func main() {
	vm, err := voicemeeter.NewRemote("potato")
	if err != nil {
		log.Fatal(err)
	}
	defer vm.Logout()

	vm.Login()
	// enable level updates (disabled by default)
	vm.EventAdd("ldirty")

	o := observer{vm}
	o.Register()
	time.Sleep(30 * time.Second)
	o.Deregister()
}
