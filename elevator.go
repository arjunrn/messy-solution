package main

import (
	"bytes"
	"math"
	"sort"
	"strconv"
)

type ElStatus struct {
	id           int
	floor        int
	destinations []int
	direction    int
	floors       int
}

func (e ElStatus) getDestinations() string {
	var buff bytes.Buffer
	for i := 0; i < len(e.destinations); i++ {
		buff.WriteString(strconv.Itoa(e.destinations[i]))
	}

	return buff.String()
}

func (e *ElStatus) addDestination(summonFloor, destinationFloor int) {
	if summonFloor != e.floor {
		e.destinations = append(e.destinations, summonFloor)
	}

	e.destinations = append(e.destinations, destinationFloor)
	sort.Ints(e.destinations)

	if e.direction == 0 {
		if summonFloor > e.floor {
			e.direction = 1
		} else {
			e.direction = -1
		}
	}
}

func (e *ElStatus) step() {
	var destinations []int
	e.floor = e.floor + e.direction

	for i := 0; i < len(e.destinations); i++ {
		if e.floor != e.destinations[i] {
			destinations = append(destinations, e.destinations[i])
		}
	}

	highest := -1
	lowest := e.floors + 1
	for i := 0; i < len(e.destinations); i++ {
		if e.destinations[i] > highest {
			highest = e.destinations[i]
		}
		if e.destinations[i] < lowest {
			lowest = e.destinations[i]
		}
	}

	if e.direction == 1 && highest != -1 && highest < e.floor {
		e.direction = -1
	}

	if e.direction == -1 && lowest != e.floors+1 && lowest > e.floor {
		e.direction = 1
	}

	e.destinations = destinations
	if len(e.destinations) == 0 {
		e.direction = 0
	}
}

func (e ElStatus) effort(summonFloor int, destination int) int {
	baseEffort := summonFloor - destination
	if baseEffort < 0 {
		baseEffort = -baseEffort
	}
	if e.direction == 1 {
		if summonFloor > e.floor {
			return (summonFloor - e.floor + baseEffort)
		} else {
			return (e.destinations[len(e.destinations)-1] - summonFloor) + baseEffort
		}
	} else if e.direction == -1 {
		if summonFloor < e.floor {
			return (e.floor - summonFloor + baseEffort)
		} else {
			return (summonFloor - e.destinations[0] + baseEffort)
		}
	} else {
		return int(math.Abs(float64(e.floor-summonFloor))) + baseEffort
	}
}
