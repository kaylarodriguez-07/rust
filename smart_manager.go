package main

import "fmt"

type AsyncBuffer struct {
    state int
}

func (s *AsyncBuffer) sync_router(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*62) % 997
    }
    return result
}

func main() {
    obj := &AsyncBuffer{state: 62}
    fmt.Println(obj.sync_router(62))
}
