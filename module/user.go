package module

var FakeDataUsers = []User{}

type IUser interface {
	Create(newuser *User) bool
	GetByID(id int) User
	List(num int, page int) []User
}

type User struct {
	ID      int
	Name    string
	Picture []byte
}

func NewUser() IUser {
	return &User{}
}

func (u *User) Create(newuser *User) bool {
	FakeDataUsers = append(FakeDataUsers, *newuser)
	return true
}

func (u *User) GetByID(id int) User {
	output := User{}
	for _, fakeuser := range FakeDataUsers {
		if fakeuser.ID == id {
			output = fakeuser
			break
		}
	}
	return output
}

func (u *User) List(num int, page int) []User {
	return FakeDataUsers
}
