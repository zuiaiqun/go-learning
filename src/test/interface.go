package test

import "fmt"

type Human struct {
	name string
	age  int
}

type Student struct {
	school string
	loan   float32
}

type Man interface {
	TryModify(newname string)
}

// implements the "Man" interface
func (h *Human) TryModify(newname string) {
	fmt.Printf("try modify human from '%s' to '%s'\n", h.name, newname)
	//h.name = newname  //also works
	(*h).name = newname
}

// implements the "Man" interface
func (h Student) TryModify(newname string) {
	fmt.Printf("try modify student from '%s' to '%s'\n", h.school, newname)
	h.school = newname
}

func TestInterface() {
	stu := Student{school: "sysu", loan: 0.1}
	human := Human{"sysu", 5}
	test_list := make([]Man, 2)
	test_list[0], test_list[1] = stu, &human
	for _, value := range test_list {
		value.TryModify("HelloWorld")
	}
	fmt.Printf("after modify student: '%s' \n", stu.school)
	fmt.Printf("after modify human: '%s' \n", human.name)
}
