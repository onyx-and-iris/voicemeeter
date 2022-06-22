package voicemeeter

import (
	"fmt"
)

// custom bus type, struct forwarding channel
type bus struct {
	channel
}

// newBus returns a strip type
// it also initializes embedded channel type
func newBus(i int, k *kind) bus {
	return bus{channel{"bus", i, *k}}
}

// String implements the stringer interface
func (b *bus) String() string {
	if b.index < b.kind.physOut {
		return fmt.Sprintf("PhysicalBus%d\n", b.index)
	}
	return fmt.Sprintf("VirtualBus%d\n", b.index)
}

// GetMute returns the value of the Mute parameter
func (b *bus) GetMute() bool {
	return b.getter_bool("Mute")
}

// SetMute sets the value of the Mute parameter
func (b *bus) SetMute(val bool) {
	b.setter_bool("Mute", val)
}

// GetEq returns the value of the Eq.On parameter
func (b *bus) GetEq() bool {
	return b.getter_bool("Eq.On")
}

// SetEq sets the value of the Eq.On parameter
func (b *bus) SetEq(val bool) {
	b.setter_bool("Eq.On", val)
}
