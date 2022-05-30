package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
}

//SQLStore provides all functions to execute db queries and transactions
//Extends the functionality of Queries (individual queries) to run in transaction
type SQLStore struct {
	*Queries
	db *sql.DB //required to create a new transaction
}

//NewStore creates a new Store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

//execTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	//create a transaction
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	//prepare query
	q := New(tx)
	//execute transaction
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

func (store *SQLStore) CreateUserWithTx(ctx context.Context, arg CreateUserParams) (User, error) {
	var result User

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result, err = q.CreateUser(ctx, arg)

		if err != nil {
			return err
		}
		return nil
	})

	return result, err
}
