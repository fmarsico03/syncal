package models

type User struct {
	name     string
	lastname string
	mail     string
}

func (u User) Lastname() string {
	return u.lastname
}

func (u User) SetLastname(lastname string) {
	u.lastname = lastname
}

func (u User) Name() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u User) Mail() string {
	return u.mail
}

func (u *User) SetMail(mail string) {
	u.mail = mail
}

func NewUser(name string, lastname string, mail string) *User {
	return &User{name: name, lastname: lastname, mail: mail}
}
