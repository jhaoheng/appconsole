package module

import "strings"

var FakeDataDevices = []Device{}

type IDevice interface{}

type Device struct {
	Name         string `json:"name"`
	IP           string `json:"ip"`
	MacAddress   string `json:"mac_address"`
	DeviceSerial string `json:"device_serial"`
	Status       bool   `json:"status"`
}

func NewDevice() IDevice {
	return &Device{}
}

func (d *Device) Create(newdevice *Device) bool {
	FakeDataDevices = append(FakeDataDevices, *newdevice)
	return true
}

func (d *Device) GetByDeviceSerial(id string) Device {
	output := Device{}
	for _, fakedevice := range FakeDataDevices {
		if strings.Compare(fakedevice.DeviceSerial, id) == 0 {
			output = fakedevice
			break
		}
	}
	return output
}

func (d *Device) List(num int, page int) []Device {
	return FakeDataDevices
}
