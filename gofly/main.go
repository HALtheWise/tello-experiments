package main

import (
	"log"
	"time"

	"github.com/SMerrony/tello"
	"github.com/davecgh/go-spew/spew"
)

var _ tello.Tello

func main() {
	drone := new(tello.Tello)
	err := drone.ControlConnectDefault()
	defer drone.ControlDisconnect()
	if err != nil {
		log.Fatalf("Unable to connect to Tello: %v", err)
	}

	drone.TakeOff()
	time.Sleep(2.0 * time.Second)

	drone.GetSSID()
	drone.GetVersion()
	drone.GetMaxHeight()
	time.Sleep(100 * time.Millisecond)

	data := drone.GetFlightData()
	spew.Dump(data)

	time.Sleep(3.0 * time.Second)
	log.Println("Flipping")
	drone.ForwardLeftFlip()
	time.Sleep(2.0 * time.Second)
	log.Println("Hand landing")
	drone.PalmLand()
	time.Sleep(5 * time.Second)
}
