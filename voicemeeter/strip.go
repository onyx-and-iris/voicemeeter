package voicemeeter

import (
	"fmt"
)

type t_strip interface {
	GetMute() bool
	SetMute(val bool)
	GetMono() bool
	SetMono(val bool)
	GetSolo() bool
	SetSolo(val bool)
	GetLimit() int
	SetLimit(val int)
	GetLabel() string
	SetLabel(val string)
	GetGain() float64
	SetGain(val float32)
	GetMc() bool
	SetMc(val bool)
	GetComp() bool
	SetComp(val bool)
	GetGate() bool
	SetGate(val bool)
	GetAudibility() bool
	SetAudibility(val bool)
}

// strip represents a strip channel
// embeds channel struct
type strip struct {
	iRemote
}

// implement stringer interface in fmt
func (s *strip) String() string {
	if s.index < s.kind.physIn {
		return fmt.Sprintf("PhysicalStrip%d\n", s.index)
	}
	return fmt.Sprintf("VirtualStrip%d\n", s.index)
}

// GetMute returns the value of the Mute parameter
func (s *strip) GetMute() bool {
	return s.getter_bool("Mute")
}

// SetMute sets the value of the Mute parameter
func (s *strip) SetMute(val bool) {
	s.setter_bool("Mute", val)
}

// GetMono returns the value of the Mute parameter
func (s *strip) GetMono() bool {
	return s.getter_bool("Mono")
}

// SetMono sets the value of the Mute parameter
func (s *strip) SetMono(val bool) {
	s.setter_bool("Mono", val)
}

// GetSolo returns the value of the Mute parameter
func (s *strip) GetSolo() bool {
	return s.getter_bool("Solo")
}

// SetSolo sets the value of the Mute parameter
func (s *strip) SetSolo(val bool) {
	s.setter_bool("Solo", val)
}

// GetLimit returns the value of the Limit parameter
func (s *strip) GetLimit() int {
	return s.getter_int("Limit")
}

// SetLimit sets the value of the Limit parameter
func (s *strip) SetLimit(val int) {
	s.setter_int("Limit", val)
}

// GetLabel returns the value of the MC parameter
func (s *strip) GetLabel() string {
	return s.getter_string("Label")
}

// SetLabel sets the value of the MC parameter
func (s *strip) SetLabel(val string) {
	s.setter_string("Label", val)
}

// GetGain returns the value of the Gain parameter
func (s *strip) GetGain() float64 {
	return s.getter_float("Gain")
}

// SetGain sets the value of the Gain parameter
func (s *strip) SetGain(val float32) {
	s.setter_float("Gain", val)
}

type physicalStrip struct {
	strip
}

func newPhysicalStrip(i int, k *kind) t_strip {
	ps := physicalStrip{strip{iRemote{"strip", i, k}}}
	return t_strip(&ps)
}

// GetComp returns the value of the Mute parameter
func (p *physicalStrip) GetComp() bool {
	return p.getter_bool("Comp")
}

// SetComp sets the value of the Mute parameter
func (p *physicalStrip) SetComp(val bool) {
	p.setter_bool("Comp", val)
}

// GetGate returns the value of the Mute parameter
func (p *physicalStrip) GetGate() bool {
	return p.getter_bool("Gate")
}

// SetGate sets the value of the Mute parameter
func (p *physicalStrip) SetGate(val bool) {
	p.setter_bool("Gate", val)
}

// GetAudibility returns the value of the Mute parameter
func (p *physicalStrip) GetAudibility() bool {
	return p.getter_bool("Audibility")
}

// SetAudibility sets the value of the Mute parameter
func (p *physicalStrip) SetAudibility(val bool) {
	p.setter_bool("Audibility", val)
}

// GetMc panics reason invalid parameter
func (p *physicalStrip) GetMc() bool {
	panic("invalid parameter MC for physicalStrip")
}

// SetMc panics reason invalid parameter
func (p *physicalStrip) SetMc(val bool) {
	panic("invalid parameter MC for physicalStrip")
}

type virtualStrip struct {
	strip
}

func newVirtualStrip(i int, k *kind) t_strip {
	vs := virtualStrip{strip{iRemote{"strip", i, k}}}
	return t_strip(&vs)
}

// GetMc returns the value of the MC parameter
func (v *virtualStrip) GetMc() bool {
	return v.getter_bool("MC")
}

// SetMc sets the value of the MC parameter
func (v *virtualStrip) SetMc(val bool) {
	v.setter_bool("MC", val)
}

// GetComp panics reason invalid parameter
func (v *virtualStrip) GetComp() bool {
	panic("invalid parameter Comp for virtualStrip")
}

// SetComp panics reason invalid parameter
func (v *virtualStrip) SetComp(val bool) {
	panic("invalid parameter Comp for virtualStrip")
}

// GetGate panics reason invalid parameter
func (v *virtualStrip) GetGate() bool {
	panic("invalid parameter Gate for virtualStrip")
}

// SetGate panics reason invalid parameter
func (v *virtualStrip) SetGate(val bool) {
	panic("invalid parameter Gate for virtualStrip")
}

// GetAudibility panics reason invalid parameter
func (v *virtualStrip) GetAudibility() bool {
	panic("invalid parameter Audibility for virtualStrip")
}

// SetAudibility panics reason invalid parameter
func (v *virtualStrip) SetAudibility(val bool) {
	panic("invalid parameter Audibility for virtualStrip")
}
