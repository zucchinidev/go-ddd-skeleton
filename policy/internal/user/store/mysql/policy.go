package mysql

import (
	"database/sql"
	"github.com/zucchinidev/go-ddd-skeleton/policy/internal/user"
	"github.com/zucchinidev/go-ddd-skeleton/policy/shared/adapters/store"
	"github.com/zucchinidev/go-ddd-skeleton/policy/shared/errors"
)

type userRepository struct {
	conn *store.Storage
}

func NewUserRepository(conn *store.Storage) *userRepository {
	return &userRepository{conn: conn}
}

func (c *userRepository) ByID(userID int) (*user.User, error) {
	query := c.getUserByIDSQL()
	stmt, err := c.conn.DB.Prepare(query)
	if err != nil {
		return nil, errors.WrapUserSearch(err, "error preparing query %s", query)
	}
	defer stmt.Close()

	var u user.User
	if err := stmt.QueryRow(userID).Scan(&u.ID, &u.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.WrapUserNotFound(err, "user with id %d not found", userID)
		}
	}
	return &u, nil
}

func (c *userRepository) getUserByIDSQL() string {
	return `SELECT u.id, u.email FROM user u where u.id = ?`
}
