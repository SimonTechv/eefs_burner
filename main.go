package main

import (
	"fmt"
	"log"

	"go.bug.st/serial"
)

func main() {

	ports, err := GetDevices()

	if err != nil {
		log.Fatal("No found devices")
	} else {
		fmt.Printf("Found %d serial devices:\n", len(ports))

		for n, port := range ports {
			fmt.Printf("\t%d: %s\n", n, port)
		}
	}

	xm := xModemInit(256, serial.Mode{
		BaudRate: 9600,
		Parity:   serial.NoParity,
	})

	log.Println(xm)
}
