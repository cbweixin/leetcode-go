package channel_sample

import "fmt"

var intChan1 chan int
var intChan2 chan int
var channels = []chan int{intChan1, intChan2}

var numbers = []int{1, 2, 3, 4, 5}

func Select_test() {
	// select's execution order, from left to right, from up to bottom
	// intChan1 and intChan2 are not initialized yet, so it would be blocked
	// but still 2 case related code got evaulated. and finally the default branch has bee selected
	select {
	case getChan(0) <- getNumber(0):
		fmt.Println("The 1th case is selected.")
	case getChan(1) <- getNumber(1):
		fmt.Println("The 2nd case is selected.")
	default:
		fmt.Println("Default!")
	}
}

func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i]
}

func getChan(i int) chan int {
	fmt.Printf("channels[%d]\n", i)
	return channels[i]
}

/**
channels[0]
numbers[0]
channels[1]
numbers[1]
Default!
*/
