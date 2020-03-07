package mysql

import (
	"github.com/zucchinidev/go-ddd-skeleton/policy/internal/policy"
	"github.com/zucchinidev/go-ddd-skeleton/policy/internal/user"
	"github.com/zucchinidev/go-ddd-skeleton/policy/shared/adapters/store"
	"github.com/zucchinidev/go-ddd-skeleton/policy/shared/errors"
)

type policyRepository struct {
	conn *store.Storage
}

func NewPolicyRepository(conn *store.Storage) *policyRepository {
	return &policyRepository{conn: conn}
}

func (c *policyRepository) OpenPolicies() ([]policy.Policy, error) {
	pp := []policy.Policy{}
	query := c.getOpenPoliciesSQL()
	stmt, err := c.conn.DB.Prepare(query)
	if err != nil {
		return nil, errors.WrapPolicySearch(err, "error preparing query %s", query)
	}
	defer stmt.Close()

	cur, err := stmt.Query(int(policy.OpenPolicy))

	if err != nil {
		return nil, errors.WrapPolicySearch(err, "error executing query %s", query)
	}
	defer cur.Close()

	for cur.Next() {
		p := policy.Policy{User: &user.User{}}
		if err = cur.Scan(&p.Identifier, &p.User.ID, &p.User.Email); err != nil {
			return nil, errors.WrapPolicySearch(err, "error scanning values in query %s", query)
		}
		pp = append(pp, p)
	}

	if err = cur.Err(); err != nil {
		return nil, errors.WrapPolicySearch(err, "error with the cursor in query %s", query)
	}
	if len(pp) == 0 {
		return pp, errors.NewPolicyNotFound("policies not found")
	}
	return pp, nil
}

func (c *policyRepository) CloseByUserID(tx policy.TransactionManager, userID int) error {
	_, err := tx.Exec(c.closePoliciesByUserIDSQL(), policy.ClosedPolicy, userID)
	return err
}

func (c *policyRepository) BlockUser(tx policy.TransactionManager, userID int) error {
	_, err := tx.Exec(c.blockUserSQL(), user.Inactive.IntValue(), userID)
	return err
}

// WithTransaction creates a new transaction and handles rollback/commit based on the
// error object returned by the `TxFn`
func (c *policyRepository) WithTransaction(fn policy.TxFn) (err error) {
	tx, err := c.conn.DB.Begin()
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and re-panic
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			_ = tx.Rollback()
		} else {
			// all good, commit
			err = tx.Commit()
		}
	}()

	err = fn(tx)
	return err
}

func (c *policyRepository) getOpenPoliciesSQL() string {
	return `SELECT p.id, u.id, u.email FROM policy p inner join user u on p.user_id = u.id where p.status_id = ?`
}

func (c *policyRepository) closePoliciesByUserIDSQL() string {
	return `UPDATE policy p SET status_id = ? where p.user_id = ?`
}
func (c *policyRepository) blockUserSQL() string {
	return `UPDATE user u SET active = ? where u.id = ?`
}
