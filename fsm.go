package fsm

type Fsm struct {
	current     string
	transitions map[transition_key]string //maps key to name of transiston
}

// An event
type Transition struct {
	Name string
	From string
	To   string
	Callback
}

// map[transition_key]string where string is the dst
type transition_key struct {
	name    string
	current string
}

type callback_key struct {
	Callback
	Name string
}

type Callback interface{}

func Fsm_init(inital string) Fsm {
	// shuold this return a pointer? Probably
	m := Fsm{
		current: inital,
		// maps key to dst
		transitions: make(map[transition_key]string),
	}

	return m
}

func (m *Fsm) Add_transition(t Transition) {
	t_state := transition_key{name: t.Name, current: t.From}
	m.transitions[t_state] = t.dst

}

func (m *Fsm) Can(name string) bool {
	// needs to be rewritten with a map.
	t_state := transition_key{name: name, current: m.current}
	_, can := m.transitions[t_state]
	return can
}

func (m *Fsm) Possible_transitions() []string {
	var transitions []string
	for key, value := range m.transitions {
		if key.current == m.current {
			transitions = append(transitions, value)
		}
	}
	return transitions

}

// param is name of event that causes transitions between states
func (m *Fsm) Event(next_state string) error {
	return nil
}
