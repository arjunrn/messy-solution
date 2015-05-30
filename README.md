## Solution

###Assumptions
1. The unit of time is the time for an elevator to move a single floor.
2. Moving up or down takes the same amount of time.
3. All elevators take the same amount to time for a unit.


### Instruction to Run
`go run *.go`

Enter the number of floors in the building the number of elevators

	Enter Number of Elevators: 5
	Enter Number of Floors: 20

Then run the simutation by stepping forward or submitting a request for the elevator

	Enter Your Choice
	Summon Elevator(S) ; Step Forward(F) ; Print Status (P); Exit(X)
	s
	Floor, Destination: 13 18
	The elevator selected to service the call is: 0
	It has to travel 18 floors.
	Enter Your Choice
	Summon Elevator(S) ; Step Forward(F) ; Print Status (P); Exit(X)
	f
	Moving
	Enter Your Choice
	Summon Elevator(S) ; Step Forward(F) ; Print Status 	(P); Exit(X)
	f
	Moving
	
## Lessons learnt
There should a 2 queues for the elevator. One queue is where it has to pickup passengers and the other where is has to drop them off. When the elevator picks up a passenger it should prompt for input. Unfortunately, this solution does not have this. As a result even the lift does not travel in the opposite direction. For example if the lift is on floor 10 and it has to pickup a passenger at floor 6 and deposit them at floor 8 it will only go from floor 10 to floor 6 stopping at floor 8 on the way.
