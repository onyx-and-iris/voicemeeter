package voicemeeter

type eQ struct {
	iRemote
}

func newEq(id string, i int) *eQ {
	return &eQ{iRemote{id, i}}
}

func (e *eQ) On() bool {
	return e.getter_bool("on")
}

func (e *eQ) SetOn(val bool) {
	e.setter_bool("on", val)
}

func (e *eQ) Ab() bool {
	return e.getter_bool("AB")
}

func (e *eQ) SetAb(val bool) {
	e.setter_bool("AB", val)
}
