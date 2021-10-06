package main

import (
	"fmt"

	fsm "github.com/baileywickham/go_fsm"
)

func main() {
	s0 := fsm.State{
		Name:   "unlocked",
		Before: func() { fmt.Println("The door is now unlocked") },
		ToState: map[string]fsm.Transition{
			"push": {
				To:       "locked",
				Callback: func() { fmt.Println("pushed an unlocked door") }},
			"coin": {
				To:       "unlocked",
				Callback: func() { fmt.Println("insert a coin into an unlocked door") }}}}
	s1 := fsm.State{
		Name:   "locked",
		Before: func() { fmt.Println("The door is now locked") },
		ToState: map[string]fsm.Transition{
			"push": {
				To:       "locked",
				Callback: func() { fmt.Println("pushed a locked door") }},
			"coin": {
				To:       "unlocked",
				Callback: func() { fmt.Println("insert a coint into a locked door") }}}}

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
