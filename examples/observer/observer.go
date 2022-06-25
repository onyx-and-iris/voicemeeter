package main

import (
	"fmt"
	"time"

	"github.com/onyx-and-iris/voicemeeter-api-go/voicemeeter"
)

type observer struct {
	i int
}

func (o observer) OnUpdate(subject string) {
	fmt.Println(o.i, subject)
}

func main() {
	vmRem := voicemeeter.GetRemote("banana")
	vmRem.Login()

	o := observer{1}
	o2 := observer{2}
	o3 := observer{3}
	o4 := observer{4}
	vmRem.Pooler.Register(o)
	vmRem.Pooler.Register(o2)
	vmRem.Pooler.Register(o3)
	vmRem.Pooler.Register(o4)

	time.Sleep(5 * time.Second)

	vmRem.Pooler.Deregister(o2)

	time.Sleep(5 * time.Second)

	vmRem.Logout()
}
