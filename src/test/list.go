// a list with head
//simple test pointer in go
//suppose the key is unique

package test

import "fmt"

type ListNode struct {
	key  int
	next *ListNode
}

func insert(head *ListNode, value int) {
	node := new(ListNode)
	node.key = value
	node.next = nil
	tmp_pointer := head

	for ; tmp_pointer.next != nil; tmp_pointer = tmp_pointer.next {
	}
	tmp_pointer.next = node
}

func delete(head *ListNode, value int) {
	parent_pointer := head
	tmp_pointer := head.next
	for ; tmp_pointer != nil && tmp_pointer.key != value; parent_pointer, tmp_pointer = tmp_pointer, tmp_pointer.next {
	}
	if tmp_pointer == nil {
		return
	} else {
		parent_pointer.next = tmp_pointer.next
	}
}

// no need: the gc does the work
func destroy(head *ListNode) {
}

func print_list(head *ListNode) {
	for tmp_pointer := head.next; tmp_pointer != nil; tmp_pointer = tmp_pointer.next {
		fmt.Printf("%d ", tmp_pointer.key)
	}
	fmt.Printf("\n")
}

func TestList() {
	var head *ListNode = new(ListNode)
	head.next = nil
	insert(head, 11)
	for i := 0; i <= 10; i += 1 {
		insert(head, i)
	}
	fmt.Println("list :")
	print_list(head)
	fmt.Println("after delete 9 :")
	delete(head, 9)
	print_list(head)
	fmt.Println("after delete 10 :")
	delete(head, 10)
	print_list(head)
	fmt.Println("try to delete 20 :")
	delete(head, 20)
	print_list(head)
}
