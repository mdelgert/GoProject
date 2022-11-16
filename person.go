//https://ieftimov.com/posts/testing-in-go-go-test/#:~:text=To%20run%20your%20tests%20in,prints%20the%20complete%20test%20output.

package main

import "errors"

type Person struct {
	age int
}

func NewPerson(age int) (*Person, error) {
	if age < 1 {
		return nil, errors.New("A person is at least 1 years old")
	}

	return &Person{
		age: age,
	}, nil
}

func (p *Person) older(other *Person) bool {
	return p.age > other.age
}
