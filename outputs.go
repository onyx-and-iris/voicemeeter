package voicemeeter

// iOutputs defines the interface outputs type must satisfy
type iOutputs interface {
	A1() bool
	SetA1(val bool)
	A2() bool
	SetA2(val bool)
	A3() bool
	SetA3(val bool)
	A4() bool
	SetA4(val bool)
	A5() bool
	SetA5(val bool)
	B1() bool
	SetB1(val bool)
	B2() bool
	SetB2(val bool)
	B3() bool
	SetB3(val bool)
}

// outputs represents the outputs field (A1 - A5, B1 - B3)
// expected to be embedded
type outputs struct {
	iRemote
}

// newOutputs returns an outputs type
func newOutputs(id string, i int) outputs {
	o := outputs{iRemote{id, i}}
	return o
}

// A1 returns the value of the A1 parameter
func (o *outputs) A1() bool {
	return o.getter_bool("A1")
}

// SetA1 sets the value of the A1 parameter
func (o *outputs) SetA1(val bool) {
	o.setter_bool("A1", val)
}

// A2 returns the value of the A2 parameter
func (o *outputs) A2() bool {
	return o.getter_bool("A2")
}

// SetA2 sets the value of the A2 parameter
func (o *outputs) SetA2(val bool) {
	o.setter_bool("A2", val)
}

// A3 returns the value of the A3 parameter
func (o *outputs) A3() bool {
	return o.getter_bool("A3")
}

// SetA3 sets the value of the A3 parameter
func (o *outputs) SetA3(val bool) {
	o.setter_bool("A3", val)
}

// A4 returns the value of the A4 parameter
func (o *outputs) A4() bool {
	return o.getter_bool("A4")
}

// SetA4 sets the value of the A4 parameter
func (o *outputs) SetA4(val bool) {
	o.setter_bool("A4", val)
}

// A5 returns the value of the A5 parameter
func (o *outputs) A5() bool {
	return o.getter_bool("A5")
}

// SetA5 sets the value of the A5 parameter
func (o *outputs) SetA5(val bool) {
	o.setter_bool("A5", val)
}

// B1 returns the value of the B1 parameter
func (o *outputs) B1() bool {
	return o.getter_bool("B1")
}

// SetB1 sets the value of the B1 parameter
func (o *outputs) SetB1(val bool) {
	o.setter_bool("B1", val)
}

// B2 returns the value of the B2 parameter
func (o *outputs) B2() bool {
	return o.getter_bool("B2")
}

// SetB2 sets the value of the B2 parameter
func (o *outputs) SetB2(val bool) {
	o.setter_bool("B2", val)
}

// B3 returns the value of the B3 parameter
func (o *outputs) B3() bool {
	return o.getter_bool("B3")
}

// SetB3 sets the value of the B3 parameter
func (o *outputs) SetB3(val bool) {
	o.setter_bool("B3", val)
}
