package utils

import "fmt"

func PrintArray(data []int) {
	length := len(data)
	for i, value := range data {
		if i == length-1 {
			fmt.Printf("%d\n", value)
		} else {
			fmt.Printf("%d,", value)
		}
	}

}
