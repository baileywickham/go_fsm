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

// map[transition_key]string where string is the name of the transition
type transition_key struct {
	to   string
	from string
}

type Callback interface{}

func Fsm_init(inital string) Fsm {
	// shuold this return a pointer? Probably
	m := Fsm{
		current:     inital,
		transitions: make(map[transition_key]string),
	}

	return m
}

func (m *Fsm) Add_transition(t Transition) {
	t_state := transition_key{from: t.From, to: t.To}
	m.transitions[t_state] = t.Name

}

func (m *Fsm) Can(next string) bool {
	// needs to be rewritten with a map.
	t_state := transition_key{from: m.current, to: next}
	_, can := m.transitions[t_state]
	return can
}
func (m *Fsm) Possible_transitions() []string {
	var transitions []string
	for t := range m.transitions {
		if t.from == m.current {
			// return names of transitions possible by checking transition map
			transitions = append(transitions, m.transitions[t])
		}
	}
	return transitions

}

func main() {

}
