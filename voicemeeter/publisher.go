package voicemeeter

import (
	"time"
)

// observer defines the interface any registered observers must satisfy
type observer interface {
	OnUpdate(subject string)
}

// publisher defines methods that support observers
type publisher struct {
	observerList []observer
}

// Register adds an observer to observerList
func (p *publisher) Register(o observer) {
	p.observerList = append(p.observerList, o)
}

// Deregister removes an observer from observerList
func (p *publisher) Deregister(o observer) {
	var indexToRemove int

	for i, observer := range p.observerList {
		if observer == o {
			indexToRemove = i
			break
		}
	}

	p.observerList = append(p.observerList[:indexToRemove], p.observerList[indexToRemove+1:]...)
}

// notify updates observers of any changes
func (p *publisher) notify(subject string) {
	for _, observer := range p.observerList {
		observer.OnUpdate(subject)
	}
}

// pooler continuously polls the dirty paramters
// it is expected to be run in a goroutine
type pooler struct {
	run bool
	publisher
}

func newPooler() *pooler {
	p := &pooler{
		run: true,
	}
	go p.runner()
	return p
}

func (p *pooler) runner() {
	for p.run {
		if pdirty() {
			p.notify("pdirty")
		}
		if mdirty() {
			p.notify("mdirty")
		}
		time.Sleep(33 * time.Millisecond)
	}
}
