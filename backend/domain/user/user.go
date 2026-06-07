package user

type User struct {
	id    int64
	name  string
	email string
}

type UserWithPurchases struct {
	User
	purchases int
}

type UserByMostExpensiveProduct struct {
	User
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

func (u *UserWithPurchases) Purchases() int {
	return u.purchases
}

func NewUser(id int64, name string, email string) *User {
	return &User{id: id, name: name, email: email}
}

func NewUserWithPurchases(id int64, name string, email string, purchases int) *UserWithPurchases {
	return &UserWithPurchases{
		User:      User{id: id, name: name, email: email},
		purchases: purchases,
	}
}

func NewUserByMostExpensiveProduct(id int64, name string, email string) *UserByMostExpensiveProduct {
	return &UserByMostExpensiveProduct{
		User: User{id: id, name: name, email: email},
	}
}
