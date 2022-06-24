package voicemeeter

import (
	"fmt"
	"strings"
)

// A kind represents a Voicemeeter kinds layout
type kind struct {
	name                                              string
	physIn, virtIn, physOut, virtOut, vbanIn, vbanOut int
}

func (k *kind) numStrip() int {
	n := k.physIn + k.virtIn
	return n
}

func (k *kind) numBus() int {
	n := k.physOut + k.virtOut
	return n
}

func (k *kind) String() string {
	return fmt.Sprintf("%s%s", strings.ToUpper(k.name[:1]), k.name[1:])
}

// newBasicKind returns a basic kind struct address
func newBasicKind() *kind {
	return &kind{"basic", 2, 1, 1, 1, 4, 4}
}

// newBananaKind returns a banana kind struct address
func newBananaKind() *kind {
	return &kind{"banana", 3, 2, 3, 2, 8, 8}
}

// newPotatoKind returns a potato kind struct address
func newPotatoKind() *kind {
	return &kind{"potato", 5, 3, 5, 3, 8, 8}
}

var (
	kindMap = map[string]*kind{
		"basic":  newBasicKind(),
		"banana": newBananaKind(),
		"potato": newPotatoKind(),
	}
)
