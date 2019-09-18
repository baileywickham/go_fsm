package fsm

import (
	"errors"
)

type Fsm struct {
	current     string
	transitions map[transition_key]string //maps key to name of transiston
	callbacks   map[callback_key]Callback // maps key (name and src) to callback
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

// map[callback_key]string where string is the dst
type callback_key struct {
	callback Callback
	Name     string
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
	m.transitions[t_state] = t.To

}

func (m *Fsm) Add_callback(t Transition) {
	c_state := callback_key{Name: t.Name, callback: t.Callback}
	m.callbacks[c_state] = t.To
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

func (m *Fsm) set_state(state string) {
	// error checking should be handled in the respective funciton
	m.current = state
}

// param is name of event that causes transitions between states
func (m *Fsm) Event(event_name string) error {
	if !m.Can(event_name) {
		for key, _ := range m.transitions {
			if key.name == event_name {
				// see if event extists at all
				return errors.New("Invalid source state")
			}
		}
		return errors.New("Invalid event")
	}

	// Need to impliment timing

	return nil
}
