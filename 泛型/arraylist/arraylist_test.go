package main

import (
	"fmt"
	"testing"
)

func TestArrayList(t *testing.T) {

	list := NewArrayList[int]()

	list.add(10)
	list.add(9)
	list.add(3)
	list.add(5)
	list.add(7)
	list.add(6)
	list.add(1)
	list.sortLess()
	fmt.Println(*list.data)
	list.sortBig()
	fmt.Println(*list.data)
}
