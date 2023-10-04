package voicemeeter

// recorder represents the recorder
type recorder struct {
	iRemote
	outputs
}

// newRecorder returns an address to a recorder struct
func newRecorder() *recorder {
	o := newOutputs("recorder", 0)
	return &recorder{iRemote{"recorder", 0}, o}
}

// Play plays the file currently loaded into the recorder
func (r *recorder) Play() {
	r.setter_float("Play", 1)
}

// Stop stops the file currently playing
func (r *recorder) Stop() {
	r.setter_float("Stop", 0)
}

// Pause pauses the file currently playing
func (r *recorder) Pause() {
	r.setter_float("Pause", 1)
}

// Restart restarts the Voicemeeter audio engine
func (r *recorder) Replay() {
	r.setter_float("Replay", 1)
}

// Record records the current track playing
func (r *recorder) Record() {
	r.setter_float("Record", 1)
}

// Ff fast forwards the recorder
func (r *recorder) Ff() {
	r.setter_float("Ff", 1)
}

// Rew rewinds the recorder
func (r *recorder) Rew() {
	r.setter_float("Rew", 1)
}

// Loop enables loop play mode
func (r *recorder) Loop(val bool) {
	r.setter_bool("Mode.Loop", val)
}

// Gain returns the value of the Gain parameter
func (r *recorder) Gain() float64 {
	return r.getter_float("Gain")
}

// SetGain sets the value of the Gain parameter
func (r *recorder) SetGain(val float64) {
	r.setter_float("Gain", val)
}
