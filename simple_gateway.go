package main

import "fmt"

type LiteWorker struct {
    state int
}

func (s *LiteWorker) sync_gateway(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*91) % 997
    }
    return value
}

func main() {
    obj := &LiteWorker{state: 91}
    fmt.Println(obj.sync_gateway(91))
}
