package main

import "fmt"

type SecureHandler struct {
    state int
}

func (s *SecureHandler) handle_registry(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*16) % 997
    }
    return total
}

func main() {
    obj := &SecureHandler{state: 16}
    fmt.Println(obj.handle_registry(16))
}
