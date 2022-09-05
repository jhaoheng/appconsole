package module

import (
	"io/ioutil"

	"github.com/google/uuid"
)

var pic = func(filename string) []byte {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return b
}

var FakeDataUsers = []User{
	0: {ID: 1, Name: "strange_1", MemberID: uuid.New().String(), Picture: pic("./module/fake_pic_1.jpg"), Phone: "1234567", Gender: "woman", PictureFilePath: "./module/fake_pic_1.jpg"},
	1: {ID: 2, Name: "strange_2", MemberID: uuid.New().String(), Picture: pic("./module/fake_pic_2.jpg"), Phone: "1234567", Gender: "woman", PictureFilePath: "./module/fake_pic_2.jpg"},
	2: {ID: 3, Name: "strange_3", MemberID: uuid.New().String(), Picture: pic("./module/fake_pic_3.jpg"), Phone: "1234567", Gender: "woman", PictureFilePath: "./module/fake_pic_3.jpg"},
	3: {ID: 4, Name: "strange_4", MemberID: uuid.New().String(), Picture: pic("./module/fake_pic_4.jpg"), Phone: "1234567", Gender: "woman", PictureFilePath: "./module/fake_pic_4.jpg"},
	4: {ID: 5, Name: "strange_5", MemberID: uuid.New().String(), Picture: pic("./module/fake_pic_5.jpg"), Phone: "1234567", Gender: "woman", PictureFilePath: "./module/fake_pic_5.jpg"},
}
