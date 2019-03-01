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

func (pool *Pool) Apply() (*Person, error) {
	if pool.free <= pool.min {
		childbirth(pool)
	}
	person := pool.persons[0]
	person.Status = StatusWaiting
	pool.free--
	return person, nil
}

func (pool *Pool) GetStatus() (total int, free int) {
	total = len(pool.persons)
	free = pool.free
	return
}

// 获取状态
func getStatus(pool *Pool, status int) (int, []*Person) {
	var count int = 0
	var persons []*Person
	for i := 0; i < len(pool.persons); i++ {
		if pool.persons[i].Status == status {
			count++
			persons = append(persons, pool.persons[i])
		}
	}
	return count, persons
}

// 生产
func childbirth(pool *Pool) {
	parent := Person{}
	for i := 0; i < pool.min; i++ {
		baby := parent.Bear()
		pool.persons = append(pool.persons, baby)
		pool.free++
	}
}
