package main

import (
	"fmt"
	"log"
	"os"

	"go.bug.st/serial"
)

/* Special symbols */
const (
	SOH   = 0x01
	EOT   = 0x04
	ACK   = 0x06
	NAK   = 0x15
	ETB   = 0x17
	CAN   = 0x18
	C     = 0x43
	CTRLZ = 0x1A
)

type xModem struct {
	payloadSize uint16

	portHandle   serial.Port
	portSettings serial.Mode
	portName     string
}

/**
 *
 *
 *
 *
 *
 *
 *
 * Public methods
 *
 *
 *
 *
 *
 *
 */

/**
 * Initialize transfer object
 */
func xModemInit(port_name string, payload_size uint16, port_settings serial.Mode) *xModem {

	xm := xModem{
		payloadSize:  payload_size,
		portSettings: port_settings,
		portName:     port_name,
	}

	fmt.Println(xm)

	/* Prepare port to transfer */
	err := error(nil)
	xm.portHandle, err = serial.Open(xm.portName, &xm.portSettings)
	if err != nil {
		log.Fatal("Error in opening port!", err)
	}

	return &xm
}

/**
 * Perform file transfer
 */
func (xm *xModem) FileTransfer(file *os.File, size int64) {

	log.Println("Perform file transfer")

	slice := []byte{1, 2, 3, 4, 5, 6, 7}

	xm.portHandle.Write(slice)

	/* Flush buffers */

	/* Wait while slave is ready */
}

/**
 *
 *
 *
 *
 *
 *
 *
 * Private methods
 *
 *
 *
 *
 *
 *
 */

/**
 * Perform block transfer
 */
func (xm *xModem) blockTransfer(file []byte, size uint32) {
	log.Println("Perform block transfer")
}

/**
 *
 *
 *
 *
 *
 *
 *
 * Wrappers
 *
 *
 *
 *
 *
 *
 */

/**
 * Scan available serial devices
 */
func GetDevices() ([]string, error) {

	/* Retrieve available serial devices */
	return serial.GetPortsList()
}
