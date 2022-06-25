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
}

// String implements the stringer interface
func (r *remote) String() string {
	return fmt.Sprintf("Voicemeeter %s", r.kind)
}

func (r *remote) Login() {
	login(r.kind.name)
}

func (r *remote) Logout() {
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

type remoteBuilder interface {
	setKind() remoteBuilder
	makeStrip() remoteBuilder
	makeBus() remoteBuilder
	makeButton() remoteBuilder
	makeCommand() remoteBuilder
	makeVban() remoteBuilder
	Get() *remote
}

type director struct {
	builder remoteBuilder
}

func (d *director) SetBuilder(b remoteBuilder) {
	d.builder = b
}

func (d *director) Construct() {
	d.builder.setKind().makeStrip().makeBus().makeButton().makeCommand().makeVban()
}

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

func (b *genericBuilder) makeButton() remoteBuilder {
	fmt.Println("building button")
	_button := make([]button, 80)
	for i := 0; i < 80; i++ {
		_button[i] = newButton(i)
	}
	b.r.Button = _button
	return b
}

func (b *genericBuilder) makeCommand() remoteBuilder {
	fmt.Println("building command")
	_command := newCommand()
	b.r.Command = _command
	return b
}

func (b *genericBuilder) makeVban() remoteBuilder {
	fmt.Println("building vban")
	_vban := newVban(b.k)
	b.r.Vban = _vban
	return b
}

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
