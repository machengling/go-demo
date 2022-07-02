package main

import "sort"

type Number interface {
	int | int32 | int64 | float64 | float32
}

type arrayList[K Number] struct {
	data  *[]K
	sortL bool
}

func NewArrayList[K Number]() arrayList[K] {
	list := arrayList[K]{}
	list.data = &[]K{}
	return list
}

func (list *arrayList[K]) Less(i, j int) bool {
	if list.sortL {
		return (*list.data)[i] < (*list.data)[j]
	}
	return (*list.data)[i] > (*list.data)[j]
}

func (list *arrayList[K]) Swap(i, j int) {
	(*list.data)[i], (*list.data)[j] = (*list.data)[j], (*list.data)[i]

}

func (list *arrayList[K]) Len() int {
	return len(*list.data)
}

func (list *arrayList[K]) add(item K) {
	*list.data = append(*list.data, item)
}

func (list *arrayList[K]) sortLess() {
	list.sortL = true
	sort.Sort(list)
}

func (list *arrayList[K]) sortBig() {
	list.sortL = false
	sort.Sort(list)
}
