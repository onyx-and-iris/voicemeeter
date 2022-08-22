package voicemeeter

import (
	"fmt"
	"os"
)

// A Remote type represents the API for a kind
type Remote struct {
	Kind     *kind
	Strip    []iStrip
	Bus      []iBus
	Button   []button
	Command  *command
	Vban     *vban
	Device   *device
	Recorder *recorder
	Midi     *midi_t

	pooler *pooler
}

// String implements the fmt.stringer interface
func (r *Remote) String() string {
	return fmt.Sprintf("Voicemeeter %s", r.Kind)
}

// Login logs into the API
// then it intializes the pooler
func (r *Remote) Login() {
	login(r.Kind.Name)
	r.pooler = newPooler(r.Kind)
}

// Logout logs out of the API
// it also terminates the pooler
func (r *Remote) Logout() {
	r.pooler.run = false
	logout()
}

// Type returns the type of Voicemeeter (basic, banana, potato)
func (r *Remote) Type() string {
	return getVMType()
}

// Version returns the version of Voicemeeter as a string
func (r *Remote) Version() string {
	return getVersion()
}

// Pdirty returns true iff a parameter value has changed
func (r *Remote) Pdirty() bool {
	return pdirty()
}

// Mdirty returns true iff a macrobutton value has changed
func (r *Remote) Mdirty() bool {
	return mdirty()
}

// Gets a float parameter value
func (r *Remote) GetFloat(name string) float64 {
	return getParameterFloat(name)
}

// Sets a float paramter value
func (r *Remote) SetFloat(name string, value float32) {
	setParameterFloat(name, value)
}

// Gets a string parameter value
func (r *Remote) GetString(name string) string {
	return getParameterString(name)
}

// Sets a string paramter value
func (r *Remote) SetString(name, value string) {
	setParameterString(name, value)
}

// SendText sets multiple parameters by script
func (r *Remote) SendText(script string) {
	setParametersMulti(script)
}

// Register forwards the register method to Pooler
func (r *Remote) Register(o observer) {
	r.pooler.Register(o)
}

// Deregister forwards the deregister method to Pooler
func (r *Remote) Deregister(o observer) {
	r.pooler.Deregister(o)
}

// EventAdd adds an event to the Pooler
func (r *Remote) EventAdd(event string) {
	r.pooler.event.Add(event)
}

// EventRemove removes an event from the Pooler
func (r *Remote) EventRemove(event string) {
	r.pooler.event.Remove(event)
}

// remoteBuilder defines the interface builder types must satisfy
type remoteBuilder interface {
	setKind() remoteBuilder
	makeStrip() remoteBuilder
	makeBus() remoteBuilder
	makeButton() remoteBuilder
	makeCommand() remoteBuilder
	makeVban() remoteBuilder
	makeDevice() remoteBuilder
	makeRecorder() remoteBuilder
	makeMidi() remoteBuilder
	Build() remoteBuilder
	Get() *Remote
}

// directory is responsible for directing the genericBuilder
type director struct {
	builder remoteBuilder
}

// SetBuilder sets the appropriate builder type for a kind
func (d *director) SetBuilder(b remoteBuilder) {
	d.builder = b
}

// Construct calls the build function for the specific builder
func (d *director) Construct() {
	d.builder.Build()
}

// Get forwards the Get method to the builder
func (d *director) Get() *Remote {
	return d.builder.Get()
}

// genericBuilder represents a generic builder type
type genericBuilder struct {
	k *kind
	r Remote
}

// setKind sets the kind for a builder of a kind
func (b *genericBuilder) setKind() remoteBuilder {
	b.r.Kind = b.k
	return b
}

// makeStrip makes a strip slice and assigns it to remote.Strip
// []iStrip comprises of both physical and virtual strip types
func (b *genericBuilder) makeStrip() remoteBuilder {
	fmt.Println("building strip")
	_strip := make([]iStrip, b.k.NumStrip())
	for i := 0; i < b.k.NumStrip(); i++ {
		if i < b.k.PhysIn {
			_strip[i] = newPhysicalStrip(i, b.k)
		} else {
			_strip[i] = newVirtualStrip(i, b.k)
		}
	}
	b.r.Strip = _strip
	return b
}

