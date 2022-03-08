package main

import (
	"fmt"
	"github.com/lesismal/nbio"
)

type GpsTrackerConnection struct {
	conn *nbio.Conn
}

func main() {
	g := nbio.NewGopher(nbio.Config{
		Network:            "tcp",
		Addrs:              []string{":8888"},
		MaxWriteBufferSize: 6 * 1024 * 1024,
	})
	g.OnOpen(func(c *nbio.Conn) {

	})
	g.OnData(func(c *nbio.Conn, data []byte) {
		var prev byte
		c.ReadBuffer = append(c.ReadBuffer, data...)
		for index, v := range c.ReadBuffer {
			if v == '\n' {
				if prev == '\r' {
					println("Index =>", index)
					println(string(c.ReadBuffer))
					c.ReadBuffer = nil
					break
				}
			}
			prev = v
		}
		//c.Write(append([]byte{}, data...))
	})

	err := g.Start()
	if err != nil {
		fmt.Printf("nbio.Start failed: %v\n", err)
		return
	}
	defer g.Stop()

	g.Wait()
}
