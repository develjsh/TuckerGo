package main

type Stringer interface {
	String() string
}

type Student struct {
	Name string
}

func (s *Student) String() string {
	return s.Name
}

type User struct {
	Name string
	Age int
}

func (u User) String() string {
	return u.Name
}