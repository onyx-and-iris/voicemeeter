package voicemeeter

import (
	"fmt"
)

type t_bus interface {
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
	SetGain(val float32)
	Mode() t_busMode
}

// bus represents a bus channel
type bus struct {
	iRemote
	mode busMode
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

// Mode returns address of a busMode struct
func (b *bus) Mode() t_busMode {
	return &b.mode
}

type physicalBus struct {
	bus
}

func newPhysicalBus(i int) t_bus {
	pb := physicalBus{bus{
		iRemote{fmt.Sprintf("bus[%d]", i), i},
		newBusMode(i),
	}}
	return t_bus(&pb)
}

// String implements the fmt.stringer interface
func (p *physicalBus) String() string {
	return fmt.Sprintf("PhysicalBus%d", p.index)
}

type virtualBus struct {
	bus
}

func newVirtualBus(i int) t_bus {
	vb := virtualBus{bus{
		iRemote{fmt.Sprintf("bus[%d]", i), i},
		newBusMode(i),
	}}
	return t_bus(&vb)
}

// String implements the fmt.stringer interface
func (v *virtualBus) String() string {
	return fmt.Sprintf("VirtualBus%d", v.index)
}

type t_busMode interface {
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

type busMode struct {
	iRemote
}

func newBusMode(i int) busMode {
	return busMode{iRemote{fmt.Sprintf("bus[%d].mode", i), i}}
}

func (bm *busMode) SetNormal(val bool) {
	bm.setter_bool("Normal", val)
}

func (bm *busMode) GetNormal() bool {
	return bm.getter_bool("Normal")
}

func (bm *busMode) SetAmix(val bool) {
	bm.setter_bool("Amix", val)
}

func (bm *busMode) GetAmix() bool {
	return bm.getter_bool("Amix")
}

func (bm *busMode) SetBmix(val bool) {
	bm.setter_bool("Bmix", val)
}

func (bm *busMode) GetBmix() bool {
	return bm.getter_bool("Bmix")
}

func (bm *busMode) SetRepeat(val bool) {
	bm.setter_bool("Repeat", val)
}

func (bm *busMode) GetRepeat() bool {
	return bm.getter_bool("Repeat")
}

func (bm *busMode) SetComposite(val bool) {
	bm.setter_bool("Composite", val)
}

func (bm *busMode) GetComposite() bool {
	return bm.getter_bool("Composite")
}

func (bm *busMode) SetTvMix(val bool) {
	bm.setter_bool("TvMix", val)
}

func (bm *busMode) GetTvMix() bool {
	return bm.getter_bool("TvMix")
}

func (bm *busMode) SetUpMix21(val bool) {
	bm.setter_bool("UpMix21", val)
}

func (bm *busMode) GetUpMix21() bool {
	return bm.getter_bool("UpMix21")
}

func (bm *busMode) SetUpMix41(val bool) {
	bm.setter_bool("UpMix41", val)
}

func (bm *busMode) GetUpMix41() bool {
	return bm.getter_bool("UpMix41")
}

func (bm *busMode) SetUpMix61(val bool) {
	bm.setter_bool("UpMix61", val)
}

func (bm *busMode) GetUpMix61() bool {
	return bm.getter_bool("UpMix61")
}

func (bm *busMode) SetCenterOnly(val bool) {
	bm.setter_bool("CenterOnly", val)
}

func (bm *busMode) GetCenterOnly() bool {
	return bm.getter_bool("CenterOnly")
}

func (bm *busMode) SetLfeOnly(val bool) {
	bm.setter_bool("LfeOnly", val)
}

func (bm *busMode) GetLfeOnly() bool {
	return bm.getter_bool("LfeOnly")
}

func (bm *busMode) SetRearOnly(val bool) {
	bm.setter_bool("RearOnly", val)
}

func (bm *busMode) GetRearOnly() bool {
	return bm.getter_bool("RearOnly")
}
