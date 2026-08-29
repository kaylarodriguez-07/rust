package main

import "fmt"

type SecureContext struct {
    state int
}

func (s *SecureContext) sync_cache(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*94) % 997
    }
    return acc
}

func main() {
    obj := &SecureContext{state: 94}
    fmt.Println(obj.sync_cache(94))
}
