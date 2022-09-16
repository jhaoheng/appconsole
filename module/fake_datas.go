package module

import (
	"embed"
	"time"

	"github.com/google/uuid"
)

//go:embed fake_pic
var fakepic embed.FS

var pic = func(file string) []byte {
	data, err := fakepic.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return data
}

var FakeUsers = []User{
	0: {ID: 1, Name: "中文支援", MemberID: uuid.New().String(), Picture: pic("fake_pic/fake_pic_1.jpg"), Phone: "(02)1111111", Gender: "woman"},
	1: {ID: 2, Name: "中文2", MemberID: uuid.New().String(), Picture: pic("fake_pic/fake_pic_2.jpg"), Phone: "(02)2222222", Gender: "woman"},
	2: {ID: 3, Name: "strange_3", MemberID: uuid.New().String(), Picture: pic("fake_pic/fake_pic_3.jpg"), Phone: "(02)3333333", Gender: "woman"},
	3: {ID: 4, Name: "strange_4", MemberID: uuid.New().String(), Picture: pic("fake_pic/fake_pic_4.jpg"), Phone: "(02)4444444", Gender: "woman"},
	4: {ID: 5, Name: "strange_5", MemberID: uuid.New().String(), Picture: pic("fake_pic/fake_pic_5.jpg"), Phone: "(02)5555555", Gender: "woman"},
}

var FakeDevices = []Device{
	0: {
		ID:           1,
		Name:         "device_01",
		IP:           "192.168.0.1",
		MacAddress:   "xx:xx:xx:xx:xx:xx",
		DeviceSerial: "J91322386",
		Status:       true,
	},
}

var FakeUserLogs = []UserLog{
	0:  {ID: 1, Name: "One", RecordTime: time.Now(), Label: "do something", Created: time.Now(), Updated: time.Now()},
	1:  {ID: 2, Name: "Two", RecordTime: time.Now(), Label: "do something", Created: time.Now(), Updated: time.Now()},
	2:  {ID: 3, Name: "Three", RecordTime: time.Now(), Label: "do something", Created: time.Now(), Updated: time.Now()},
	3:  {ID: 4, Name: "Four", RecordTime: time.Now(), Label: "do something", Created: time.Now(), Updated: time.Now()},
	4:  {ID: 5, Name: "Five", RecordTime: time.Now(), Label: "do something", Created: time.Now(), Updated: time.Now()},
	5:  {ID: 6, Name: "Six", RecordTime: time.Now(), Label: "do something", Created: time.Now(), Updated: time.Now()},
	6:  {ID: 7, Name: "Seven", RecordTime: time.Now(), Label: "do something", Created: time.Now(), Updated: time.Now()},
	7:  {ID: 8, Name: "Eight", RecordTime: time.Now(), Label: "do something", Created: time.Now(), Updated: time.Now()},
	8:  {ID: 9, Name: "Nine", RecordTime: time.Now(), Label: "do something", Created: time.Now(), Updated: time.Now()},
	9:  {ID: 10, Name: "Ten", RecordTime: time.Now(), Label: "do something", Created: time.Now(), Updated: time.Now()},
	10: {ID: 11, Name: "Eleven", RecordTime: time.Now(), Label: "toooo long name, will be truncate", Created: time.Now(), Updated: time.Now()},
	11: {ID: 12, Name: "Twelve", RecordTime: time.Now(), Label: "truncate", Created: time.Now(), Updated: time.Now()},
}
