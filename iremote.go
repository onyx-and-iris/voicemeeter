package voicemeeter

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

// iRemote provides an interface between higher methods and lower functions
// expected to be embedded
type iRemote struct {
	_identifier string
	index       int
}

// identifier returns a string identifier
func (ir *iRemote) identifier() string {
	return ir._identifier
}

// getter_bool returns the value of a boolean parameter
func (ir *iRemote) getter_bool(p string) bool {
	param := fmt.Sprintf("%s.%s", ir.identifier(), p)
	log.Debug("getter_bool::", param)
	val, err := getParameterFloat(param)
	if err != nil {
		fmt.Println(err)
	}
	return val == 1
}

// setter_bool sets the value of a boolean parameter
func (ir *iRemote) setter_bool(p string, v bool) {
	param := fmt.Sprintf("%s.%s", ir.identifier(), p)
	var value float32
	if v {
		value = 1
	} else {
		value = 0
	}
	log.Debug("setter_bool::", param, "=", v)
	err := setParameterFloat(param, float64(value))
	if err != nil {
		fmt.Println(err)
	}
}

// getter_int returns the value of an int parameter p
func (ir *iRemote) getter_int(p string) int {
	param := fmt.Sprintf("%s.%s", ir.identifier(), p)
	log.Debug("getter_int::", param)
	val, err := getParameterFloat(param)
	if err != nil {
		fmt.Println(err)
	}
	return int(val)
}

// setter_int sets the value v of an int parameter p
func (ir *iRemote) setter_int(p string, v int) {
	param := fmt.Sprintf("%s.%s", ir.identifier(), p)
	log.Debug("setter_int::", param, "=", v)
	err := setParameterFloat(param, float64(v))
	if err != nil {
		fmt.Println(err)
	}
}

// getter_float returns the value of a float parameter p
func (ir *iRemote) getter_float(p string) float64 {
	var param string
	if p != "" {
		param = fmt.Sprintf("%s.%s", ir.identifier(), p)
	} else {
		param = ir.identifier()
	}
	log.Debug("getter_float::", param)
	val, err := getParameterFloat(param)
	if err != nil {
		fmt.Println(err)
	}
	return val
}

// setter_float sets the value v of an float parameter p
func (ir *iRemote) setter_float(p string, v float64) {
	var param string
	if p != "" {
		param = fmt.Sprintf("%s.%s", ir.identifier(), p)
	} else {
		param = ir.identifier()
	}
	log.Debug("setter_float::", param, "=", v)
	err := setParameterFloat(param, float64(v))
	if err != nil {
		fmt.Println(err)
	}
}

// getter_string returns the value of a string parameter p
func (ir *iRemote) getter_string(p string) string {
	param := fmt.Sprintf("%s.%s", ir.identifier(), p)
	log.Debug("getter_string::", param)
	val, err := getParameterString(param)
	if err != nil {
		fmt.Println(err)
	}
	return val
}

// setter_string sets the value v of a string parameter p
func (ir *iRemote) setter_string(p, v string) {
	param := fmt.Sprintf("%s.%s", ir.identifier(), p)
	log.Debug("setter_string::", param, "=", v)
	err := setParameterString(param, v)
	if err != nil {
		fmt.Println(err)
	}
}
