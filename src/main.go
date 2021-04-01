package main

import (
	"fmt"
	"log"
	"time"

	"bno055"

	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/host"
)

func main() {

	var chip bno055.IMU

	// Initialize normally. Your driver will be loaded:
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// shared acces for main
	bus, err := i2creg.Open("1")
	if err != nil {
		log.Fatal(err)
	}
	defer bus.Close()
	fmt.Printf("Current active bus: %s\n", bus.String())

	chip.Init(bus)

	for {
		read, _ := chip.GetCalibratedVal()
		fmt.Printf("%x\n", read)

		stat, _ := chip.ReadBytes(0x36, 1)
		fmt.Printf("stat: %x\n\n", stat)

		time.Sleep(200 * time.Millisecond)

	}
}
