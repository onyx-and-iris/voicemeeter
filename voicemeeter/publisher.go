package voicemeeter

// observer defines the interface any registered observers must satisfy
type observer interface {
	OnUpdate(subject string)
}

// Publisher defines methods that support observers
type Publisher struct {
	observerList []observer
}

// Register adds an observer to observerList
func (p *Publisher) Register(o observer) {
	p.observerList = append(p.observerList, o)
}

// Deregister removes an observer from observerList
func (p *Publisher) Deregister(o observer) {
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
func (p *Publisher) notify(subject string) {
	for _, observer := range p.observerList {
		observer.OnUpdate(subject)
	}
}

// Pooler continuously polls the dirty paramters
// it is expected to be run in a goroutine
type Pooler struct {
	run bool
	Publisher
}

func newPooler() *Pooler {
	p := &Pooler{
		run: true,
	}
	go p.runner()
	return p
}

func (p *Pooler) runner() {
	for p.run {
		if pdirty() {
			p.notify("pdirty")
		}
		if mdirty() {
			p.notify("mdirty")
		}
	}
}
