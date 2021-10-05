package main

import (
	"fmt"

	fsm "github.com/baileywickham/go_fsm"
)

func main() {
	s0 := fsm.State{
		Name:   "unlocked",
		Before: func() { fmt.Println("The door is now unlocked") },
		ToState: map[string]string{
			"push": "locked",
			"coin": "unlocked"}}
	s1 := fsm.State{
		Name:   "locked",
		Before: func() { fmt.Println("The door is now locked") },
		ToState: map[string]string{
			"push": "locked",
			"coin": "unlocked"}}

	m := fsm.InitFsm("locked")

	m.AddState(s0)
	m.AddState(s1)

	println("Possible states to transition to: ")
	for _, tran := range m.PossibleStates() {
		println(tran)
	}

	m.Event("coin")
	m.Event("coin")
	m.Event("push")
	m.Event("coin")
}
