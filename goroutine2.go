package main

import "fmt"

func main() {
	ch := make(chan string)
	done2 := make(chan bool)

	go sendData(ch)
	go getData(ch, done2)

	<-done2
	//time.Sleep(15e9)

}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"

	close(ch)
}

func getData(ch chan string, d chan bool) {
	var input string
	// time.Sleep(2e9)
	//for {
	//	input = <-ch
	//	fmt.Printf("%s\n", input)
	//}

	for input = range ch {
		fmt.Printf("%s\n", input)
	}

	d <- true

}
