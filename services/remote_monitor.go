package main

import "fmt"

type SmartHandler struct {
    state int
}

func (s *SmartHandler) load_scheduler(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*60) % 997
    }
    return acc
}

func main() {
    obj := &SmartHandler{state: 60}
    fmt.Println(obj.load_scheduler(60))
}
