package fsm

import (
	"errors"
)

type State struct {
	Name   string
	Before func()
	After  func()
	// event -> end state
	ToState map[string]string
}

type Fsm struct {
	current string
	states  map[string]State
}

var InvalidEventErr = errors.New("Cannot trigger event from here")

func InitFsm(inital string) Fsm {
	m := Fsm{
		current: inital,
		states:  make(map[string]State),
	}

	return m
}

func (m *Fsm) AddState(s State) {
	m.states[s.Name] = s
}

func (m *Fsm) SetState(state State) {
	m.current = state.Name
}

func (m *Fsm) Current() State {
	return m.states[m.current]
}

func (m *Fsm) Can(event string) bool {
	// needs to be rewritten with a map.
	currentState := m.Current()
	_, can := currentState.ToState[event]

	return can
}

func (m *Fsm) PossibleStates() []string {
	states := make([]string, 0)
	for _, state := range m.Current().ToState {
		states = append(states, state)
	}
	return states
}

func (m *Fsm) Event(event string) error {
	if !m.Can(event) {
		return InvalidEventErr
	}

	currentState := m.Current()
	newState := m.states[currentState.ToState[event]]

	if currentState.After != nil {
		currentState.After()
	}
	m.SetState(newState)
	if newState.Before != nil {
		newState.Before()
	}
	return nil

}
