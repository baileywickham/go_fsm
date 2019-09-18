package main

import (
	"fmt"

	fsm "github.com/baileywickham/go_fsm"
)

func main() {
	t1 := fsm.Transition{
		Name:     "push",
		From:     "unlocked",
		To:       "locked",
		Callback: func() { fmt.Println("You entered the gate") },
	}

	t2 := fsm.Transition{
		Name:     "push",
		From:     "locked",
		To:       "locked",
		Callback: func() { fmt.Println("You pushed but it was locked") },
	}
	t3 := fsm.Transition{
		Name:     "coin",
		From:     "locked",
		To:       "unlocked",
		Callback: func() { fmt.Println("You inserted a coin and unlocked the gate") },
	}
	t4 := fsm.Transition{
		Name:     "coin",
		From:     "unlocked",
		To:       "unlocked",
		Callback: func() { fmt.Println("You inserted a coin but the gate was already unlocked") },
	}
	t := []fsm.Transition{t1, t2, t3, t4}
	m := fsm.Fsm_init("locked")

	for _, tran := range t {
		println("adding transition")
		m.Add_transition(tran)
	}

	println("Possible states to transition to: ")
	for _, tran := range m.Possible_transitions() {
		println(tran)
	}
}
