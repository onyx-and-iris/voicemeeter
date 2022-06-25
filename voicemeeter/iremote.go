package voicemeeter

import (
	"fmt"
)

// iRemote provides a set of common forwarding methods
type iRemote struct {
	_identifier string
	index       int
}

func (c *iRemote) identifier() string {
	return c._identifier
}

// getter_bool returns the value of a boolean parameter
func (c *iRemote) getter_bool(p string) bool {
	param := fmt.Sprintf("%s[%d].%s", c.identifier(), c.index, p)
	return getParameterFloat(param) == 1
}

// setter_bool sets the value of a boolean parameter
func (c *iRemote) setter_bool(p string, v bool) {
	param := fmt.Sprintf("%s[%d].%s", c.identifier(), c.index, p)
	var value float32
	if v {
		value = 1
	} else {
		value = 0
	}
	setParameterFloat(param, float32(value))
}

// getter_int returns the value of an int parameter p
func (c *iRemote) getter_int(p string) int {
	param := fmt.Sprintf("%s[%d].%s", c.identifier(), c.index, p)
	return int(getParameterFloat(param))
}

// setter_int sets the value v of an int parameter p
func (c *iRemote) setter_int(p string, v int) {
	param := fmt.Sprintf("%s[%d].%s", c.identifier(), c.index, p)
	setParameterFloat(param, float32(v))
}

// getter_float returns the value of an int parameter p
func (c *iRemote) getter_float(p string) float64 {
	param := fmt.Sprintf("%s[%d].%s", c.identifier(), c.index, p)
	return getParameterFloat(param)
}

// setter_float sets the value v of an int parameter p
func (c *iRemote) setter_float(p string, v float32) {
	param := fmt.Sprintf("%s[%d].%s", c.identifier(), c.index, p)
	setParameterFloat(param, float32(v))
}

// getter_string returns the value of a string parameter p
func (c *iRemote) getter_string(p string) string {
	param := fmt.Sprintf("%s[%d].%s", c.identifier(), c.index, p)
	return getParameterString(param)
}

// setter_string sets the value v of a string parameter p
func (c *iRemote) setter_string(p, v string) {
	param := fmt.Sprintf("%s[%d].%s", c.identifier(), c.index, p)
	setParameterString(param, v)
}
