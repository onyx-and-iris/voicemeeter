package voicemeeter

import (
	"fmt"
	"time"
)

// iBus defines the interface bus types must satisfy
type iBus interface {
	String() string
	Mute() bool
	SetMute(val bool)
	Mono() bool
	SetMono(val bool)
	Label() string
	SetLabel(val string)
	Gain() float64
	SetGain(val float64)
	Eq() *eQ
	Mode() *busMode
	Levels() *levels
	FadeTo(target float32, time_ int)
	FadeBy(change float32, time_ int)
}

// bus represents a bus channel
type bus struct {
	iRemote
	eQ   *eQ
	mode *busMode
	levels
}

// Mute returns the value of the Mute parameter
func (b *bus) Mute() bool {
	return b.getter_bool("Mute")
}

// SetMute sets the value of the Mute parameter
func (b *bus) SetMute(val bool) {
	b.setter_bool("Mute", val)
}

// Mono returns the value of the Mute parameter
func (b *bus) Mono() bool {
	return b.getter_bool("Mono")
}

// SetMono sets the value of the Mute parameter
func (b *bus) SetMono(val bool) {
	b.setter_bool("Mono", val)
}

// Label returns the value of the MC parameter
func (b *bus) Label() string {
	return b.getter_string("Label")
}

// SetLabel sets the value of the MC parameter
func (b *bus) SetLabel(val string) {
	b.setter_string("Label", val)
}

// Gain returns the value of the Gain parameter
func (b *bus) Gain() float64 {
	return b.getter_float("Gain")
}

// SetGain sets the value of the Gain parameter
func (b *bus) SetGain(val float64) {
	b.setter_float("Gain", val)
}

// Eq returns the eQ field
func (b *bus) Eq() *eQ {
	return b.eQ
}

// Mode returns address of a busMode struct
func (b *bus) Mode() *busMode {
	return b.mode
}

// Levels returns the levels field
func (b *bus) Levels() *levels {
	return &b.levels
}

// FadeTo sets the value of gain to target over at time interval of time_
func (b *bus) FadeTo(target float32, time_ int) {
	b.setter_string("FadeTo", fmt.Sprintf("(\"%f\", %d)", target, time_))
	time.Sleep(time.Millisecond)
}

// FadeBy adjusts the value of gain by change over a time interval of time_
func (b *bus) FadeBy(change float32, time_ int) {
	b.setter_string("FadeBy", fmt.Sprintf("(\"%f\", %d)", change, time_))
	time.Sleep(time.Millisecond)
}

// physicalBus represents a single physical bus
type physicalBus struct {
	bus
}

// newPhysicalBus returns a physicalBus type cast to an iBus
func newPhysicalBus(i int, k *kind) iBus {
	e := newEq(fmt.Sprintf("bus[%d].EQ", i), i)
	b := newBusMode(i)
	l := newBusLevels(i, k)
	pb := physicalBus{bus{iRemote{fmt.Sprintf("bus[%d]", i), i}, e, b, l}}

	return &pb
}

// String implements the fmt.stringer interface
func (p *physicalBus) String() string {
	return fmt.Sprintf("PhysicalBus%d", p.index)
}

// virtualBus represents a single virtual bus
type virtualBus struct {
	bus
}

// newVirtualBus returns a virtualBus type cast to an iBus
func newVirtualBus(i int, k *kind) iBus {
	e := newEq(fmt.Sprintf("bus[%d].EQ", i), i)
	b := newBusMode(i)
	l := newBusLevels(i, k)
	vb := virtualBus{bus{iRemote{fmt.Sprintf("bus[%d]", i), i}, e, b, l}}
	return &vb
}

// String implements the fmt.stringer interface
func (v *virtualBus) String() string {
	return fmt.Sprintf("VirtualBus%d", v.index)
}

// busMode offers methods for getting/setting bus mode states
type busMode struct {
	iRemote
}

// newBusMode returns a busMode struct
func newBusMode(i int) *busMode {
	return &busMode{iRemote{fmt.Sprintf("bus[%d].mode", i), i}}
}

// Normal gets the value of the Mode.Normal parameter
func (bm *busMode) Normal() bool {
	return bm.getter_bool("Normal")
}

