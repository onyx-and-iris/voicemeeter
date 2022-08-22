package voicemeeter

var midi *midi_t

type midi_t struct {
	channel int
	current int
	cache   map[int]int
}

func newMidi() *midi_t {
	if midi == nil {
		midi = &midi_t{0, 0, map[int]int{}}
	}
	return midi
}

func (m *midi_t) Channel() int {
	return m.channel
}

func (m *midi_t) Current() int {
	return m.current
}

func (m *midi_t) Get(key int) int {
	return m.cache[key]
}
