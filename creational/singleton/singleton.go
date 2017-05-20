package singleton

type Singleton interface {
	AddOne() int
}

type singleton struct {
	cout int
}

var instance *singleton

func GetInstance() Singleton {
	return nil
}

func (s *singleton) AddOne() int {
	return 0
}
