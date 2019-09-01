package main

import (
	"fmt"
)

func runner() {
	t1 := transition{
		name:     "push",
		from:     "unlocked",
		to:       []string{"locked"},
		callback: func() { fmt.Println("You entered the gate") },
	}

	t2 := transition{
		name:     "push",
		from:     "locked",
		to:       []string{"locked"},
		callback: func() { fmt.Println("You pushed but it was locked") },
	}
	t3 := transition{
		name:     "coin",
		from:     "locked",
		to:       []string{"unlocked"},
		callback: func() { fmt.Println("You inserted a coin and unlocked the gate") },
	}
	t4 := transition{
		name:     "coin",
		from:     "unlocked",
		to:       []string{"unlocked"},
		callback: func() { fmt.Println("You inserted a coin but the gate was already unlocked") },
	}
	t := []transition{t1, t2, t3, t4}
	m := fsm_init("locked", []string{"locked", "unlocked"})

	for _, tran := range t {
		println("adding transition")
		m.add_transition(tran)
	}

}
