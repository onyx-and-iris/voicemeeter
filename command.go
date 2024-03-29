package voicemeeter

// command represents command (action) type parameters
type command struct {
	iRemote
}

// newCommand returns a pointer to a command type
func newCommand() *command {
	return &command{iRemote{"command", 0}}
}

// Show shows the Voicemeete GUI if it's hidden
func (c *command) Show() {
	c.setter_float("Show", 1)
}

// Hide hides the Voicemeete GUI if it's shown
func (c *command) Hide() {
	c.setter_float("Show", 0)
}

// Shutdown shutdown the Voicemeeter GUI
func (c *command) Shutdown() {
	c.setter_float("Shutdown", 1)
}

// Restart restarts the Voicemeeter audio engine
func (c *command) Restart() {
	c.setter_float("Restart", 1)
}

// Lock locks or unlocks the Voiceemeter GUI
func (c *command) Lock(val bool) {
	var value float64
	if val {
		value = 1
	} else {
		value = 0
	}
	c.setter_float("Lock", value)
}
