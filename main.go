package main

import (
	"log"
	"os"

	"go.bug.st/serial"
)

func main() {

	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	for _, port := range ports {
		log.Printf("Found port: %v\n", port)
	}

	xm := xModemInit("COM11", 256, serial.Mode{
		BaudRate: 115200,
		Parity:   serial.NoParity,
	})

	/* Open file & prepare to transfer */
	fd, err := os.Open("test.img") // For read access.
	if err != nil {
		log.Fatal("Error in file open")
	}
	defer fd.Close()

	var fs os.FileInfo

	fs, err = fd.Stat()
	if err != nil {
		log.Fatal("Error in retrieve file info")
	}
	xm.FileTransfer(fd, fs.Size())

}
