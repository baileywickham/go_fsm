package main

type fsm struct {
	current     string
	transitions []transition
}

type transition struct {
	name string
	from string
	to   []string
	callback
}

type callback interface{}

func fsm_init(inital string, states []string) fsm {
	m := fsm{current: inital}
	return m
}

func (m *fsm) add_transition(t transition) {
	m.transitions = append(m.transitions, t)

}

func (m *fsm) can(next string) bool {
	// needs to be rewritten with a map.
	for _, tran := range m.transitions {
		if tran.name == m.current {
			for _, possible := range tran.to {
				if next == possible {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	runner()
}
