package main

import "fmt"

func printEvenOdd() {

	lst_number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, i := range lst_number {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
}

func main() {

	printEvenOdd()

}
