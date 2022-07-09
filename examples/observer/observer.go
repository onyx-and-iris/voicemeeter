package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/onyx-and-iris/voicemeeter-api-go"
)

type observer struct {
	vm *voicemeeter.Remote
}

func (o observer) OnUpdate(subject string) {
	if strings.Compare(subject, "pdirty") == 0 {
		fmt.Println("pdirty!")
	}
	if strings.Compare(subject, "mdirty") == 0 {
		fmt.Println("mdirty!")
	}
	if strings.Compare(subject, "ldirty") == 0 {
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
	vmRem := voicemeeter.GetRemote("potato")
	vmRem.Login()

	o := observer{vmRem}
	vmRem.Register(o)
	time.Sleep(30 * time.Second)
	vmRem.Deregister(o)

	vmRem.Logout()
}
