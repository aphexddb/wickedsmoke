package main

// via https://github.com/d2r2/go-i2c
//
// Package i2c provides low level control over the linux i2c bus.
//
// Before usage you should load the i2c-dev kernel module
//
//      sudo modprobe i2c-dev
//
// Each i2c bus can address 127 independent i2c devices, and most
// linux systems contain several buses.
//
// To allow access to the i2c devices as a standard user
//
//			sudo chmod o+rw /dev/i2c*
//
// Also this
// https://learn.adafruit.com/adafruit-16-channel-servo-driver-with-raspberry-pi/configuring-your-pi-for-i2c
//

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

const (
	I2C_SLAVE = 0x0703
)

// I2C represents a connection to an i2c device.
type I2C struct {
	rc *os.File
}

// NewI2C opens a connection to an i2c device.
func NewI2C(addr uint8, bus int) (*I2C, error) {
	f, err := os.OpenFile(fmt.Sprintf("/dev/i2c-%d", bus), os.O_RDWR, 0600)
	if err != nil {
		return nil, err
	}
	if err := ioctl(f.Fd(), I2C_SLAVE, uintptr(addr)); err != nil {
		return nil, err
	}
	this := &I2C{rc: f}
	return this, nil
}

// Write sends byte array to the remote i2c device. The interpretation of
// the message is implementation dependant.
func (d *I2C) Write(buf []byte) (int, error) {
	return d.rc.Write(buf)
}

// WriteByte byte to the remote i2c device. The interpretation of
// the message is implementation dependant.
func (d *I2C) WriteByte(b byte) (int, error) {
	var buf [1]byte
	buf[0] = b
	return d.rc.Write(buf[:])
}

// Read closes the remote i2c device
func (d *I2C) Read(p []byte) (int, error) {
	return d.rc.Read(p)
}

// Close closes the remote i2c device
func (d *I2C) Close() error {
	return d.rc.Close()
}

// ReadRegU8 reads byte from i2c device register specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (d *I2C) ReadRegU8(reg byte) (byte, error) {
	_, err := d.Write([]byte{reg})
	if err != nil {
		return 0, err
	}
	buf := make([]byte, 1)
	_, err = d.Read(buf)
	if err != nil {
		return 0, err
	}
	log.Printf("Read U8 %d from reg 0x%0X\n", buf[0], reg)
	return buf[0], nil
}

// WriteRegU8 writes byte to i2c device register specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (d *I2C) WriteRegU8(reg byte, value byte) error {
	buf := []byte{reg, value}
	_, err := d.Write(buf)
	if err != nil {
		return err
	}
	log.Printf("Write U8 %d to reg 0x%0X\n", value, reg)
	return nil
}

// ReadRegU16BE Read unsigned big endian word (16 bits) from i2c device
// starting from address specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (d *I2C) ReadRegU16BE(reg byte) (uint16, error) {
	_, err := d.Write([]byte{reg})
	if err != nil {
		return 0, err
	}
	buf := make([]byte, 2)
	_, err = d.Read(buf)
	if err != nil {
		return 0, err
	}
	w := uint16(buf[0])<<8 + uint16(buf[1])
	log.Printf("Read U16 %d from reg 0x%0X\n", w, reg)
	return w, nil
}

// ReadRegU16LE reads unsigned little endian word (16 bits) from i2c device
// starting from address specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (d *I2C) ReadRegU16LE(reg byte) (uint16, error) {
	w, err := d.ReadRegU16BE(reg)
	if err != nil {
		return 0, err
	}
	// exchange bytes
	w = (w&0xFF)<<8 + w>>8
	return w, nil
}

// ReadRegS16BE reads signed big endian word (16 bits) from i2c device
// starting from address specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (d *I2C) ReadRegS16BE(reg byte) (int16, error) {
	_, err := d.Write([]byte{reg})
	if err != nil {
		return 0, err
	}
	buf := make([]byte, 2)
	_, err = d.Read(buf)
	if err != nil {
		return 0, err
	}
	w := int16(buf[0])<<8 + int16(buf[1])
	log.Printf("Read S16 %d from reg 0x%0X\n", w, reg)
	return w, nil
}

// ReadRegS16LE reads unsigned little endian word (16 bits) from i2c device
// starting from address specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (d *I2C) ReadRegS16LE(reg byte) (int16, error) {
	w, err := d.ReadRegS16BE(reg)
	if err != nil {
		return 0, err
	}
	// exchange bytes
	w = (w&0xFF)<<8 + w>>8
	return w, nil

}

// WriteRegU16BE writes unsigned big endian word (16 bits) value to i2c device
// starting from address specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (d *I2C) WriteRegU16BE(reg byte, value uint16) error {
	buf := []byte{reg, byte((value & 0xFF00) >> 8), byte(value & 0xFF)}
	_, err := d.Write(buf)
	if err != nil {
		return err
	}
	log.Printf("Write U16 %d to reg 0x%0X\n", value, reg)
	return nil
}

// WriteRegU16LE writes unsigned big endian word (16 bits) value to i2c device
// starting from address specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (d *I2C) WriteRegU16LE(reg byte, value uint16) error {
	w := (value*0xFF00)>>8 + value<<8
	return d.WriteRegU16BE(reg, w)
}

// WriteRegS16BE writes signed big endian word (16 bits) value to i2c device
// starting from address specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (d *I2C) WriteRegS16BE(reg byte, value int16) error {
	buf := []byte{reg, byte((uint16(value) & 0xFF00) >> 8), byte(value & 0xFF)}
	_, err := d.Write(buf)
	if err != nil {
		return err
	}
	log.Printf("Write S16 %d to reg 0x%0X\n", value, reg)
	return nil
}

// WriteRegS16LE writes signed big endian word (16 bits) value to i2c device
// starting from address specified in reg.
// SMBus (System Management Bus) protocol over I2C.
func (d *I2C) WriteRegS16LE(reg byte, value int16) error {
	w := int16((uint16(value)*0xFF00)>>8) + value<<8
	return d.WriteRegS16BE(reg, w)
}

func ioctl(fd, cmd, arg uintptr) error {
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, fd, cmd, arg, 0, 0, 0)
	if err != 0 {
		return err
	}
	return nil
}
