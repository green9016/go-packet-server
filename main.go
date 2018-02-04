package main

import (
	"fmt"
  "./server"
  "./api"
)

/*
	This example takes you to character selection screen when using game client of game KalOnline.
 */

func main() {
  go api.StartApi();
	game := server.New("0.0.0.0:3000")

  // Subscribe to connection event.
  game.OnConnected(func(s *server.Session) {
    fmt.Println("Client connected.")
  })

  // Subscribe to disconnection event.
  game.OnDisconnected(func(s *server.Session) {
    fmt.Println("Client disconnected.")
  })

  // This event fires when undefined packet was received.
  // It means client send packet of type you did not subscribed to.
  game.OnUnknownPacket(func(s *server.Session, p *server.Packet) {
    fmt.Println("Unknown packet:", p.Type())
  })

  game.On("BP00", func(s *server.Session, p *server.Packet) {
    sn:=p.SN()
    fmt.Println("Sync:", sn)

    s.Send(sn, "AP01", "HSO")
  })

  //login
  game.On("BP05", func(s *server.Session, p *server.Packet) {
    sn:=p.SN()
    _, data:=p.ExtractString()
    fmt.Println("Login:", sn)

    go server.ProcessDeviceLogin(sn, data)
    s.Send(sn, "AP05", "")
  })

  game.On("BO01", func(s *server.Session, p *server.Packet) {
    sn:=p.SN()
    _, data:=p.ExtractString()
    fmt.Println("Alarm:", sn, data)

    server.ProcessAlarm(sn, data)
    s.Send(sn, "AS01", data[0:1])
  })

  game.On("BR03", func(s *server.Session, p *server.Packet) {
    sn:=p.SN()
    _, data:=p.ExtractString()
    fmt.Println("Location:", sn, data)

    go server.ProcessLocation(sn, data)
  })

  game.On("BR01", func(s *server.Session, p *server.Packet) {
    sn:=p.SN()
    _, data:=p.ExtractString()
    fmt.Println("Location:", sn, data)

    go server.ProcessLocation(sn, data)
  })

  game.On("BR02", func(s *server.Session, p *server.Packet) {
    sn:=p.SN()
    _, data:=p.ExtractString()
    fmt.Println("Location:", sn, data)

    go server.ProcessLocation(sn, data)
  })

  game.On("BR03", func(s *server.Session, p *server.Packet) {
    sn:=p.SN()
    _, data:=p.ExtractString()
    fmt.Println("Location:", sn, data)

    go server.ProcessLocation(sn, data)
  })

  game.Start()
}