package voicemeeter

import (
	"fmt"
	"time"
)

// iBus defines the interface bus types must satisfy
type iBus interface {
	String() string
	GetMute() bool
	SetMute(val bool)
	GetEq() bool
	SetEq(val bool)
	GetMono() bool
	SetMono(val bool)
	GetLabel() string
	SetLabel(val string)
	GetGain() float64
	SetGain(val float64)
	Mode() iBusMode
	Levels() *levels
	FadeTo(target float32, time_ int)
	FadeBy(change float32, time_ int)
}

// bus represents a bus channel
type bus struct {
	iRemote
	mode busMode
	levels
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
func (b *bus) SetGain(val float64) {
	b.setter_float("Gain", val)
}

// Mode returns address of a busMode struct
func (b *bus) Mode() iBusMode {
	return &b.mode
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
	b := newBusMode(i)
	l := newBusLevels(i, k)
	pb := physicalBus{bus{iRemote{fmt.Sprintf("bus[%d]", i), i}, b, l}}

	return iBus(&pb)
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
	b := newBusMode(i)
	l := newBusLevels(i, k)
	vb := virtualBus{bus{iRemote{fmt.Sprintf("bus[%d]", i), i}, b, l}}
	return iBus(&vb)
}

// String implements the fmt.stringer interface
func (v *virtualBus) String() string {
	return fmt.Sprintf("VirtualBus%d", v.index)
}

// iBusMode defines the interface busMode type must satisfy
type iBusMode interface {
	SetNormal(val bool)
	GetNormal() bool
	SetAmix(val bool)
	GetAmix() bool
	SetBmix(val bool)
	GetBmix() bool
	SetRepeat(val bool)
	GetRepeat() bool
	SetComposite(val bool)
	GetComposite() bool
	SetTvMix(val bool)
	GetTvMix() bool
	SetUpMix21(val bool)
	GetUpMix21() bool
	SetUpMix41(val bool)
	GetUpMix41() bool
	SetUpMix61(val bool)
	GetUpMix61() bool
	SetCenterOnly(val bool)
	GetCenterOnly() bool
	SetLfeOnly(val bool)
	GetLfeOnly() bool
	SetRearOnly(val bool)
	GetRearOnly() bool
}

// busMode offers methods for getting/setting bus mode states
type busMode struct {
	iRemote
}

// newBusMode returns a busMode struct
func newBusMode(i int) busMode {
	return busMode{iRemote{fmt.Sprintf("bus[%d].mode", i), i}}
}

// GetNormal gets the value of the Mode.Normal parameter
func (bm *busMode) GetNormal() bool {
	return bm.getter_bool("Normal")
}

// SetNormal sets the value of the Mode.Normal parameter
func (bm *busMode) SetNormal(val bool) {
	bm.setter_bool("Normal", val)
}

// GetAmix gets the value of the Mode.Amix parameter
func (bm *busMode) GetAmix() bool {
	return bm.getter_bool("Amix")
}

// SetAmix sets the value of the Mode.Amix parameter
func (bm *busMode) SetAmix(val bool) {
	bm.setter_bool("Amix", val)
}

// GetBmix gets the value of the Mode.Bmix parameter
func (bm *busMode) GetBmix() bool {
	return bm.getter_bool("Bmix")
}

// SetBmix sets the value of the Mode.Bmix parameter
func (bm *busMode) SetBmix(val bool) {
	bm.setter_bool("Bmix", val)
}

// GetRepeat gets the value of the Mode.Repeat parameter
func (bm *busMode) GetRepeat() bool {
	return bm.getter_bool("Repeat")
}

// SetRepeat sets the value of the Mode.Repeat parameter
func (bm *busMode) SetRepeat(val bool) {
	bm.setter_bool("Repeat", val)
}

// GetComposite gets the value of the Mode.Composite parameter
func (bm *busMode) GetComposite() bool {
	return bm.getter_bool("Composite")
}

// SetComposite sets the value of the Mode.Composite parameter
func (bm *busMode) SetComposite(val bool) {
	bm.setter_bool("Composite", val)
}

// GetTvMix gets the value of the Mode.TvMix parameter
func (bm *busMode) GetTvMix() bool {
	return bm.getter_bool("TvMix")
}

// SetTvMix sets the value of the Mode.TvMix parameter
func (bm *busMode) SetTvMix(val bool) {
	bm.setter_bool("TvMix", val)
}

// GetUpMix21 gets the value of the Mode.UpMix21 parameter
func (bm *busMode) GetUpMix21() bool {
	return bm.getter_bool("UpMix21")
}

// SetUpMix21 sets the value of the Mode.UpMix21 parameter
func (bm *busMode) SetUpMix21(val bool) {
	bm.setter_bool("UpMix21", val)
}

// GetUpMix41 gets the value of the Mode.UpMix41 parameter
func (bm *busMode) GetUpMix41() bool {
	return bm.getter_bool("UpMix41")
}

// SetUpMix41 sets the value of the Mode.UpMix41 parameter
func (bm *busMode) SetUpMix41(val bool) {
	bm.setter_bool("UpMix41", val)
}

// GetUpMix61 gets the value of the Mode.UpMix61 parameter
func (bm *busMode) GetUpMix61() bool {
	return bm.getter_bool("UpMix61")
}

// SetUpMix61 sets the value of the Mode.UpMix61 parameter
func (bm *busMode) SetUpMix61(val bool) {
	bm.setter_bool("UpMix61", val)
}

// GetCenterOnly gets the value of the Mode.CenterOnly parameter
func (bm *busMode) GetCenterOnly() bool {
	return bm.getter_bool("CenterOnly")
}

// SetCenterOnly sets the value of the Mode.CenterOnly parameter
func (bm *busMode) SetCenterOnly(val bool) {
	bm.setter_bool("CenterOnly", val)
}

// GetLfeOnly gets the value of the Mode.LFE parameter
func (bm *busMode) GetLfeOnly() bool {
	return bm.getter_bool("LfeOnly")
}

// SetLfeOnly sets the value of the Mode.LFE parameter
func (bm *busMode) SetLfeOnly(val bool) {
	bm.setter_bool("LfeOnly", val)
}

// GetRearOnly gets the value of the Mode.RearOnly parameter
func (bm *busMode) GetRearOnly() bool {
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
