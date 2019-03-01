package models

import (
	"github.com/satori/go.uuid"
	"math/rand"
	"time"
)

const (
	SexMale   = 1
	SexFemale = 0
)

const (
	StatusFree = iota
	StatusWaiting
	StatusRunning
	StatusDeath
)

type Person struct {
	Id          string
	Code        string
	Name        string
	Sex         int
	Birthday    time.Time
	DateOfDeath time.Time
	Status      int
	Applicant   string
}

func genSex() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(1)
}

func (*Person) Bear() *Person {
	uid, _ := uuid.NewV4()
	baby := &Person{
		Id:       uid.String(),
		Sex:      genSex(),
		Name:     "未定义",
		Birthday: time.Now(),
		Status:   StatusFree,
	}
	return baby
}

func (person *Person) Death() {
	person.DateOfDeath = time.Now()
	person.Status = StatusDeath
}