// SetNormal sets the value of the Mode.Normal parameter
func (bm *busMode) SetNormal(val bool) {
	bm.setter_bool("Normal", val)
}

// Amix gets the value of the Mode.Amix parameter
func (bm *busMode) Amix() bool {
	return bm.getter_bool("Amix")
}

// SetAmix sets the value of the Mode.Amix parameter
func (bm *busMode) SetAmix(val bool) {
	bm.setter_bool("Amix", val)
}

// Bmix gets the value of the Mode.Bmix parameter
func (bm *busMode) Bmix() bool {
	return bm.getter_bool("Bmix")
}

// SetBmix sets the value of the Mode.Bmix parameter
func (bm *busMode) SetBmix(val bool) {
	bm.setter_bool("Bmix", val)
}

// Repeat gets the value of the Mode.Repeat parameter
func (bm *busMode) Repeat() bool {
	return bm.getter_bool("Repeat")
}

// SetRepeat sets the value of the Mode.Repeat parameter
func (bm *busMode) SetRepeat(val bool) {
	bm.setter_bool("Repeat", val)
}

// Composite gets the value of the Mode.Composite parameter
func (bm *busMode) Composite() bool {
	return bm.getter_bool("Composite")
}

// SetComposite sets the value of the Mode.Composite parameter
func (bm *busMode) SetComposite(val bool) {
	bm.setter_bool("Composite", val)
}

// TvMix gets the value of the Mode.TvMix parameter
func (bm *busMode) TvMix() bool {
	return bm.getter_bool("TvMix")
}

// SetTvMix sets the value of the Mode.TvMix parameter
func (bm *busMode) SetTvMix(val bool) {
	bm.setter_bool("TvMix", val)
}

// UpMix21 gets the value of the Mode.UpMix21 parameter
func (bm *busMode) UpMix21() bool {
	return bm.getter_bool("UpMix21")
}

// SetUpMix21 sets the value of the Mode.UpMix21 parameter
func (bm *busMode) SetUpMix21(val bool) {
	bm.setter_bool("UpMix21", val)
}

// UpMix41 gets the value of the Mode.UpMix41 parameter
func (bm *busMode) UpMix41() bool {
	return bm.getter_bool("UpMix41")
}

// SetUpMix41 sets the value of the Mode.UpMix41 parameter
func (bm *busMode) SetUpMix41(val bool) {
	bm.setter_bool("UpMix41", val)
}

// UpMix61 gets the value of the Mode.UpMix61 parameter
func (bm *busMode) UpMix61() bool {
	return bm.getter_bool("UpMix61")
}

// SetUpMix61 sets the value of the Mode.UpMix61 parameter
func (bm *busMode) SetUpMix61(val bool) {
	bm.setter_bool("UpMix61", val)
}

// CenterOnly gets the value of the Mode.CenterOnly parameter
func (bm *busMode) CenterOnly() bool {
	return bm.getter_bool("CenterOnly")
}

// SetCenterOnly sets the value of the Mode.CenterOnly parameter
func (bm *busMode) SetCenterOnly(val bool) {
	bm.setter_bool("CenterOnly", val)
}

// LfeOnly gets the value of the Mode.LFE parameter
func (bm *busMode) LfeOnly() bool {
	return bm.getter_bool("LfeOnly")
}

// SetLfeOnly sets the value of the Mode.LFE parameter
func (bm *busMode) SetLfeOnly(val bool) {
	bm.setter_bool("LfeOnly", val)
}

// RearOnly gets the value of the Mode.RearOnly parameter
func (bm *busMode) RearOnly() bool {
	return bm.getter_bool("RearOnly")
}

// SetRearOnly sets the value of the Mode.RearOnly parameter
func (bm *busMode) SetRearOnly(val bool) {
	bm.setter_bool("RearOnly", val)
}

// newBusLevels represents the levels field for a channel
func newBusLevels(i int, k *kind) levels {
	init := i * 8
	return levels{iRemote{fmt.Sprintf("bus[%d]", i), i}, k, init, 8, "bus"}
}

// All returns the level values for a bus
func (l *levels) All() []float64 {
	var levels []float64
	for i := l.init; i < l.init+l.offset; i++ {
		levels = append(levels, convertLevel(_levelCache.busLevels[i]))
	}
	return levels
}
