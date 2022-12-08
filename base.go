package voicemeeter

import (
	"bytes"
	"fmt"
	"math"
	"strings"
	"syscall"
	"time"
	"unsafe"

	log "github.com/sirupsen/logrus"
)

var (
	mod = syscall.NewLazyDLL(getDllPath())

	vmLogin          = mod.NewProc("VBVMR_Login")
	vmLogout         = mod.NewProc("VBVMR_Logout")
	vmRunvm          = mod.NewProc("VBVMR_RunVoicemeeter")
	vmGetvmType      = mod.NewProc("VBVMR_GetVoicemeeterType")
	vmGetvmVersion   = mod.NewProc("VBVMR_GetVoicemeeterVersion")
	vmPdirty         = mod.NewProc("VBVMR_IsParametersDirty")
	vmGetParamFloat  = mod.NewProc("VBVMR_GetParameterFloat")
	vmGetParamString = mod.NewProc("VBVMR_GetParameterStringA")

	vmGetLevelFloat = mod.NewProc("VBVMR_GetLevel")

	vmSetParamFloat  = mod.NewProc("VBVMR_SetParameterFloat")
	vmSetParameters  = mod.NewProc("VBVMR_SetParameters")
	vmSetParamString = mod.NewProc("VBVMR_SetParameterStringA")

	vmGetDevNumIn   = mod.NewProc("VBVMR_Input_GetDeviceNumber")
	vmGetDevDescIn  = mod.NewProc("VBVMR_Input_GetDeviceDescA")
	vmGetDevNumOut  = mod.NewProc("VBVMR_Output_GetDeviceNumber")
	vmGetDevDescOut = mod.NewProc("VBVMR_Output_GetDeviceDescA")

	vmMdirty         = mod.NewProc("VBVMR_MacroButton_IsDirty")
	vmGetMacroStatus = mod.NewProc("VBVMR_MacroButton_GetStatus")
	vmSetMacroStatus = mod.NewProc("VBVMR_MacroButton_SetStatus")

	vmGetMidiMessage = mod.NewProc("VBVMR_GetMidiMessage")
)

// login logs into the API,
// attempts to launch Voicemeeter if it's not running,
// initializes dirty parameters.
func login(kindId string) error {
	res, _, _ := vmLogin.Call()
	if res == 1 {
		runVoicemeeter(kindId)
		time.Sleep(time.Second)
	} else if res != 0 {
		err := fmt.Errorf("VBVMR_Login returned %d", res)
		return err
	}
	log.Info("Logged into Voicemeeter ", kindId)
	clear()
	return nil
}

// logout logs out of the API,
// delayed for 100ms to allow final operation to complete.
func logout(kindId string) error {
	time.Sleep(100 * time.Millisecond)
	res, _, _ := vmLogout.Call()
	if int32(res) != 0 {
		err := fmt.Errorf("VBVMR_Logout returned %d", int32(res))
		return err
	}
	log.Info("Logged out of Voicemeeter ", kindId)
	return nil
}

// runVoicemeeter attempts to launch a Voicemeeter GUI of a kind.
func runVoicemeeter(kindId string) error {
	vals := map[string]uint64{
		"basic":  1,
		"banana": 2,
		"potato": 3,
	}
	res, _, _ := vmRunvm.Call(uintptr(vals[kindId]))
	if int32(res) != 0 {
		err := fmt.Errorf("VBVMR_RunVoicemeeter returned %d", int32(res))
		return err
	}
	return nil
}

// getVersion returns the version of Voicemeeter as a string
func getVersion() (string, error) {
	var ver uint64
	res, _, _ := vmGetvmVersion.Call(uintptr(unsafe.Pointer(&ver)))
	if int32(res) != 0 {
		err := fmt.Errorf("VBVMR_GetVoicemeeterVersion returned %d", int32(res))
		return "", err
	}
	v1 := (ver & 0xFF000000) >> 24
	v2 := (ver & 0x00FF0000) >> 16
	v3 := (ver & 0x0000FF00) >> 8
	v4 := ver & 0x000000FF
	return fmt.Sprintf("%d.%d.%d.%d", v1, v2, v3, v4), nil
}

// pdirty returns true iff a parameter value has changed
func pdirty() (bool, error) {
	res, _, _ := vmPdirty.Call()
	if int32(res) < 0 {
		err := fmt.Errorf("VBVMR_IsParametersDirty returned %d", int32(res))
		log.Error(err.Error())
		return false, err
	}
	return int32(res) == 1, nil
}

