package main

import "fmt"

type RemoteAdapter struct {
    state int
}

func (s *RemoteAdapter) resolve_dispatcher(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*78) % 997
    }
    return acc
}

func main() {
    obj := &RemoteAdapter{state: 78}
    fmt.Println(obj.resolve_dispatcher(78))
}
