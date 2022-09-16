package module

import "strings"

type IDevice interface {
	Create(new *Device) bool
	GetByDeviceSerial(serial_id string) Device
	List(num int, page int) []Device
	Del(id int) error
	Count() int
}

type Device struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	IP           string `json:"ip"`
	MacAddress   string `json:"mac_address"`
	DeviceSerial string `json:"device_serial"`
	Status       bool   `json:"status"`
}

func NewDevice() IDevice {
	return &Device{}
}

func (d *Device) Create(new *Device) bool {
	FakeDevices = append(FakeDevices, *new)
	return true
}

func (d *Device) GetByDeviceSerial(serial_id string) Device {
	output := Device{}
	for _, fakedevice := range FakeDevices {
		if strings.Compare(fakedevice.DeviceSerial, serial_id) == 0 {
			output = fakedevice
			break
		}
	}
	return output
}

func (d *Device) List(num int, page int) []Device {
	return FakeDevices
}

func (d *Device) Del(id int) error {
	NewFakeDevices := []Device{}
	for _, v := range FakeDevices {
		if v.ID != id {
			NewFakeDevices = append(NewFakeDevices, v)
		}
	}
	FakeDevices = NewFakeDevices
	return nil
}

func (d *Device) Count() int {
	return len(FakeDevices)
}
