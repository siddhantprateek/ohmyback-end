package main

import "fmt"

/*
Builder pattern separates the construction of complex
object from its representation, allowing the same construction process to create
different representation.
*/

type User struct {
	Name   string
	Role   string
	Salary int
}

type UserBuilder struct {
	User
}

func (ub *UserBuilder) SetName(name string) *UserBuilder {
	ub.User.Name = name
	return ub
}

func (ub *UserBuilder) SetRole(role string) *UserBuilder {
	ub.User.Role = role
	return ub
}

func (ub *UserBuilder) SetSalary(sal int) *UserBuilder {
	ub.User.Salary = sal
	return ub
}

func (ub *UserBuilder) Build() User {
	return ub.User
}

func main() {
	ub := &UserBuilder{}

	user := ub.
		SetName("John").
		SetRole("Developer").
		SetSalary(1000000).
		Build()

	fmt.Println(user.Name, user.Role, user.Salary)
}
