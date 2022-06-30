package voicemeeter

import (
	"fmt"
	"os"
)

// A remote type represents the API for a kind,
// comprised of slices representing each member
type remote struct {
	kind     *kind
	Strip    []t_strip
	Bus      []t_bus
	Button   []button
	Command  *command
	Vban     *vban
	Device   *device
	Recorder *recorder

	pooler *pooler
}

// String implements the fmt.stringer interface
func (r *remote) String() string {
	return fmt.Sprintf("Voicemeeter %s", r.kind)
}

// Login logs into the API
// then it intializes the pooler
func (r *remote) Login() {
	r.pooler = newPooler()
	login(r.kind.name)
}

// Logout logs out of the API
// it also terminates the pooler
func (r *remote) Logout() {
	r.pooler.run = false
	logout()
}

func (r *remote) Type() string {
	return getVMType()
}

func (r *remote) Version() string {
	return getVersion()
}

// Pdirty returns true iff a parameter value has changed
func (r *remote) Pdirty() bool {
	return pdirty()
}

// Mdirty returns true iff a macrobutton value has changed
func (r *remote) Mdirty() bool {
	return mdirty()
}

func (r *remote) SendText(script string) {
	setParametersMulti(script)
}

// Register forwards the register method to Pooler
func (r *remote) Register(o observer) {
	r.pooler.Register(o)
}

// Register forwards the deregister method to Pooler
func (r *remote) Deregister(o observer) {
	r.pooler.Deregister(o)
}

type remoteBuilder interface {
	setKind() remoteBuilder
	makeStrip() remoteBuilder
	makeBus() remoteBuilder
	makeButton() remoteBuilder
	makeCommand() remoteBuilder
	makeVban() remoteBuilder
	makeDevice() remoteBuilder
	makeRecorder() remoteBuilder
	Build() remoteBuilder
	Get() *remote
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
func (d *director) Get() *remote {
	return d.builder.Get()
}

type genericBuilder struct {
	k *kind
	r remote
}

func (b *genericBuilder) setKind() remoteBuilder {
	b.r.kind = b.k
	return b
}

// makeStrip makes a strip slice and assigns it to remote.Strip
// []t_strip comprises of both physical and virtual strip types
func (b *genericBuilder) makeStrip() remoteBuilder {
	fmt.Println("building strip")
	_strip := make([]t_strip, b.k.numStrip())
	for i := 0; i < b.k.numStrip(); i++ {
		if i < b.k.physIn {
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
	_bus := make([]t_bus, b.k.numBus())
	for i := 0; i < b.k.numBus(); i++ {
		if i < b.k.physOut {
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

// makeCommand makes a Command type and assigns it to remote.Command
func (b *genericBuilder) makeCommand() remoteBuilder {
	fmt.Println("building command")
	b.r.Command = newCommand()
	return b
}

// makeVban makes a Vban type and assigns it to remote.Vban
func (b *genericBuilder) makeVban() remoteBuilder {
	fmt.Println("building vban")
	b.r.Vban = newVban(b.k)
	return b
}

// makeVban makes a Vban type and assigns it to remote.Vban
func (b *genericBuilder) makeDevice() remoteBuilder {
	fmt.Println("building recorder")
	b.r.Device = newDevice()
	return b
}

// makeRecorder makes a recorder type and assigns it to remote.Vban
func (b *genericBuilder) makeRecorder() remoteBuilder {
	fmt.Println("building recorder")
	b.r.Recorder = newRecorder()
	return b
}

// Get returns a fully constructed remote type for a kind
func (b *genericBuilder) Get() *remote {
	return &b.r
}

type basicBuilder struct {
	genericBuilder
}

// Build defines the steps required to build a basic type
func (basb *genericBuilder) Build() remoteBuilder {
	return basb.setKind().makeStrip().makeBus().makeButton().makeCommand().makeVban().makeDevice()
}

type bananaBuilder struct {
	genericBuilder
}

// Build defines the steps required to build a banana type
func (banb *bananaBuilder) Build() remoteBuilder {
	return banb.setKind().makeStrip().makeBus().makeButton().makeCommand().makeVban().makeDevice().makeRecorder()
}

type potatoBuilder struct {
	genericBuilder
}

// Build defines the steps required to build a potato type
func (potb *potatoBuilder) Build() remoteBuilder {
	return potb.setKind().makeStrip().makeBus().makeButton().makeCommand().makeVban().makeDevice().makeRecorder()
}

// GetRemote returns a remote type for a kind
// this is the interface entry point
func GetRemote(kindId string) *remote {
	_kind, ok := kindMap[kindId]
	if !ok {
		err := fmt.Errorf("unknown Voicemeeter kind '%s'", kindId)
		fmt.Println(err)
		os.Exit(1)
	}

	director := director{}
	switch _kind.name {
	case "basic":
		director.SetBuilder(&basicBuilder{genericBuilder{_kind, remote{}}})
	case "banana":
		director.SetBuilder(&bananaBuilder{genericBuilder{_kind, remote{}}})
	case "potato":
		director.SetBuilder(&potatoBuilder{genericBuilder{_kind, remote{}}})
	}
	director.Construct()
	return director.Get()
}
