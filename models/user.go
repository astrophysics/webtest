package models

type User struct {
	ID        int64
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID int64 = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}
