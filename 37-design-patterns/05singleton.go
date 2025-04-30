package main

import "sync"
import "fmt"

type mySingletonResource struct {
    // Singleton data here
}

/*
 * Feel free to add functions on your resource.
 *
 * func (res *mySingletonResource) doSomething() {
 *   ...  
 * }
 * 
 */ 

var instance *mySingletonResource
var once sync.Once

func getInstance() *mySingletonResource {
    once.Do(func() {
        instance = &mySingletonResource{}
        fmt.Println("Creating fresh resource", &instance)
    })
    fmt.Println("Returning resource", &instance)
    return instance
}

func main() {
  for i := 0; i < 5; i++ {
    getInstance()
  }
}