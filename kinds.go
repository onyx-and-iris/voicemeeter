package voicemeeter

import (
	"fmt"
	"strings"
)

var basic, banana, potato *kind

// A kind represents a Voicemeeter kinds layout
type kind struct {
	Name                                              string
	PhysIn, VirtIn, PhysOut, VirtOut, VbanIn, VbanOut int
}

// numStrip returns the total number of strips for a kind
func (k *kind) NumStrip() int {
	n := k.PhysIn + k.VirtIn
	return n
}

// numBus returns the total number of buses for a kind
func (k *kind) NumBus() int {
	n := k.PhysOut + k.VirtOut
	return n
}

// String implements the fmt.stringer interface
func (k *kind) String() string {
	return fmt.Sprintf("%s%s", strings.ToUpper(k.Name[:1]), k.Name[1:])
}

// newBasicKind returns a basic kind struct address
func newBasicKind() *kind {
	if basic == nil {
		basic = &kind{"basic", 2, 1, 1, 1, 4, 4}
	}
	return basic
}

// newBananaKind returns a banana kind struct address
func newBananaKind() *kind {
	if banana == nil {
		banana = &kind{"banana", 3, 2, 3, 2, 8, 8}
	}
	return banana
}

// newPotatoKind returns a potato kind struct address
func newPotatoKind() *kind {
	if potato == nil {
		potato = &kind{"potato", 5, 3, 5, 3, 8, 8}
	}
	return potato
}

var (
	kindMap = map[string]*kind{
		"basic":  newBasicKind(),
		"banana": newBananaKind(),
		"potato": newPotatoKind(),
	}
)
