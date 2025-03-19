package main

import (
	"log"

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

	/* Settings of selected COM port */
	portSettings serial.Mode
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
func xModemInit(payload_size uint16, pors_settings serial.Mode) *xModem {

	xm := xModem{
		payloadSize:  payload_size,
		portSettings: pors_settings,
	}

	return &xm
}

/**
 * Perform file transfer
 */
func (xm *xModem) FileTransfer(file []byte, size uint32) {

	log.Println("Perform file transfer")

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
