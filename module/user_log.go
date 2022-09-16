package module

import (
	"strings"
	"time"
)

type UserLog struct {
	ID         int
	Name       string
	RecordTime time.Time
	Label      string
	Created    time.Time
	Updated    time.Time
}

func NewUserLog() *UserLog {
	return &UserLog{}
}

func (module *UserLog) GetAll() []UserLog {
	return FakeUserLogs
}

func (module *UserLog) SearchNameLike(name string) ([]UserLog, error) {
	output := []UserLog{}
	name = strings.ToLower(name)
	for _, userlog := range FakeUserLogs {
		if strings.Contains(strings.ToLower(userlog.Name), name) {
			output = append(output, userlog)
		}
	}
	return output, nil
}
