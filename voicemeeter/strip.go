package voicemeeter

import (
	"fmt"
)

// custom strip type, struct forwarding channel
type strip struct {
	channel
}

// constructor method for strip
func newStrip(i int, k *kind) strip {
	return strip{channel{"strip", i, *k}}
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

// GetLimit returns the value of the Limit parameter
func (s *strip) GetLimit() int {
	return s.getter_int("Limit")
}

// SetLimit sets the value of the Limit parameter
func (s *strip) SetLimit(val int) {
	s.setter_int("Limit", val)
}

// GetMc returns the value of the MC parameter
func (s *strip) GetMc() bool {
	return s.getter_bool("MC")
}

// SetMc sets the value of the MC parameter
func (s *strip) SetMc(val bool) {
	s.setter_bool("MC", val)
}

// GetMc returns the value of the MC parameter
func (s *strip) GetLabel() string {
	return s.getter_string("Label")
}

// SetMc sets the value of the MC parameter
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
