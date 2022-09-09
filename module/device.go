package module

import "strings"

var FakeDataDevices = []Device{
	0: {
		ID:           1,
		Name:         "device_01",
		IP:           "192.168.0.1",
		MacAddress:   "xx:xx:xx:xx:xx:xx",
		DeviceSerial: "J91322386",
		Status:       true,
	},
}

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
	FakeDataDevices = append(FakeDataDevices, *new)
	return true
}

func (d *Device) GetByDeviceSerial(serial_id string) Device {
	output := Device{}
	for _, fakedevice := range FakeDataDevices {
		if strings.Compare(fakedevice.DeviceSerial, serial_id) == 0 {
			output = fakedevice
			break
		}
	}
	return output
}

func (d *Device) List(num int, page int) []Device {
	return FakeDataDevices
}

func (d *Device) Del(id int) error {
	NewFakeDataDevices := []Device{}
	for _, v := range FakeDataDevices {
		if v.ID != id {
			NewFakeDataDevices = append(NewFakeDataDevices, v)
		}
	}
	FakeDataDevices = NewFakeDataDevices
	return nil
}

func (d *Device) Count() int {
	return len(FakeDataDevices)
}
