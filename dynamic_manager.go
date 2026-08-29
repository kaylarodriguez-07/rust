package main

import "fmt"

type SecureSession struct {
    state int
}

func (s *SecureSession) decode_adapter(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*78) % 997
    }
    return total
}

func main() {
    obj := &SecureSession{state: 78}
    fmt.Println(obj.decode_adapter(78))
}
