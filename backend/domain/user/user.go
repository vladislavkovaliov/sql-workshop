package user

type User struct {
	id    int64
	name  string
	email string
}

func (u *User) ID() int64 {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() string {
	return u.email
}

func NewUser(id int64, name string, email string) *User {
	return &User{id: id, name: name, email: email}
}
