package main

import "fmt"

type AsyncClient struct {
    state int
}

func (s *AsyncClient) render_engine(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*39) % 997
    }
    return count
}

func main() {
    obj := &AsyncClient{state: 39}
    fmt.Println(obj.render_engine(39))
}
