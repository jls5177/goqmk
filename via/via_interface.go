package via

import (
	"fmt"
	"github.com/bearsh/hid"
	"log"
	"time"
)

type Device struct {
	handle *hid.Device
}

const (
	ViaDeviceUsage = 0x61
	ViaDeviceUsagePage = 0xff60
)

func New(vendorId, productId uint16) (*Device, error) {
	for _, device := range hid.Enumerate(vendorId, productId) {
		if device.Usage == ViaDeviceUsage && device.UsagePage == ViaDeviceUsagePage {
			handle, err := device.Open()
			if err != nil {
				return nil, fmt.Errorf("failed opening handle: %v", err)
			}
			return &Device{handle: handle}, nil
		}
	}
	return nil, fmt.Errorf("no Via devices found")
}

func (d *Device) Close() {
	err := d.handle.Close()
	if err != nil {
		log.Fatalf("failed opening handle: %v", err)
	}
}

func (d *Device) GotoBootloader() error {
	return d.sendCommand(IdBootloaderJump, []uint8{})
}

func (d *Device) Debug(value uint8) error {
	return d.sendCommand(IdDebug, []uint8{value})
}

func (d *Device) Lighting() *DeviceLighting {
	// TODO: add support to determine if using RGB or Backlight
	return &DeviceLighting{device: d}
}

type DeviceLighting struct {
	device *Device
}

func (l *DeviceLighting) setValue(lightingCmdId ViaLightingSubCmdId, data []uint8) error {
	defer time.Sleep(100*time.Millisecond)
	return l.device.sendCommand(IdLightingSetValue, append([]uint8{lightingCmdId}, data...))
}

func (l *DeviceLighting) Save() error {
	return l.device.sendCommand(IdLightingSave, []uint8{})
}

func (l *DeviceLighting) Brightness(value uint8) error {
	return l.setValue(IdQmkRgblightBrightness, []uint8{value})
}

func (l *DeviceLighting) Effect(value uint8) error {
	return l.setValue(IdQmkRgblightEffect, []uint8{value})
}

func (l *DeviceLighting) EffectSpeed(value uint8) error {
	return l.setValue(IdQmkRgblightEffectSpeed, []uint8{value})
}

func (l *DeviceLighting) Color(hue, saturation uint8) error {
	return l.setValue(IdQmkRgblightColor, []uint8{hue, saturation})
}

func (d *Device) sendCommand(commandId ViaCommandId, data []uint8) error {
	buf := []uint8{commandId}
	buf = append(buf, data...)
	size, err := d.handle.Write(buf)  // reboot to bootloader
	if err != nil {
		return fmt.Errorf("failed opening handle: %v", err)
	} else if size != len(buf) {
		return fmt.Errorf("failed to write all bytes: %v, expected %v", size, len(buf))
	}
	return nil
}