package models

type Pool struct {
	min     int
	free    int
	persons []*Person
}

func (pool *Pool) Init(min_ int) {
	pool.min = min_
	parent := Person{}
	var persons []*Person
	for i := 0; i < pool.min; i++ {
		baby := parent.Bear()
		persons = append(persons, baby)
	}
	pool.persons = persons
	pool.free = len(persons)
}

func (pool *Pool) Request() *Person {
	if pool.free < pool.min {
		childbirth(pool)
	}
	person := pool.persons[0]
	pool.free--
	return person
}

func (pool *Pool) GetStatus() (total int, free int) {
	total = len(pool.persons)
	free = pool.free
	return
}

func childbirth(pool *Pool) {
	parent := Person{}
	for i := 0; i < pool.min; i++ {
		baby := parent.Bear()
		pool.persons = append(pool.persons, baby)
		pool.free++
	}
}
