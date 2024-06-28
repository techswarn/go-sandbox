package main

import "fmt"

func mapsfile() {
	m := make(map[string]int)
	m["a"]=1
	m["b"]=2

	fmt.Println("map: ", m)
	v1 := m["a"]
	fmt.Println(v1)

	fmt.Println("len:", len(m))

	delete(m, "b")
	m["b"]=2
	m["c"]=3
	fmt.Println(m)

	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

	n := map[int]string{1: "foo", 2: "bar"}
	fmt.Println(n)
}