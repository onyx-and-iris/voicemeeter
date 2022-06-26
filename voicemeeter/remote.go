package voicemeeter

import (
	"fmt"
	"os"
)

// A remote type represents the API for a kind,
// comprised of slices representing each member
type remote struct {
	kind    *kind
	Strip   []t_strip
	Bus     []t_bus
	Button  []button
	Command *command
	Vban    *vban

	Pooler *Pooler
}

// String implements the fmt.stringer interface
func (r *remote) String() string {
	return fmt.Sprintf("Voicemeeter %s", r.kind)
}

// Login logs into the API
// then it intializes the pooler
func (r *remote) Login() {
	r.Pooler = newPooler()
	login(r.kind.name)
}

// Logout logs out of the API
// it also terminates the pooler
func (r *remote) Logout() {
	r.Pooler.run = false
	logout()
}

func (r *remote) Type() string {
	return getVMType()
}

func (r *remote) Version() string {
	return getVersion()
}

func (r *remote) SendText(script string) {
	setParametersMulti(script)
}

// Register forwards the register method to Pooler
func (r *remote) Register(o observer) {
	r.Pooler.Register(o)
}

// Register forwards the deregister method to Pooler
func (r *remote) Deregister(o observer) {
	r.Pooler.Deregister(o)
}

type remoteBuilder interface {
	setKind() remoteBuilder
	makeStrip() remoteBuilder
	makeBus() remoteBuilder
	makeButton() remoteBuilder
	makeCommand() remoteBuilder
	makeVban() remoteBuilder
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

// Construct defines the steps required for building a remote type
func (d *director) Construct() {
	d.builder.setKind().makeStrip().makeBus().makeButton().makeCommand().makeVban()
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
			_strip[i] = newPhysicalStrip(i)
		} else {
			_strip[i] = newVirtualStrip(i)
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
			_bus[i] = newPhysicalBus(i)
		} else {
			_bus[i] = newVirtualBus(i)
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

// makeCommand makes a Command type and assignss it to remote.Command
func (b *genericBuilder) makeCommand() remoteBuilder {
	fmt.Println("building command")
	b.r.Command = newCommand()
	return b
}

// makeVban makes a Vban type and assignss it to remote.Vban
func (b *genericBuilder) makeVban() remoteBuilder {
	fmt.Println("building vban")
	b.r.Vban = newVban(b.k)
	return b
}

// Get returns a fully constructed remote type for a kind
func (b *genericBuilder) Get() *remote {
	return &b.r
}

type basicBuilder struct {
	genericBuilder
}

type bananaBuilder struct {
	genericBuilder
}

type potatoBuilder struct {
	genericBuilder
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
