package main

import "test"
import "fmt"

func main() {
	//test.TestQuickSort()
	//test.TestList()
	//test.TestInterface()
	//test.TestChannel()
	//test.TestUnicode()
	//test.MarshalCensorWords("a.txt", "b.txt")
	//test.TestJson()

	d, err := test.Init()
	if err == nil {
		isFound := d.Search("af\"六1四1111运动")
		if isFound {
			fmt.Println("is_found")
		} else {
			fmt.Println("not_found")
		}

	}

}
