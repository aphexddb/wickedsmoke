// dependencies managed via https://github.com/kardianos/govendor
package main

import (
	"fmt"
	"os"

	"github.com/mrmorphic/hwio"
)

func main() {
	fmt.Println("Wickedsmoke alpha")

	// set Raspberry Pi driver
	hwio.SetDriver(new(hwio.RaspberryPiDTDriver))

	myPin, pinErr := hwio.GetPinWithMode("gpio4", hwio.OUTPUT)
	if pinErr != nil {
		fmt.Printf("hwio: error getting GPIO gpio4 %s", pinErr.Error())
		os.Exit(1)
	}

	value, readErr := hwio.DigitalRead(myPin)
	if readErr != nil {
		fmt.Printf("hwio: error reading GPIO gpio4 %s", readErr.Error())
		os.Exit(1)
	}
	fmt.Println(value)

}
