package main

import (
	"bytes"
	"log"

	"github.com/chia0hwan9/goeip"
)

// Demo program for reading and writing BOOL tags on Omron NJ/NX and Inovance Easy521
// series PLCs that use 2-byte BOOL alignment on the wire (only the LSB is valid).
//
// Rockwell/Allen-Bradley Logix PLCs use 1-byte BOOLs by default; set BoolSize = 2
// for controllers that pack 16 BOOLs per WORD with 2-byte per BOOL on the wire.
func main() {
	var err error

	// setup the client for an Omron/Inovance controller
	client := goeip.NewClient("192.168.2.100")
	// Micro8xx / Omron NJ use no backplane path. Adjust as needed for your device.
	client.Controller.Path = &bytes.Buffer{}
	// KEY: set BoolSize to 2 for Omron/Inovance 2-byte BOOL alignment
	client.BoolSize = 2

	// connect using parameters in the client struct
	err = client.Connect()
	if err != nil {
		log.Printf("Error opening client: %v", err)
		return
	}
	defer client.Disconnect()

	// --- Read a single BOOL ---
	var myBool bool
	err = client.Read("MyBoolTag", &myBool)
	if err != nil {
		log.Printf("error reading 'MyBoolTag': %v", err)
	} else {
		log.Printf("MyBoolTag = %v", myBool)
	}

	// --- Write a single BOOL ---
	err = client.Write("MyBoolTag", true)
	if err != nil {
		log.Printf("error writing 'MyBoolTag': %v", err)
	}

	// --- Read an array of BOOLs (must be multiple of 16 with BoolSize=2) ---
	boolArray := make([]bool, 16)
	err = client.Read("MyBoolArray", boolArray)
	if err != nil {
		log.Printf("error reading 'MyBoolArray': %v", err)
	} else {
		log.Printf("MyBoolArray = %v", boolArray)
	}

	// --- Write an array of BOOLs ---
	err = client.Write("MyBoolArray", []bool{true, false, true, false})
	if err != nil {
		log.Printf("error writing 'MyBoolArray': %v", err)
	}

	// --- Read a DINT (other types work normally regardless of BoolSize) ---
	var myDint int32
	err = client.Read("MyDintTag", &myDint)
	if err != nil {
		log.Printf("error reading 'MyDintTag': %v", err)
	} else {
		log.Printf("MyDintTag = %d", myDint)
	}
}
