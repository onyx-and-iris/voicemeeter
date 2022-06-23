package voicemeeter

import (
	"fmt"
)

type t_bus interface {
	GetMute() bool
	SetMute(val bool)
	GetEq() bool
	SetEq(val bool)
	GetMono() bool
	SetMono(val bool)
	GetLabel() string
	SetLabel(val string)
	GetGain() float64
	SetGain(val float32)
}

// bus represents a bus channel
// embeds channel struct
type bus struct {
	channel
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

// GetMono returns the value of the Mute parameter
func (b *bus) GetMono() bool {
	return b.getter_bool("Mono")
}

// SetMono sets the value of the Mute parameter
func (b *bus) SetMono(val bool) {
	b.setter_bool("Mono", val)
}

// GetLabel returns the value of the MC parameter
func (b *bus) GetLabel() string {
	return b.getter_string("Label")
}

// SetLabel sets the value of the MC parameter
func (b *bus) SetLabel(val string) {
	b.setter_string("Label", val)
}

// GetGain returns the value of the Gain parameter
func (b *bus) GetGain() float64 {
	return b.getter_float("Gain")
}

// SetGain sets the value of the Gain parameter
func (b *bus) SetGain(val float32) {
	b.setter_float("Gain", val)
}

type physicalBus struct {
	bus
}

func newPhysicalBus(i int, k *kind) t_bus {
	pb := physicalBus{bus{channel{"bus", i, *k}}}
	return t_bus(&pb)
}

type virtualBus struct {
	bus
}

func newVirtualBus(i int, k *kind) t_bus {
	vb := virtualBus{bus{channel{"bus", i, *k}}}
	return t_bus(&vb)
}
