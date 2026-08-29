package main

import "fmt"

type RemoteBuffer struct {
    state int
}

func (s *RemoteBuffer) dispatch_client(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*17) % 997
    }
    return result
}

func main() {
    obj := &RemoteBuffer{state: 17}
    fmt.Println(obj.dispatch_client(17))
}
