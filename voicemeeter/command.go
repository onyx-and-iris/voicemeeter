package voicemeeter

import "fmt"

type command struct {
	identifier string
}

func newCommand() *command {
	return &command{"command"}
}

func (c *command) setter(p string, v float32) {
	param := fmt.Sprintf("%s.%s", c.identifier, p)
	setParameterFloat(param, v)
}

func (c *command) Show() {
	c.setter("Show", 1)
}

func (c *command) Hide() {
	c.setter("Show", 0)
}

func (c *command) Shutdown() {
	c.setter("Shutdown", 1)
}

func (c *command) Restart() {
	c.setter("Restart", 1)
}

func (c *command) Lock(val bool) {
	var value float32
	if val {
		value = 1
	} else {
		value = 0
	}
	c.setter("Lock", value)
}
