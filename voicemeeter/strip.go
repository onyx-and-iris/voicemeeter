package voicemeeter

import (
	"fmt"
)

type t_strip interface {
	String() string
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
	GetComp() float64
	SetComp(val float32)
	GetGate() float64
	SetGate(val float32)
	GetAudibility() float64
	SetAudibility(val float32)
	GainLayer() []gainLayer
	Levels() *levels
	t_outputs
}

// strip represents a strip channel
type strip struct {
	iRemote
	outputs
	gainLayer []gainLayer
	levels
}

// GetMute returns the value of the Mute parameter
func (s *strip) GetMute() bool {
	return s.getter_bool("Mute")
}

// SetMute sets the value of the Mute parameter
func (s *strip) SetMute(val bool) {
	s.setter_bool("Mute", val)
}

// GetMono returns the value of the Mono parameter
func (s *strip) GetMono() bool {
	return s.getter_bool("Mono")
}

// SetMono sets the value of the Mono parameter
func (s *strip) SetMono(val bool) {
	s.setter_bool("Mono", val)
}

// GetSolo returns the value of the Solo parameter
func (s *strip) GetSolo() bool {
	return s.getter_bool("Solo")
}

// SetSolo sets the value of the Solo parameter
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

// GetLabel returns the value of the Label parameter
func (s *strip) GetLabel() string {
	return s.getter_string("Label")
}

// SetLabel sets the value of the Label parameter
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

// GainLayer returns the gainlayer field
func (s *strip) GainLayer() []gainLayer {
	return s.gainLayer
}

// Levels returns the gainlayer field
func (s *strip) Levels() *levels {
	return &s.levels
}

type physicalStrip struct {
	strip
}

func newPhysicalStrip(i int, k *kind) t_strip {
	o := newOutputs("strip[%d]", i)
	gl := make([]gainLayer, 8)
	for j := 0; j < 8; j++ {
		gl[j] = newGainLayer(i, j)
	}
	l := newStripLevels(i, k)
	ps := physicalStrip{strip{iRemote{fmt.Sprintf("strip[%d]", i), i}, o, gl, l}}
	return t_strip(&ps)
}

// implement fmt.stringer interface in fmt
func (p *physicalStrip) String() string {
	return fmt.Sprintf("PhysicalStrip%d", p.index)
}

// GetComp returns the value of the Comp parameter
func (p *physicalStrip) GetComp() float64 {
	return p.getter_float("Comp")
}

// SetComp sets the value of the Comp parameter
func (p *physicalStrip) SetComp(val float32) {
	p.setter_float("Comp", val)
}

// GetGate returns the value of the Gate parameter
func (p *physicalStrip) GetGate() float64 {
	return p.getter_float("Gate")
}

// SetGate sets the value of the Gate parameter
func (p *physicalStrip) SetGate(val float32) {
	p.setter_float("Gate", val)
}

// GetAudibility returns the value of the Audibility parameter
func (p *physicalStrip) GetAudibility() float64 {
	return p.getter_float("Audibility")
}

// SetAudibility sets the value of the Audibility parameter
func (p *physicalStrip) SetAudibility(val float32) {
	p.setter_float("Audibility", val)
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
	o := newOutputs("strip[%d]", i)
	gl := make([]gainLayer, 8)
	for j := 0; j < 8; j++ {
		gl[j] = newGainLayer(i, j)
	}
	l := newStripLevels(i, k)
	vs := virtualStrip{strip{iRemote{fmt.Sprintf("strip[%d]", i), i}, o, gl, l}}
	return t_strip(&vs)
}

// implement fmt.stringer interface in fmt
func (v *virtualStrip) String() string {
	return fmt.Sprintf("VirtualStrip%d", v.index)
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
func (v *virtualStrip) GetComp() float64 {
	panic("invalid parameter Comp for virtualStrip")
}

// SetComp panics reason invalid parameter
func (v *virtualStrip) SetComp(val float32) {
	panic("invalid parameter Comp for virtualStrip")
}

// GetGate panics reason invalid parameter
func (v *virtualStrip) GetGate() float64 {
	panic("invalid parameter Gate for virtualStrip")
}

// SetGate panics reason invalid parameter
func (v *virtualStrip) SetGate(val float32) {
	panic("invalid parameter Gate for virtualStrip")
}

// GetAudibility panics reason invalid parameter
func (v *virtualStrip) GetAudibility() float64 {
	panic("invalid parameter Audibility for virtualStrip")
}

// SetAudibility panics reason invalid parameter
func (v *virtualStrip) SetAudibility(val float32) {
	panic("invalid parameter Audibility for virtualStrip")
}

type gainLayer struct {
	iRemote
	index int
}

func newGainLayer(i, j int) gainLayer {
	return gainLayer{iRemote{fmt.Sprintf("strip[%d]", i), i}, j}
}

func (gl *gainLayer) Get() float64 {
	return gl.getter_float(fmt.Sprintf("gainlayer[%d]", gl.index))
}

func (gl *gainLayer) Set(val float32) {
	gl.setter_float(fmt.Sprintf("gainlayer[%d]", gl.index), val)
}

func newStripLevels(i int, k *kind) levels {
	var init int
	var os int
	if i < k.physIn {
		init = i * 2
		os = 2
	} else {
		init = (k.physIn * 2) + ((i - k.physIn) * 8)
		os = 8
	}
	return levels{iRemote{fmt.Sprintf("strip[%d]", i), i}, k, init, os}
}

func (l *levels) PreFader() []float32 {
	var levels []float32
	for i := l.init; i < l.init+l.offset; i++ {
		levels = append(levels, l.convertLevel(getLevel(0, i)))
	}
	return levels
}

func (l *levels) PostFader() []float32 {
	var levels []float32
	for i := l.init; i < l.init+l.offset; i++ {
		levels = append(levels, l.convertLevel(getLevel(1, i)))
	}
	return levels
}

func (l *levels) PostMute() []float32 {
	var levels []float32
	for i := l.init; i < l.init+l.offset; i++ {
		levels = append(levels, l.convertLevel(getLevel(2, i)))
	}
	return levels
}
