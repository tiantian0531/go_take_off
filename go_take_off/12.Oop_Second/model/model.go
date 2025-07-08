package model

import "strconv"

type User struct {
	name string
	age  int
}

func NewUser(name string, age int) *User {
	user := new(User)
	user.age = age
	user.name = name
	return user
}

func (user *User) GetInfo() string {

	return user.name + strconv.Itoa(user.age)

}