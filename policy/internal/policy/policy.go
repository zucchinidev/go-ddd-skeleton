package policy

import (
	"database/sql"
	"github.com/zucchinidev/go-ddd-skeleton/policy/internal/user"
)

// domain
// I assume a small intrusion of the infrastructure component (sql) in order to simplify the programming and
// trie apply Don't Repeat Your-self
type Policy struct {
	Identifier int
	User       *user.User
}

type PolicyStatus int

const (
	OpenPolicy   PolicyStatus = 100
	ClosedPolicy PolicyStatus = 200
)

type PolicyRepository interface {
	OpenPolicies() ([]Policy, error)
	CloseByUserID(tx TransactionManager, userID int) error
	BlockUser(tx TransactionManager, userID int) error
	WithTransaction(fn TxFn) (err error)
}

// TransactionManager is an interface that models the standard transaction in
// `database/sql`.
//
// To ensure `TxFn` function cannot commit or rollback a transaction (which is
// handled by `WithTransaction`), those methods are not included here.
type TransactionManager interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// A TxFn is a function that will be called with an initialized `TransactionManager` object
// that can be used for executing statements and queries against a database.
type TxFn func(TransactionManager) error
