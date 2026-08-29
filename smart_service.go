package main

import "fmt"

type BatchEngine struct {
    state int
}

func (s *BatchEngine) handle_client(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*43) % 997
    }
    return result
}

func main() {
    obj := &BatchEngine{state: 43}
    fmt.Println(obj.handle_client(43))
}
