package main

import (
	"fmt"
)

type Controller struct {
	floors int
	status []ElStatus
}

func (c *Controller) RegisterElevator(s ElStatus) {
	for i := 0; i < len(c.status); i++ {
		if c.status[i].id == s.id {
			fmt.Println("Already Registered")
			return
		}
	}
	c.status = append(c.status, s)
}

func (c Controller) Elevators() {
	fmt.Println("Elevator   |   Floor   |   Goal")
	fmt.Println("-----------------------------------")
	for i := 0; i < len(c.status); i++ {
		fmt.Printf("%11d|%11d|%v\n", c.status[i].id, c.status[i].floor, c.status[i].destinations)
	}
}

func (c *Controller) step() {
	for i := 0; i < len(c.status); i++ {
		c.status[i].step()
	}
}

func (c *Controller) registerCall(floor int, destination int) {
	if floor == destination {
		fmt.Println("Stop messing with the elevator")
		return
	}

	if floor > c.floors || floor < 0 {
		fmt.Println("Invalid Floor Entry")
		return
	}

	if destination > c.floors || destination < 0 {
		fmt.Println("Invalid Destination Entry")
	}

	selectedEl := -1
	distance := c.floors + 1
	for i := 0; i < len(c.status); i++ {
		el := c.status[i]
		if el.effort(floor, destination) < distance {
			selectedEl = i
			distance = el.effort(floor, destination)
		}
	}

	fmt.Printf("The elevator selected to service the call is: %d\n", c.status[selectedEl].id)
	fmt.Printf("It has to travel %d floors.\n", distance)

	c.status[selectedEl].addDestination(floor, destination)
}
