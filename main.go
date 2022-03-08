package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/lesismal/nbio"
	"github.com/snksoft/crc"
)

type GpsTrackerConnection struct {
	conn *nbio.Conn
}

var maxLengthPacket int = 1024

func hasGps(protocol byte) bool {
	switch protocol {
	case 0x12, 0x16, 0x22, 0x27:
		return true
	}
	return false
}
func hasHearbeat(protocol byte) bool {
	switch protocol {
	case 0x13, 0x23:
		return true
	}
	return false
}

func decode_gt06(c *nbio.Conn) {
	/*
		header 		2byte 0x78 0x78
		length		1byte (from protocol to Error Check)
		protocol	1byte
		Content		NByte
		Serial		2byte
		Error Check	2Byte
		End			2byte 0x0D 0x0A

	*/
	if c.ReadBuffer[0] == 0x78 && c.ReadBuffer[1] == 0x78 {
		var lengthPacket = c.ReadBuffer[2]
		//var protocol = c.ReadBuffer[3]
		if int(lengthPacket+5) < len(c.ReadBuffer) {
			return
		}
		var protocol = c.ReadBuffer[3]
		if protocol == 0x01 {
			/*
				Protocol Login
				header 		2byte 0x78 0x78
				length		1byte (from protocol to Error Check)
				protocol	1byte
				IMEI		8Byte
				Serial		2byte
				Error Check	2Byte
				End			2byte 0x0D 0x0A

			*/
			var imei = hex.EncodeToString(c.ReadBuffer[4:12])
			if imei == "437444" {
				println("GPS Found")
			}

			//answer data
			response := []byte{}
			response = append(response, 0x78, 0x78, 0x1)
			response = append(response, c.ReadBuffer[12], c.ReadBuffer[13]) //add serial

			ccittCrc := uint16(crc.CalculateCRC(crc.X25, response))
			b := make([]byte, 2)
			binary.BigEndian.PutUint16(b, ccittCrc)
			response = append(response, b...)
			response = append(response, 0xd, 0xA)
			c.Write(response)

			//parseLogin(c, c.ReadBuffer[4:12])
		} else if hasGps(protocol) {
			//
		}
	}
}
func parseLogin(c *nbio.Conn, data []byte) {
	//var imei = hex.EncodeToString(data)

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
		var index = 0
		for _, v := range c.ReadBuffer {
			if v == '\n' {
				if prev == '\r' {
					println(string(c.ReadBuffer))
					c.ReadBuffer = nil
					break
				}
			}
			prev = v
			index++
		}

		decode_gt06(c)
	})

	err := g.Start()
	if err != nil {
		fmt.Printf("nbio.Start failed: %v\n", err)
		return
	}
	defer g.Stop()

	g.Wait()
}
