package voicemeeter

type observer interface {
	OnUpdate(subject string)
}

type Publisher struct {
	observerList []observer
}

func (p *Publisher) Register(o observer) {
	p.observerList = append(p.observerList, o)
}

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

func (p *Publisher) notify(subject string) {
	for _, observer := range p.observerList {
		observer.OnUpdate(subject)
	}
}

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

func (r *Pooler) runner() {
	for r.run {
		if pdirty() {
			r.notify("pdirty")
		}
		if mdirty() {
			r.notify("mdirty")
		}
	}
}
