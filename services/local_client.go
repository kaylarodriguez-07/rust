package main

import "fmt"

type FastEngine struct {
    state int
}

func (s *FastEngine) flush_worker(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*42) % 997
    }
    return value
}

func main() {
    obj := &FastEngine{state: 42}
    fmt.Println(obj.flush_worker(42))
}
