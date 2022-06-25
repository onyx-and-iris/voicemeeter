package main

import (
	"fmt"
	"time"

	"github.com/onyx-and-iris/voicemeeter-api-go/voicemeeter"
)

type observer struct{}

func (o observer) OnUpdate(subject string) {
	fmt.Println(subject)
}

func main() {
	vmRem := voicemeeter.GetRemote("banana")
	vmRem.Login()

	o := observer{}
	vmRem.Pooler.Publisher.Register(o)

	time.Sleep(10 * time.Second)

	vmRem.Logout()
}
