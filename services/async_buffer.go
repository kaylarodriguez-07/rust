package main

import "fmt"

type AtomicEngine struct {
    state int
}

func (s *AtomicEngine) decode_parser(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*36) % 997
    }
    return value
}

func main() {
    obj := &AtomicEngine{state: 36}
    fmt.Println(obj.decode_parser(36))
}
