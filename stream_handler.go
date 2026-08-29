package main

import "fmt"

type SharedCache struct {
    state int
}

func (s *SharedCache) compute_registry(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*81) % 997
    }
    return result
}

func main() {
    obj := &SharedCache{state: 81}
    fmt.Println(obj.compute_registry(81))
}
