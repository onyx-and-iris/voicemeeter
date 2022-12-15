package voicemeeter

import (
	"time"

	log "github.com/sirupsen/logrus"
)

// publisher defines the list of observer channels
type publisher struct {
	observers []chan string
}

// Register adds an observer channel to the channelList
func (p *publisher) Register(channel chan string) {
	p.observers = append(p.observers, channel)
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

func (e *event) Add(events ...string) {
	for _, event := range events {
		switch event {
		case "pdirty":
			e.pdirty = true
		case "mdirty":
			e.mdirty = true
		case "midi":
			e.midi = true
		case "ldirty":
			e.ldirty = true
		}
		log.Info(event, " added to the pooler")
	}
}

func (e *event) Remove(events ...string) {
	for _, event := range events {
		switch event {
		case "pdirty":
			e.pdirty = false
		case "mdirty":
			e.mdirty = false
		case "midi":
			e.midi = false
		case "ldirty":
			e.ldirty = false
		}
		log.Info(event, " removed from the pooler")
	}
}

var p *pooler

// pooler continuously polls the dirty parameters
// it is expected to be run in a goroutine
type pooler struct {
	k          *kind
	run        bool
	event      *event
	pdirtyDone chan bool
	mdirtyDone chan bool
	midiDone   chan bool
	ldirtyDone chan bool
	publisher
}

func newPooler(k *kind) *pooler {
	p = &pooler{
		k:          k,
		run:        true,
		event:      newEvent(),
		pdirtyDone: make(chan bool),
		mdirtyDone: make(chan bool),
		midiDone:   make(chan bool),
		ldirtyDone: make(chan bool),
	}
	go p.done()
	go p.parameters()
	go p.macrobuttons()
	go p.midi()
	go p.levels()
	return p
}

func (p *pooler) done() {
	for {
		select {
		case _, ok := <-p.pdirtyDone:
			if !ok {
				p.pdirtyDone = nil
			}
		case _, ok := <-p.mdirtyDone:
			if !ok {
				p.mdirtyDone = nil
			}
		case _, ok := <-p.midiDone:
			if !ok {
				p.midiDone = nil
			}
		case _, ok := <-p.ldirtyDone:
			if !ok {
				p.ldirtyDone = nil
			}
		}
		if p.pdirtyDone == nil && p.mdirtyDone == nil && p.midiDone == nil && p.ldirtyDone == nil {
			for _, ch := range p.observers {
				close(ch)
			}
			break
		}
	}
}

func (p *pooler) parameters() {
	for p.run {
		pdirty, err := pdirty()
		if err != nil {
			close(p.pdirtyDone)
			break
		}
		if p.event.pdirty && pdirty {
			for _, ch := range p.observers {
				ch <- "pdirty"
			}
		}
		time.Sleep(33 * time.Millisecond)
	}
}

func (p *pooler) macrobuttons() {
	for p.run {
		mdirty, err := mdirty()
		if err != nil {
			close(p.mdirtyDone)
			break
		}
		if p.event.mdirty && mdirty {
			for _, ch := range p.observers {
				ch <- "mdirty"
			}
		}
		time.Sleep(33 * time.Millisecond)
	}
}

func (p *pooler) midi() {
	for p.run {
		midi, err := getMidiMessage()
		if err != nil {
			close(p.midiDone)
			break
		}
		if p.event.midi && midi {
			for _, ch := range p.observers {
				ch <- "midi"
			}
		}
		time.Sleep(33 * time.Millisecond)
	}
}

func (p *pooler) levels() {
	_levelCache = newLevelCache(p.k)

	for p.run {
		ldirty, err := ldirty(p.k)
		if err != nil {
			close(p.ldirtyDone)
			break
		}
		if p.event.ldirty && ldirty {
			update(_levelCache.stripLevels, _levelCache.stripLevelsBuff, (2*p.k.PhysIn)+(8*p.k.VirtIn))
			update(_levelCache.busLevels, _levelCache.busLevelsBuff, 8*p.k.NumBus())
			for _, ch := range p.observers {
				ch <- "ldirty"
			}
		}
		time.Sleep(33 * time.Millisecond)
	}
}
