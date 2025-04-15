package main

import (
	"fmt"
)

type DynamicArray struct {
	Capacity, Length int
	Arr              []int
}

func NewDynamicArray() *DynamicArray {
	return &DynamicArray{
		Capacity: 2,
		Length:   0,
		Arr:      make([]int, 2),
	}
}

func (d *DynamicArray) Pushback(n int) {
	if d.Length == d.Capacity {
		d.Resize()
	}

	d.Arr[d.Length] = n
	d.Length++
}

func (d *DynamicArray) Resize() {
	d.Capacity = d.Capacity * 2
	newArr := make([]int, d.Capacity)

	copy(newArr, d.Arr)
	d.Arr = newArr
}

func (d *DynamicArray) Popback() {
	if d.Length > 0 {
		d.Length--
	}
}

func (d *DynamicArray) Get(i int) int {
	if i < d.Length {
		return d.Arr[i]
	}
	return -1
}

func (d *DynamicArray) Insert(i, n int) {
	if i < d.Length {
		d.Arr[i] = n
	}
}

func (d *DynamicArray) Print() {
	for i := 0; i < d.Length; i++ {
		fmt.Println(d.Arr[i])
	}
}
