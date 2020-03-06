package user

type UserRepository interface {
	ByID(userID int) (*User, error)
}

type User struct {
	ID    int
	Email string
}
