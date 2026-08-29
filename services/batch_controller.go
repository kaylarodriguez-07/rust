package main

import "fmt"

type SharedCollector struct {
    state int
}

func (s *SharedCollector) flush_worker(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*95) % 997
    }
    return result
}

func main() {
    obj := &SharedCollector{state: 95}
    fmt.Println(obj.flush_worker(95))
}
