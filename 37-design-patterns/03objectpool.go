package main

import (
    "fmt"
    "log"
    "strconv"
    "sync"
)

type poolResource interface {
  // Unique identifier for a pool resource.
  getID() string
}

/*
 * The following struct and the methods on it 
 * will be used to handle resource pooling.
 */

type pool struct {
  idle     []poolResource
  active   []poolResource
  capacity int
  mutex    *sync.Mutex
}

func initPool(poolObjects []poolResource) (*pool, error) {
  if len(poolObjects) == 0 {
    return nil, fmt.Errorf("Cannot create a pool of 0 length")
  }
  active := make([]poolResource, 0)
  pool := &pool{
    idle:     poolObjects,
    active:   active,
    capacity: len(poolObjects),
    mutex:   new(sync.Mutex),
  }
  return pool, nil
}

func (p *pool) acquire() (poolResource, error) {
  p.mutex.Lock()
  defer p.mutex.Unlock()
  if len(p.idle) == 0 {
    return nil, fmt.Errorf("No pool resource free. Please request after sometime")
  }
  obj := p.idle[0]
  p.idle = p.idle[1:]
  p.active = append(p.active, obj)
  fmt.Printf("Acquired Pool Resource with ID: %s\n", obj.getID())
  return obj, nil
}

func (p *pool) release(target poolResource) error {
  p.mutex.Lock()
  defer p.mutex.Unlock()
  err := p.remove(target)
  if err != nil {
    return err
  }
  p.idle = append(p.idle, target)
  fmt.Printf("Released Pool Resource with ID: %s\n", target.getID())
  return nil
}

func (p *pool) remove(target poolResource) error {
  currentActiveLength := len(p.active)
  for i, obj := range p.active {
    if obj.getID() == target.getID() {
      p.active[currentActiveLength-1], p.active[i] = p.active[i], p.active[currentActiveLength-1]
      p.active = p.active[:currentActiveLength-1]
      return nil
    }
  }
  return fmt.Errorf("Target pool resource doesn't belong to the pool")
}

type connection struct {
  id string
}

func (c *connection) getID() string {
  return c.id
}

func main() {
  connections := make([]poolResource, 0)
  for i := 0; i < 3; i++ {
    c := &connection{id: strconv.Itoa(i)}
    connections = append(connections, c)
  }
  pool, err := initPool(connections)
  if err != nil {
    log.Fatalf("Init Pool Error: %s", err)
  }
  firstConn, err := pool.acquire()
  if err != nil {
    log.Fatalf("Pool acquire Error: %s", err)
  }
  secondConn, err := pool.acquire()
  if err != nil {
    log.Fatalf("Pool acquire Error: %s", err)
  }
  pool.release(firstConn)
  pool.release(secondConn)
}