// mdirty returns true iff a macrobutton value has changed
func mdirty() (bool, error) {
	res, _, _ := vmMdirty.Call()
	if int32(res) < 0 {
		err := fmt.Errorf("VBVMR_MacroButton_IsDirty returned %d", int32(res))
		log.Error(err.Error())
		return false, err
	}
	return int32(res) == 1, nil
}

// ldirty returns true iff a level value has changed
func ldirty(k *kind) (bool, error) {
	_levelCache.stripLevelsBuff = make([]float64, (2*k.PhysIn)+(8*k.VirtIn))
	_levelCache.busLevelsBuff = make([]float64, 8*k.NumBus())

	for i := 0; i < (2*k.PhysIn)+(8*k.VirtIn); i++ {
		val, err := getLevel(_levelCache.stripMode, i)
		if err != nil {
			log.Error(err.Error())
			return false, err
		}
		_levelCache.stripLevelsBuff[i] = val
		_levelCache.stripComp[i] = _levelCache.stripLevelsBuff[i] == _levelCache.stripLevels[i]
	}
	for i := 0; i < 8*k.NumBus(); i++ {
		val, err := getLevel(3, i)
		if err != nil {
			log.Error(err.Error())
			return false, err
		}
		_levelCache.busLevelsBuff[i] = val
		_levelCache.busComp[i] = _levelCache.busLevelsBuff[i] == _levelCache.busLevels[i]
	}
	return !(allTrue(_levelCache.stripComp, (2*k.PhysIn)+(8*k.VirtIn)) && allTrue(_levelCache.busComp, 8*k.NumBus())), nil
}

func clear() {
	for {
		pdirty, _ := pdirty()
		mdirty, _ := mdirty()
		if !(pdirty || mdirty) {
			break
		}
	}
}

// getVMType returns the type of Voicemeeter, as a string
func getVMType() (string, error) {
	var type_ uint64
	res, _, _ := vmGetvmType.Call(
		uintptr(unsafe.Pointer(&type_)),
	)
	if int32(res) != 0 {
		err := fmt.Errorf("VBVMR_GetVoicemeeterType returned %d", int32(res))
		return "", err
	}
	vals := map[uint64]string{
		1: "basic",
		2: "banana",
		3: "potato",
	}
	return vals[type_], nil
}

// getParameterFloat gets the value of a float parameter
func getParameterFloat(name string) (float64, error) {
	if vmsync {
		time.Sleep(time.Duration(vmdelay) * time.Millisecond)
		clear()
	}
	var value float32
	b := append([]byte(name), 0)
	res, _, _ := vmGetParamFloat.Call(
		uintptr(unsafe.Pointer(&b[0])),
		uintptr(unsafe.Pointer(&value)),
	)
	if int32(res) != 0 {
		err := fmt.Errorf("VBVMR_GetParameterFloat returned %d", int32(res))
		return 0, err
	}
	return math.Round(float64(value)*10) / 10, nil
}

// setParameterFloat sets the value of a float parameter
func setParameterFloat(name string, value float64) error {
	b1 := append([]byte(name), 0)
	b2 := math.Float32bits(float32(value))
	res, _, _ := vmSetParamFloat.Call(
		uintptr(unsafe.Pointer(&b1[0])),
		uintptr(b2),
	)
	if int32(res) != 0 {
		err := fmt.Errorf("VBVMR_SetParameterFloat returned %d", int32(res))
		return err
	}
	return nil
}

// getParameterString gets the value of a string parameter
func getParameterString(name string) (string, error) {
	if vmsync {
		time.Sleep(time.Duration(vmdelay) * time.Millisecond)
		clear()
	}
	b1 := append([]byte(name), 0)
	var b2 [512]byte
	res, _, _ := vmGetParamString.Call(
		uintptr(unsafe.Pointer(&b1[0])),
		uintptr(unsafe.Pointer(&b2[0])),
	)
	if int32(res) != 0 {
		err := fmt.Errorf("VBVMR_GetParameterStringA returned %d", int32(res))
		return "", err
	}
	str := bytes.Trim(b2[:], "\x00")
	return string(str), nil
}

// setParameterString sets the value of a string parameter
func setParameterString(name, value string) error {
	b1 := append([]byte(name), 0)
	b2 := append([]byte(value), 0)
	res, _, _ := vmSetParamString.Call(
		uintptr(unsafe.Pointer(&b1[0])),
		uintptr(unsafe.Pointer(&b2[0])),
	)
	if int32(res) != 0 {
		err := fmt.Errorf("VBVMR_SetParameterStringA returned %d", int32(res))
		return err
	}
	return nil
}

