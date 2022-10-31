package main

func incr(p *int) int {
	*p++
	return *p
}