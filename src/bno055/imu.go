// Package bno055 provides tools for interfacing with Bosch BNO055 IMU chip.
package bno055

import (
	"fmt"
	"log"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/i2c"
)

const CHIP_ADDRESS = 0x29

// IMU servers as a  base datatype for this chip.
type IMU struct {
	i2cDev i2c.Dev
	pinSDA gpio.PinIO
	pinSCL gpio.PinIO

	// data
	// Lat float64
}

//Init initializes the lora communication channel
//chipAddress is the i2c address of the BNO055 chip.
//GPIOdev can be left "" to pick the first usable i2c device
func (n *IMU) Init(bus i2c.BusCloser) {
	// creates new i2cdevice on default address
	n.i2cDev = i2c.Dev{Addr: 0x29, Bus: bus}
	if pins, err := bus.(i2c.Pins); err {
		n.pinSDA = pins.SDA()
		n.pinSDA = pins.SCL()
	} else {
		log.Fatal(err)
	}

	// Send a command 0x10 and expect a 1 byte reply.
	write := []byte{0x36} // calib state
	read := make([]byte, 1)
	if err := n.i2cDev.Tx(write, read); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x\n", read)
}

//Calibrate calibrates the chip...
func (n *IMU) Calibrate() {

}

//GetCalibratedVal gets the calibration statuses as a byte
func (n *IMU) GetCalibratedVal() (read []byte, err error) {
	// Send a command 0x10 and expect a 1 byte reply.
	read, err = n.ReadBytes(0x35, 1)
	return
}

func (n *IMU) ReadBytes(addr byte, length int) (read []byte, err error) {
	write := []byte{addr} // calib state
	read = make([]byte, length)
	if err := n.i2cDev.Tx(write, read); err != nil {
		return nil, err
	}
	return
}

// func (n *imuBNO055) Calibrate(chipAddress int){

// }
