package voicemeeter

import "fmt"

type t_outputs interface {
	GetA1() bool
	SetA1(val bool)
	GetA2() bool
	SetA2(val bool)
	GetA3() bool
	SetA3(val bool)
	GetA4() bool
	SetA4(val bool)
	GetA5() bool
	SetA5(val bool)
	GetB1() bool
	SetB1(val bool)
	GetB2() bool
	SetB2(val bool)
	GetB3() bool
	SetB3(val bool)
}

type outputs struct {
	iRemote
}

func newOutputs(i int) outputs {
	o := outputs{iRemote{fmt.Sprintf("strip[%d]", i), i}}
	return o
}

// GetA1 returns the value of the A1 parameter
func (o *outputs) GetA1() bool {
	return o.getter_bool("A1")
}

// SetA1 sets the value of the A1 parameter
func (o *outputs) SetA1(val bool) {
	o.setter_bool("A1", val)
}

// GetA2 returns the value of the A2 parameter
func (o *outputs) GetA2() bool {
	return o.getter_bool("A2")
}

// SetA2 sets the value of the A2 parameter
func (o *outputs) SetA2(val bool) {
	o.setter_bool("A2", val)
}

// GetA3 returns the value of the A3 parameter
func (o *outputs) GetA3() bool {
	return o.getter_bool("A3")
}

// SetA3 sets the value of the A3 parameter
func (o *outputs) SetA3(val bool) {
	o.setter_bool("A3", val)
}

// GetA4 returns the value of the A4 parameter
func (o *outputs) GetA4() bool {
	return o.getter_bool("A4")
}

// SetA4 sets the value of the A4 parameter
func (o *outputs) SetA4(val bool) {
	o.setter_bool("A4", val)
}

// GetA5 returns the value of the A5 parameter
func (o *outputs) GetA5() bool {
	return o.getter_bool("A5")
}

// SetA5 sets the value of the A5 parameter
func (o *outputs) SetA5(val bool) {
	o.setter_bool("A5", val)
}

// GetB1 returns the value of the B1 parameter
func (o *outputs) GetB1() bool {
	return o.getter_bool("B1")
}

// SetB1 sets the value of the B1 parameter
func (o *outputs) SetB1(val bool) {
	o.setter_bool("B1", val)
}

// GetB2 returns the value of the B2 parameter
func (o *outputs) GetB2() bool {
	return o.getter_bool("B2")
}

// SetB2 sets the value of the B2 parameter
func (o *outputs) SetB2(val bool) {
	o.setter_bool("B2", val)
}

// GetB3 returns the value of the B3 parameter
func (o *outputs) GetB3() bool {
	return o.getter_bool("B3")
}

// SetB3 sets the value of the B3 parameter
func (o *outputs) SetB3(val bool) {
	o.setter_bool("B3", val)
}
