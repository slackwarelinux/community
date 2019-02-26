package models

import (
	"github.com/satori/go.uuid"
	"math/rand"
	"time"
)

const (
	Male   = 1
	Female = 0
	Unknow = 99
)

type Person struct {
	Id          string
	DisplayName string
	FirstName   string
	LastName    string
	Sex         int
	Birthday    string
}

func genSex() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(1)
}

func (*Person) Bear() *Person {
	baby := new(Person)
	baby.Id = uuid.NewV4().String()
	baby.Sex = genSex()
	return baby
}
