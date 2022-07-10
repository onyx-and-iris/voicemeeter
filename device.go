package voicemeeter

type devDesc struct {
	Name, Type, Hwid string
}

type device struct {
}

func newDevice() *device {
	return &device{}
}

// Ins returns the total number of physical input devices
func (d *device) Ins() int {
	return int(getNumDevices("in"))
}

// Ins returns the total number of physical input devices
func (d *device) Outs() int {
	return int(getNumDevices("out"))
}

func (d *device) Input(i int) devDesc {
	n, t_, id := getDeviceDescription(i, "in")
	vals := map[uint64]string{
		1: "mme",
		3: "wdm",
		4: "ks",
		5: "asio",
	}
	return devDesc{Name: n, Type: vals[t_], Hwid: id}
}

func (d *device) Output(i int) devDesc {
	n, t_, id := getDeviceDescription(i, "out")
	vals := map[uint64]string{
		1: "mme",
		3: "wdm",
		4: "ks",
		5: "asio",
	}
	return devDesc{Name: n, Type: vals[t_], Hwid: id}
}
