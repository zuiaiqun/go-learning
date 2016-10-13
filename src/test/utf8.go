package test

import "fmt"

func TestUnicode() {
	str := "你好a吗"
	for i, v := range str {
		fmt.Printf("%#U starts at byte position %d %d\n", v, i, int(v))
	}
	new := []rune(str)
	for i, v := range new {
		fmt.Printf("%#U starts at byte position %d %d\n", v, i, int(v))
	}
}
