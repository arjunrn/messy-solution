package main

import (
	"fmt"
)

type Controller struct {
	status []ElStatus
}

func (c Controller) RegisterElevator(s ElStatus) {
	for i := 0; i < len(c.status); i++ {
		if c.status[i].id == s.id {
			fmt.Println("Already Registered")
			return
		}
	}
	c.status = append(c.status, s)
}

func (c Controller) run() {
	fmt.Printf("Running Controller")
}

func (c Controller) Elevators() []ElStatus {
	return c.status
}

func (c Controller) step() {

}

func (c Controller) registerCall(floor int, destination int) {

}
