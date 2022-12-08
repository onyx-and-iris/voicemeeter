package voicemeeter

import "fmt"

// button represents a single macrobuttton
type button struct {
	index int
}

// newButton returns a button type
func newButton(i int) button {
	return button{i}
}

// getter returns the value of a macrobutton parameter
func (m *button) getter(mode int) bool {
	val, err := getMacroStatus(m.index, mode)
	if err != nil {
		fmt.Println(err)
	}
	return val == 1
}

// setter sets the value of a macrobutton parameter
func (m *button) setter(v bool, mode int) {
	var value int
	if v {
		value = 1
	} else {
		value = 0
	}
	setMacroStatus(m.index, value, mode)
}

// String implements the fmt.stringer interface
func (m *button) String() string {
	return fmt.Sprintf("MacroButton%d", m.index)
}

// State returns the value of the State parameter
func (m *button) State() bool {
	return m.getter(1)
}

// SetState sets the value of the State parameter
func (m *button) SetState(val bool) {
	m.setter(val, 1)
}

// StateOnly returns the value of the StateOnly parameter
func (m *button) StateOnly() bool {
	return m.getter(2)
}

// SetStateOnly sets the value of the StateOnly parameter
func (m *button) SetStateOnly(val bool) {
	m.setter(val, 2)
}

// Trigger returns the value of the Trigger parameter
func (m *button) Trigger() bool {
	return m.getter(2)
}

// SetTrigger returns the value of the Trigger parameter
func (m *button) SetTrigger(val bool) {
	m.setter(val, 2)
}
