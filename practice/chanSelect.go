package main

import "fmt"
import "math/rand"

func main() {
	intChannels := [3]chan int{
		make(chan int, 1),
		make(chan int, 1),
		make(chan int, 1),
	}

	index := rand.Intn(3)
	fmt.Printf("The index: %d\n", index)
	intChannels[index] <- index

	select {
	case <-intChannels[0]:
		fmt.Printf("The first candidate case is selected")
	case <-intChannels[1]:
		fmt.Printf("The second candidate case is selected")
	case elem := <-intChannels[2]:
		fmt.Printf("The third candidate case is selected, the element is %d.\n", elem)
	default:
		fmt.Printf("No candidate case is selected!")
	}
}
