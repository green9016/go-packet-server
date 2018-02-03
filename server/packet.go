package server

import (
	"fmt"
	"encoding/binary"
	"bytes"
)

type Packet struct {
	data []byte
}

func ToPacket(data []byte) *Packet {
	return &Packet{data}
}

func NewPacket(sn string, _type string) *Packet {
	var packet *Packet
	packet = &Packet{[]byte{0x28}}
	packet.Write(sn)
	packet.Write(_type)
	return packet
}

func (p *Packet) EndPacket() {
	p.Write(0x29)
}

func (p *Packet) Size() uint16 {
	return uint16(len(p.data))
}

func (p *Packet) Type() string {
	return string(p.data[13:17])
}

func (p *Packet) SN() string {
	return string(p.data[1:13])
}

func (p *Packet) containsString() bool {
	if len(p.data) <= 18 {
		return false
	}

	return true
}

func (p *Packet) ExtractString() (uint16, string) {
	if !p.containsString() {
		return 0, ""
	}

	size := uint16(len(p.data))

	return size-18, string(p.data[17:size-1])
}

func (p *Packet) Buffer() []byte {
	return p.data
}

func (p *Packet) Stringify() string {
	return string(p.data[:])
}

func (p *Packet) String() string {
	return fmt.Sprintf("%d ", p.data)
}

func (p *Packet) Write(data ...interface{}) {
	for _, value := range data {
		switch v := value.(type) {
		case uint8:
			p.data = append(p.data, v)
		case int8:
			p.data = append(p.data, byte(v))
		case uint16:
			b := make([]byte, 2)
			binary.LittleEndian.PutUint16(b, uint16(v))
			p.data = append(p.data, b...)
		case int16:
			b := make([]byte, 2)
			binary.LittleEndian.PutUint16(b, uint16(v))
			p.data = append(p.data, b...)
		case uint32:
			b := make([]byte, 4)
			binary.LittleEndian.PutUint32(b, uint32(v))
			p.data = append(p.data, b...)
		case int32:
			b := make([]byte, 4)
			binary.LittleEndian.PutUint32(b, uint32(v))
			p.data = append(p.data, b...)
		case int:
			b := make([]byte, 4)
			binary.LittleEndian.PutUint32(b, uint32(v))
			p.data = append(p.data, b...)
		case uint64:
			b := make([]byte, 8)
			binary.LittleEndian.PutUint64(b, uint64(v))
			p.data = append(p.data, b...)
		case int64:
			b := make([]byte, 8)
			binary.LittleEndian.PutUint64(b, uint64(v))
			p.data = append(p.data, b...)
		case string:
			b := []byte(v)
			p.data = append(p.data, b...)
		case float32:
			var b bytes.Buffer
			err := binary.Write(&b, binary.LittleEndian, float32(v))
			if err == nil {
				p.data = append(p.data, b.Bytes()...)
			}
		case float64:
			var b bytes.Buffer
			err := binary.Write(&b, binary.LittleEndian, float64(v))
			if err == nil {
				p.data = append(p.data, b.Bytes()...)
			}
		}
	}
}