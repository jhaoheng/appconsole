package module

type IUser interface {
	SetName(name string) IUser
	SetMemberID(mid string) IUser
	//
	Create(newuser *User) bool
	GetByID(id int) User
	GetAll() ([]User, error)
	List(num int, page int) []User
	Del(id int) error
	Count() int
}

type User struct {
	ID              int
	MemberID        string
	Name            string
	Picture         []byte
	PictureFilePath string
	Phone           string
	Gender          string
}

func NewUser() IUser {
	return &User{}
}

func (u *User) SetName(name string) IUser {
	u.Name = name
	return u
}

func (u *User) SetMemberID(mid string) IUser {
	u.MemberID = mid
	return u
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

func (u *User) GetAll() ([]User, error) {
	return FakeDataUsers, nil
}

func (u *User) List(num int, page int) []User {
	return FakeDataUsers
}

func (u *User) Del(id int) error {
	NewFakeDataUsers := []User{}
	for _, v := range FakeDataUsers {
		if v.ID != id {
			NewFakeDataUsers = append(NewFakeDataUsers, v)
		}
	}
	FakeDataUsers = NewFakeDataUsers
	return nil
}

func (u *User) Count() int {
	return len(FakeDataUsers)
}
