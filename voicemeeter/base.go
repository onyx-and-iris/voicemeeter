package voicemeeter

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"syscall"
	"time"
	"unsafe"
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

	//vmGetLevelFloat = mod.NewProc("VBVMR_GetLevel")

	vmSetParamFloat  = mod.NewProc("VBVMR_SetParameterFloat")
	vmSetParameters  = mod.NewProc("VBVMR_SetParameters")
	vmSetParamString = mod.NewProc("VBVMR_SetParameterStringA")

	//vmGetDevNumOut  = mod.NewProc("VBVMR_Output_GetDeviceNumber")
	//vmGetDevDescOut = mod.NewProc("VBVMR_Output_GetDeviceDescA")
	//vmGetDevNumIn   = mod.NewProc("VBVMR_Input_GetDeviceNumber")
	//vmGetDevDescIn  = mod.NewProc("VBVMR_Input_GetDeviceDescA")

	vmMdirty         = mod.NewProc("VBVMR_MacroButton_IsDirty")
	vmGetMacroStatus = mod.NewProc("VBVMR_MacroButton_GetStatus")
	vmSetMacroStatus = mod.NewProc("VBVMR_MacroButton_SetStatus")
)

// login logs into the API,
// then attempts to launch Voicemeeter if it's not running.
func login(kind_id string) {
	res, _, _ := vmLogin.Call()
	if res == 1 {
		runVoicemeeter(kind_id)
		time.Sleep(time.Second)
	} else if res != 0 {
		err := fmt.Errorf("VBVMR_Login returned %d", res)
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Logged into API")
	sync()
}

// logout logs out of the API,
// delayed for 100ms to allow final operation to complete.
func logout() {
	time.Sleep(100 * time.Millisecond)
	res, _, _ := vmLogout.Call()
	if res != 0 {
		err := fmt.Errorf("VBVMR_Logout returned %d", res)
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Logged out of API")
}

// runVoicemeeter attempts to launch a Voicemeeter GUI of a kind.
func runVoicemeeter(kind_id string) {
	vals := map[string]uint64{
		"basic":  1,
		"banana": 2,
		"potato": 3,
	}
	res, _, _ := vmRunvm.Call(uintptr(vals[kind_id]))
	if res != 0 {
		err := fmt.Errorf("VBVMR_RunVoicemeeter returned %d", res)
		fmt.Println(err)
		os.Exit(1)
	}
}

// getVersion returns the version of Voicemeeter as a string
func getVersion() string {
	var ver uint64
	res, _, _ := vmGetvmVersion.Call(uintptr(unsafe.Pointer(&ver)))
	if res != 0 {
		err := fmt.Errorf("VBVMR_GetVoicemeeterVersion returned %d", res)
		fmt.Println(err)
		os.Exit(1)
	}
	v1 := (ver & 0xFF000000) >> 24
	v2 := (ver & 0x00FF0000) >> 16
	v3 := (ver & 0x0000FF00) >> 8
	v4 := ver & 0x000000FF
	return fmt.Sprintf("%d.%d.%d.%d", v1, v2, v3, v4)
}

// pdirty returns true iff a parameter value has changed
func pdirty() bool {
	res, _, _ := vmPdirty.Call()
	return int(res) == 1
}

// mdirty returns true iff a macrobutton value has changed
func mdirty() bool {
	res, _, _ := vmMdirty.Call()
	return int(res) == 1
}

func sync() {
	time.Sleep(5 * time.Millisecond)
	for pdirty() || mdirty() {
	}
}

// getVMType returns the type of Voicemeeter, as a string
func getVMType() string {
	var type_ uint64
	res, _, _ := vmGetvmType.Call(
		uintptr(unsafe.Pointer(&type_)),
	)
	if res != 0 {
		err := fmt.Errorf("VBVMR_GetVoicemeeterType returned %d", res)
		fmt.Println(err)
		os.Exit(1)
	}
	vals := map[uint64]string{
		1: "basic",
		2: "banana",
		3: "potato",
	}
	return vals[type_]
}

// getParameterFloat gets the value of a float parameter
func getParameterFloat(name string) float64 {
	var value float32
	b := append([]byte(name), 0)
	res, _, _ := vmGetParamFloat.Call(
		uintptr(unsafe.Pointer(&b[0])),
		uintptr(unsafe.Pointer(&value)),
	)
	if res != 0 {
		err := fmt.Errorf("VBVMR_GetParameterFloat returned %d", res)
		fmt.Println(err)
		os.Exit(1)
	}
	return math.Round(float64(value)*10) / 10
}

// getParameterFloat sets the value of a float parameter
func setParameterFloat(name string, value float32) {
	b1 := append([]byte(name), 0)
	b2 := math.Float32bits(value)
	res, _, _ := vmSetParamFloat.Call(
		uintptr(unsafe.Pointer(&b1[0])),
		uintptr(b2),
	)
	if res != 0 {
		err := fmt.Errorf("VBVMR_SetParameterFloat returned %d", res)
		fmt.Println(err)
		os.Exit(1)
	}
	sync()
}

// getParameterString gets the value of a string parameter
func getParameterString(name string) string {
	b1 := append([]byte(name), 0)
	var b2 [512]byte
	res, _, _ := vmGetParamString.Call(
		uintptr(unsafe.Pointer(&b1[0])),
		uintptr(unsafe.Pointer(&b2[0])),
	)
	if res != 0 {
		err := fmt.Errorf("VBVMR_GetParameterStringA returned %d", res)
		fmt.Println(err)
		os.Exit(1)
	}
	str := bytes.Trim(b2[:], "\x00")
	return string(str)
}

// getParameterString sets the value of a string parameter
func setParameterString(name, value string) {
	b1 := append([]byte(name), 0)
	b2 := append([]byte(value), 0)
	res, _, _ := vmSetParamString.Call(
		uintptr(unsafe.Pointer(&b1[0])),
		uintptr(unsafe.Pointer(&b2[0])),
	)
	if res != 0 {
		err := fmt.Errorf("VBVMR_SetParameterStringA returned %d", res)
		fmt.Println(err)
		os.Exit(1)
	}
	sync()
}

// setParametersMulti sets multiple parameters with a script
func setParametersMulti(script string) {
	b1 := append([]byte(script), 0)
	res, _, _ := vmSetParameters.Call(
		uintptr(unsafe.Pointer(&b1[0])),
	)
	if res != 0 {
		err := fmt.Errorf("VBVMR_SetParameters returned %d", res)
		fmt.Println(err)
		os.Exit(1)
	}
}

// getMacroStatus gets a macrobutton value
func getMacroStatus(id, mode int) float32 {
	var state float32
	res, _, _ := vmGetMacroStatus.Call(
		uintptr(id),
		uintptr(unsafe.Pointer(&state)),
		uintptr(mode),
	)
	if res != 0 {
		err := fmt.Errorf("VBVMR_MacroButton_GetStatus returned %d", res)
		fmt.Println(err)
		os.Exit(1)
	}
	return state
}

// setMacroStatus sets a macrobutton value
func setMacroStatus(id, state, mode int) {
	res, _, _ := vmSetMacroStatus.Call(
		uintptr(id),
		uintptr(state),
		uintptr(mode),
	)
	if res != 0 {
		err := fmt.Errorf("VBVMR_MacroButton_SetStatus returned %d", res)
		fmt.Println(err)
		os.Exit(1)
	}
	time.Sleep(30 * time.Millisecond)
	sync()
}

/*
// getLevel returns a single level value of type type_ for channel[i]
func getLevel(type_, i int) float32 {
	var val float32
	res, _, _ := getLevelFloat.Call(
		uintptr(type_),
		uintptr(i),
		uintptr(unsafe.Pointer(&val)),
	)
	if res != 0 {
		err := fmt.Errorf("VBVMR_GetLevel returned %d", res)
		fmt.Println(err)
		os.Exit(1)
	}
	return val
}
*/