// setParametersMulti sets multiple parameters with a script
func setParametersMulti(script string) error {
	b1 := append([]byte(script), 0)
	res, _, _ := vmSetParameters.Call(
		uintptr(unsafe.Pointer(&b1[0])),
	)
	if int32(res) != 0 {
		err := fmt.Errorf("VBVMR_SetParameters returned %d", int32(res))
		return err
	}
	return nil
}

// getMacroStatus gets a macrobutton value
func getMacroStatus(id, mode int) (float64, error) {
	if vmsync {
		time.Sleep(time.Duration(vmdelay) * time.Millisecond)
		clear()
	}
	var state float32
	res, _, _ := vmGetMacroStatus.Call(
		uintptr(id),
		uintptr(unsafe.Pointer(&state)),
		uintptr(mode),
	)
	if int32(res) != 0 {
		err := fmt.Errorf("VBVMR_MacroButton_GetStatus returned %d", int32(res))
		return 0, err
	}
	return float64(state), nil
}

// setMacroStatus sets a macrobutton value
func setMacroStatus(id, state, mode int) error {
	res, _, _ := vmSetMacroStatus.Call(
		uintptr(id),
		uintptr(state),
		uintptr(mode),
	)
	if int32(res) != 0 {
		err := fmt.Errorf("VBVMR_MacroButton_SetStatus returned %d", int32(res))
		return err
	}
	return nil
}

// getNumDevices returns the number of hardware input/output devices
func getNumDevices(dir string) uint64 {
	if strings.Compare(dir, "in") == 0 {
		res, _, _ := vmGetDevNumIn.Call()
		return uint64(res)
	} else {
		res, _, _ := vmGetDevNumOut.Call()
		return uint64(res)
	}
}

// getDeviceDescription returns name, driver type and hwid for a given device
func getDeviceDescription(i int, dir string) (string, uint64, string, error) {
	var t_ uint64
	var b1 [512]byte
	var b2 [512]byte
	if strings.Compare(dir, "in") == 0 {
		res, _, _ := vmGetDevDescIn.Call(
			uintptr(i),
			uintptr(unsafe.Pointer(&t_)),
			uintptr(unsafe.Pointer(&b1[0])),
			uintptr(unsafe.Pointer(&b2[0])),
		)
		if int32(res) != 0 {
			err := fmt.Errorf("VBVMR_Input_GetDeviceDescA returned %d", int32(res))
			return "", 0, "", err
		}
	} else {
		res, _, _ := vmGetDevDescOut.Call(
			uintptr(i),
			uintptr(unsafe.Pointer(&t_)),
			uintptr(unsafe.Pointer(&b1[0])),
			uintptr(unsafe.Pointer(&b2[0])),
		)
		if int32(res) != 0 {
			err := fmt.Errorf("VBVMR_Output_GetDeviceDescA returned %d", int32(res))
			return "", 0, "", err
		}
	}
	name := bytes.Trim(b1[:], "\x00")
	hwid := bytes.Trim(b2[:], "\x00")
	return string(name), t_, string(hwid), nil
}

// getLevel returns a single level value of type type_ for channel[i]
func getLevel(type_, i int) (float64, error) {
	var val float32
	res, _, _ := vmGetLevelFloat.Call(
		uintptr(type_),
		uintptr(i),
		uintptr(unsafe.Pointer(&val)),
	)
	if int32(res) != 0 {
		err := fmt.Errorf("VBVMR_GetLevel returned %d", int32(res))
		return 0, err
	}
	return float64(val), nil
}

// getMidiMessage gets midi channel, pitch and velocity for a single midi input
func getMidiMessage() (bool, error) {
	var midi = newMidi()
	var b1 [1024]byte
	res, _, _ := vmGetMidiMessage.Call(
		uintptr(unsafe.Pointer(&b1[0])),
		uintptr(1024),
	)
	x := int32(res)
	if x < 0 {
		if x == -2 {
			err := fmt.Errorf("VBVMR_GetMidiMessage returned %d", int32(res))
			if err != nil {
				log.Error(err)
				return false, err
			}
		}
	}
	msg := bytes.Trim(b1[:], "\x00")
	if len(msg) > 0 {
		for i := 0; i < len(msg)%3; i++ {
			msg = append(msg, 0)
		}

		for i := 0; i < len(msg); i += 3 {
			var ch = int(msg[i])
			var pitch = int(msg[i+1])
			var vel = int(msg[i+2])
			midi.channel = ch
			midi.current = pitch
			midi.cache[pitch] = vel
		}
	}
	return len(msg) > 0, nil
}
