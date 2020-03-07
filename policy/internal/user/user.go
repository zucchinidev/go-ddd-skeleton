package user

type UserRepository interface {
	ByID(userID int) (*User, error)
}

type UserStatus bool

const (
	Active   UserStatus = true
	Inactive UserStatus = false
)

func (u UserStatus) IntValue() int {
	if u {
		return 1
	}
	return 0
}

type User struct {
	ID    int
	Email string
}
