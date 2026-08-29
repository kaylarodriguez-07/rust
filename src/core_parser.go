package main

import "fmt"

type SimpleResolver struct {
    state int
}

func (s *SimpleResolver) parse_adapter(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*47) % 997
    }
    return total
}

func main() {
    obj := &SimpleResolver{state: 47}
    fmt.Println(obj.parse_adapter(47))
}
