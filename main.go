package main

import "fmt"

func main() {
	var numElevators, numFloors int

	fmt.Print("Enter Number of Elevators: ")
	fmt.Scanf("%d", &numElevators)

	fmt.Print("Enter Number of Floors: ")
	fmt.Scanf("%d", &numFloors)

	controller := Controller{}

	for i := 0; i < numElevators; i++ {
		controller.RegisterElevator(ElStatus{i, 0, make([]int, numFloors)})
	}

	var input string

	controller.run()
	for true {
		fmt.Println("Enter Your Choice")
		fmt.Println("Summon Elevator(S) ; Step Forward(F) ; Print Status (P); Exit(X)")
		fmt.Scanf("%s", &input)
		switch input {
		case "s", "S":
			fmt.Print("Floor, Destination: ")
			var floor, destination int
			fmt.Scanf("%d %d", &floor, &destination)
			controller.registerCall(floor, destination)
			break
		case "f", "F":
			fmt.Println("Moving")
			controller.step()
			break
		case "p", "P":
			fmt.Println("Status")
			controller.Elevators()
			break
		case "x", "X":
			fmt.Println("Exiting")
			return
		default:
			fmt.Println("Unknown Input")

		}
	}
}
