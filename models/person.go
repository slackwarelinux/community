package models

import (
	"github.com/satori/go.uuid"
	"math/rand"
	"time"
)

const (
	Male   = 1
	Female = 0
)

type Person struct {
	Id        string
	Name      string
	FirstName string
	LastName  string
	Sex       int
	Birthday  string
}

func genSex() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(1)
}

func (*Person) Bear() *Person {
	uid, _ := uuid.NewV4()
	baby := &Person{
		Id:   uid.String(),
		Sex:  genSex(),
		Name: "未定义",
	}
	return baby
}
