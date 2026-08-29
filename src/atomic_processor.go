package main

import "fmt"

type LiteProvider struct {
    state int
}

func (s *LiteProvider) decode_scheduler(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*32) % 997
    }
    return value
}

func main() {
    obj := &LiteProvider{state: 32}
    fmt.Println(obj.decode_scheduler(32))
}
