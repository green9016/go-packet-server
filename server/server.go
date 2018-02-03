package server

import (
	"fmt"
	"net"
	"io"
	"bufio"
//	"encoding/binary"
	"github.com/satori/go.uuid"
)

type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}

type Server struct {
	address	string
	events map[string]func(s *Session, p* Packet)
	sessions map[string]*Session
	onConnected func(s *Session)
	onDisconnected func(s *Session)
	onUnknownPacket func(s *Session, p* Packet)
}

type Session struct {
	conn net.Conn
	reader *bufio.Reader
	id uuid.UUID
}

func New(address string) *Server {
	return &Server{
		address,
		make(map[string]func(s *Session, p *Packet)),
		make(map[string]*Session),
		func(s *Session) {},
		func(s *Session) {},
		func(s *Session, p *Packet) {}}
}

func (s *Server) OnConnected(callback func(s *Session)) {
	s.onConnected = callback
}

func (s *Server) OnDisconnected(callback func(s *Session)) {
	s.onDisconnected = callback
}

func (s *Server) OnUnknownPacket(callback func(s *Session, p *Packet)) {
	s.onUnknownPacket = callback
}

func (s *Server) On(type_ string, callback func(s *Session, p *Packet)) {
	s.events[type_] = callback
}

func (s *Server) ForEachSession(callback func(session *Session)) {
	for _, sess := range s.sessions {
		callback(sess)
	}
}

func (s *Server) ForEach(ids []string, callback func(session *Session)) {
	for _, id := range ids {
		if sess, ok := s.sessions[id]; ok {
			callback(sess)
		}
	}
}

func (s *Server) For(id string, callback func(session *Session)) {
	if sess, ok := s.sessions[id]; ok {
		callback(sess)
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go s.listen(&Session{
			conn,
			bufio.NewReader(conn),
			uuid.Must(uuid.NewV4())})
	}
}

func (s *Server) listen(session *Session) {
	s.sessions[session.Id()] = session
	s.onConnected(session)

	defer func() {
		s.onDisconnected(session)
		delete(s.sessions, session.Id())
		session.conn.Close()
	}()

	for {
		packet, err := session.receive()
		if err != nil {
			return
		}

		fmt.Printf("%v\n", packet.data)
		fmt.Printf("%s\n", string(packet.data))
		fmt.Printf("%s\n", packet.Type())

		if event, ok := s.events[packet.Type()]; ok {
			event(session, packet)
		} else {
			s.onUnknownPacket(session, packet)
		}
	}
}

func (s *Session) receive() (*Packet, error) {
	buffer := make([]byte, 1024)
	n, err := s.reader.Read(buffer)
	fmt.Printf("Read data=%d\n", n)
	if err == io.EOF || err != nil {
		return nil, err
	}

	if(buffer[0] != 0x28 || n > 1024 ||  buffer[n - 1] != 0x29) {
		return nil, &errorString{"Wrong packet"};
	}

	packet := ToPacket(buffer[:n])

	return packet, nil
}

func (s *Session) Send(sn string, _type string, data ...interface{}) int {
	p := NewPacket(sn, _type)
	p.Write(data...)
	p.EndPacket()
	return s.SendPacket(p)
}

func (s *Session) SendPacket(p *Packet) int {
	n, _ := s.conn.Write(p.Buffer())
	return n
}

func (s *Server) Broadcast(sn string,  _type string, data ...interface{}) {
	p := NewPacket(sn, _type)
	p.Write(data...)
	p.EndPacket()
	s.BroadcastPacket(p)
}

func (s *Server) BroadcastPacket(p *Packet) {
	for _, session := range s.sessions {
		session.SendPacket(p)
	}
}

func (s *Session) Id() string {
	return s.id.String()
}

