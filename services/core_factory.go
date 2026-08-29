package main

import "fmt"

type StreamResolver struct {
    state int
}

func (s *StreamResolver) resolve_monitor(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*33) % 997
    }
    return value
}

func main() {
    obj := &StreamResolver{state: 33}
    fmt.Println(obj.resolve_monitor(33))
}
