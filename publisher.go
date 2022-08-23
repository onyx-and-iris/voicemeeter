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

type event struct {
	pdirty bool
	mdirty bool
	midi   bool
	ldirty bool
}

func newEvent() *event {
	return &event{true, true, true, false}
}

func (e *event) Add(ev string) {
	switch ev {
	case "pdirty":
		e.pdirty = true
	case "mdirty":
		e.mdirty = true
	case "midi":
		e.midi = true
	case "ldirty":
		e.ldirty = true
	}
}

func (e *event) Remove(ev string) {
	switch ev {
	case "pdirty":
		e.pdirty = false
	case "mdirty":
		e.mdirty = false
	case "midi":
		e.midi = false
	case "ldirty":
		e.ldirty = false
	}
}

// pooler continuously polls the dirty paramters
// it is expected to be run in a goroutine
type pooler struct {
	k     *kind
	run   bool
	event *event
	publisher
}

func newPooler(k *kind) *pooler {
	p := &pooler{
		k:     k,
		run:   true,
		event: newEvent(),
	}
	go p.parameters()
	go p.macrobuttons()
	go p.midi()
	go p.levels()
	return p
}

func (p *pooler) parameters() {
	for p.run {
		if p.event.pdirty && pdirty() {
			p.notify("pdirty")
		}
		time.Sleep(33 * time.Millisecond)
	}
}

func (p *pooler) macrobuttons() {
	for p.run {
		if p.event.mdirty && mdirty() {
			p.notify("mdirty")
		}
		time.Sleep(33 * time.Millisecond)
	}
}

func (p *pooler) midi() {
	for p.run {
		if p.event.midi && getMidiMessage() {
			p.notify("midi")
		}
		time.Sleep(33 * time.Millisecond)
	}
}

func (p *pooler) levels() {
	_levelCache = newLevelCache(p.k)

	for p.run {
		if p.event.ldirty && ldirty(p.k) {
			update(_levelCache.stripLevels, _levelCache.stripLevelsBuff, (2*p.k.PhysIn)+(8*p.k.VirtIn))
			update(_levelCache.busLevels, _levelCache.busLevelsBuff, 8*p.k.NumBus())
			p.notify("ldirty")
		}
		time.Sleep(33 * time.Millisecond)
	}
}
