package main

import "fmt"

func main() {

	var balance = []int32{9, 3, 6, 1, 2, 8}

	insertSorter(balance)
	fmt.Println(balance)
}

func insertSorter(sorterArray []int32) {
	for i := 1; i < len(sorterArray); i++ {
		if sorterArray[i-1] > sorterArray[i] {
			temp := sorterArray[i]
			for j := i; j > 0 && sorterArray[j-1] > sorterArray[j]; j-- {
				sorterArray[j] = sorterArray[j-1]
				sorterArray[j-1] = temp
			}
		}
	}
}
