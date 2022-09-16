package module

import "strings"

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
	SearchNameLike(name string) ([]User, error)
	SearchMemberIDLike(mid string) ([]User, error)
}

type User struct {
	ID       int
	MemberID string
	Name     string
	Picture  []byte
	Phone    string
	Gender   string
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
	FakeUsers = append(FakeUsers, *newuser)
	return true
}

func (u *User) GetByID(id int) User {
	output := User{}
	for _, fakeuser := range FakeUsers {
		if fakeuser.ID == id {
			output = fakeuser
			break
		}
	}
	return output
}

func (u *User) GetAll() ([]User, error) {
	return FakeUsers, nil
}

func (u *User) List(num int, page int) []User {
	return FakeUsers
}

func (u *User) Del(id int) error {
	NewFakeUsers := []User{}
	for _, v := range FakeUsers {
		if v.ID != id {
			NewFakeUsers = append(NewFakeUsers, v)
		}
	}
	FakeUsers = NewFakeUsers
	return nil
}

func (u *User) Count() int {
	return len(FakeUsers)
}

func (u *User) SearchNameLike(name string) ([]User, error) {
	output := []User{}
	name = strings.ToLower(name)
	for _, user := range FakeUsers {
		if strings.Contains(strings.ToLower(user.Name), name) {
			output = append(output, user)
		}
	}
	return output, nil
}

func (u *User) SearchMemberIDLike(mid string) ([]User, error) {
	output := []User{}
	for _, user := range FakeUsers {
		if strings.Contains(user.MemberID, mid) {
			output = append(output, user)
		}
	}
	return output, nil
}
