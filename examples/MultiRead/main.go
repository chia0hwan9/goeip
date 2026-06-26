package main

import (
	"log"

	"github.com/chia0hwan9/goeip"
)

// Demo program for readng an INT tag named "TestInt" in the controller.
func main() {
	var err error

	// setup the client.  If you need a different path you'll have to set that.
	client := goeip.NewClient("192.168.2.241")

	// for example, to have a controller on slot 1 instead of 0 you could do this
	//client.Path, err = goeip.Serialize(goeip.CIPPort{PortNo: 1}, goeip.CIPAddress(1))
	// or this
	// client.Path, err = goeip.ParsePath("1,1")

	// connect using parameters in the client struct
	err = client.Connect()
	if err != nil {
		log.Printf("Error opening client. %v", err)
		return
	}
	// setup a deffered disconnect.  If you don't disconnect you might have trouble reconnecting because
	// you won't have sent the close forward open.  You'll have to wait for the CIP connection to time out
	// if that happens (about a minute)
	defer client.Disconnect()

	// define a struct where fields have the tag to read from the controller specified
	// note that tag names are case insensitive.
	type multiread struct {
		TestInt  int16   `goeip:"TestInt"`
		TestDint int32   `goeip:"TestDint"`
		TestArr  []int32 `goeip:"TestDintArr[2]"`
	}
	var mr multiread
	mr.TestArr = make([]int32, 5)

	// call the read multi function with the structure passed in as a pointer.
	err = client.ReadMulti(&mr)
	if err != nil {
		log.Printf("error reading testint. %v", err)
	}
	// do whatever you want with the values
	log.Printf("multiread struct has values %+v", mr)

}
