// go run main.go  -logtostderr=true -v=3
package main

import (
	// "encoding/json"
	"flag"
	"fmt"
	"time"

	// "github.com/golang/glog"
	"github.com/kidoman/embd"
	"github.com/kidoman/embd/convertors/mcp3008"
	_ "github.com/kidoman/embd/host/all"
)

const (
	channel = 0
	speed   = 1000000
	bpw     = 8
	delay   = 0
)

const Vref = float64(3.3)
const Vresolution = float64(1023)

// ADC is a MCP3008 ADC device
type ADC struct {
	values []float64
	adc *MCP3008
}

// NewADC creates a new ADC using SPI to read data
func NewADC() *ADC {
	if err := embd.InitSPI(); err != nil {
		panic(err)
	}
	defer embd.CloseSPI()

	spiBus := embd.NewSPIBus(embd.SPIMode0, channel, speed, bpw, delay)
	defer spiBus.Close()

	adc := mcp3008.New(mcp3008.SingleMode, spiBus)

	return &ADC{
		values: [8]float64{0,0,0,0,0,0,0,0},
		adc: adc
	}
}

// Read gets sensor voltages for each channel
func (adc *ADC) Read() {
	for i := 0; i < 8; i++ {
		val, err := adc.adc.AnalogValueAt(i)
		if err != nil {
			panic(err)
		}
		sensorVoltage := convertADCVoltage(val, 2)
		adc.values[i] = sensorVoltage
		// glog.Infof("analog value %v is: %v sensorVoltage -> %v\n", i, val, sensorVoltage)
	}
}

// The MCP3008 ADC is powered via 3V3 from the Raspberry Pi.
// This is also its default reference voltage Vref.
// So if the measured analog input signal is 3.3V,  the ADC will output 1023.
// If the input signal is 0V, the ADC output will be 0.
// If we divide 3.3V by 1023, we get the resolution of the device.
// 3.3V / 1023 = 0.00322 V/step. That's 3 mV.
//
// If we want to read an analog sensor's voltage we do the following calculation...
// ADC reading / 1023 * 3.3 V = Sensor Voltage
func convertADCVoltage(adcReading int, decimalPlaces int) float64 {
	sensorVoltage := float64(adcReading) / Vresolution * Vref
	return sensorVoltage
}


// func main() {
// 	flag.Parse()
// 	glog.Infoln("this is a sample code for mcp3008 10bit 8 channel ADC")

// 	if err := embd.InitSPI(); err != nil {
// 		panic(err)
// 	}
// 	defer embd.CloseSPI()

// 	spiBus := embd.NewSPIBus(embd.SPIMode0, channel, speed, bpw, delay)
// 	defer spiBus.Close()

// 	adc := mcp3008.New(mcp3008.SingleMode, spiBus)

// 	values := [8]float64{0, 0, 0, 0, 0, 0, 0, 0}

// 	for {
// 		for i := 0; i < 8; i++ {
// 			time.Sleep(100 * time.Millisecond)
// 			val, err := adc.AnalogValueAt(0)
// 			if err != nil {
// 				panic(err)
// 			}
// 			sensorVoltage := convertADCVoltage(val, 2)
// 			values[i] = sensorVoltage

// 			// glog.Infof("analog value %v is: %v sensorVoltage -> %v\n", i, val, sensorVoltage)
// 		}
// 		v, _ := json.Marshal(values)
// 		fmt.Println(string(v))
// 	}
// }