// makeBus makes a bus slice and assigns it to remote.Bus
// []t_bus comprises of both physical and virtual bus types
func (b *genericBuilder) makeBus() remoteBuilder {
	fmt.Println("building bus")
	_bus := make([]iBus, b.k.NumBus())
	for i := 0; i < b.k.NumBus(); i++ {
		if i < b.k.PhysOut {
			_bus[i] = newPhysicalBus(i, b.k)
		} else {
			_bus[i] = newVirtualBus(i, b.k)
		}
	}
	b.r.Bus = _bus
	return b
}

// makeButton makes a button slice and assigns it to remote.Button
func (b *genericBuilder) makeButton() remoteBuilder {
	fmt.Println("building button")
	_button := make([]button, 80)
	for i := 0; i < 80; i++ {
		_button[i] = newButton(i)
	}
	b.r.Button = _button
	return b
}

// makeCommand makes a command type and assigns it to remote.Command
func (b *genericBuilder) makeCommand() remoteBuilder {
	fmt.Println("building command")
	b.r.Command = newCommand()
	return b
}

// makeVban makes a vban type and assigns it to remote.Vban
func (b *genericBuilder) makeVban() remoteBuilder {
	fmt.Println("building vban")
	b.r.Vban = newVban(b.k)
	return b
}

// makeDevice makes a device type and assigns it to remote.Device
func (b *genericBuilder) makeDevice() remoteBuilder {
	fmt.Println("building device")
	b.r.Device = newDevice()
	return b
}

// makeRecorder makes a recorder type and assigns it to remote.Recorder
func (b *genericBuilder) makeRecorder() remoteBuilder {
	fmt.Println("building recorder")
	b.r.Recorder = newRecorder()
	return b
}

// makeMidi makes a midi type and assigns it to remote.Midi
func (b *genericBuilder) makeMidi() remoteBuilder {
	fmt.Println("building midi")
	b.r.Midi = newMidi()
	return b
}

// Get returns a fully constructed remote type for a kind
func (b *genericBuilder) Get() *Remote {
	return &b.r
}

// basicBuilder represents a builder specific to basic type
type basicBuilder struct {
	genericBuilder
}

// Build defines the steps required to build a basic type
func (basb *genericBuilder) Build() remoteBuilder {
	return basb.setKind().
		makeStrip().
		makeBus().
		makeButton().
		makeCommand().
		makeVban().
		makeDevice().
		makeMidi()
}

// bananaBuilder represents a builder specific to banana type
type bananaBuilder struct {
	genericBuilder
}

// Build defines the steps required to build a banana type
func (banb *bananaBuilder) Build() remoteBuilder {
	return banb.setKind().
		makeStrip().
		makeBus().
		makeButton().
		makeCommand().
		makeVban().
		makeDevice().
		makeRecorder().
		makeMidi()
}

// potatoBuilder represents a builder specific to potato type
type potatoBuilder struct {
	genericBuilder
}

// Build defines the steps required to build a potato type
func (potb *potatoBuilder) Build() remoteBuilder {
	return potb.setKind().
		makeStrip().
		makeBus().
		makeButton().
		makeCommand().
		makeVban().
		makeDevice().
		makeRecorder().
		makeMidi()
}

// NewRemote returns a Remote type for a kind
// this is the interface entry point
func NewRemote(kindId string) *Remote {
	_kind, ok := kindMap[kindId]
	if !ok {
		err := fmt.Errorf("unknown Voicemeeter kind '%s'", kindId)
		fmt.Println(err)
		os.Exit(1)
	}

	director := director{}
	switch _kind.Name {
	case "basic":
		director.SetBuilder(&basicBuilder{genericBuilder{_kind, Remote{}}})
	case "banana":
		director.SetBuilder(&bananaBuilder{genericBuilder{_kind, Remote{}}})
	case "potato":
		director.SetBuilder(&potatoBuilder{genericBuilder{_kind, Remote{}}})
	}
	director.Construct()
	return director.Get()
}
