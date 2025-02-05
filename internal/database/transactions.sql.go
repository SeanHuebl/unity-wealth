// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: transactions.sql

package database

import (
	"context"
)

const createTransaction = `-- name: CreateTransaction :exec
INSERT INTO transactions (
        id,
        user_id,
        transaction_date,
        merchant,
        amount_cents,
        detailed_category_id
    )
VALUES (?1, ?2, ?3, ?4, ?5, ?6)
`

type CreateTransactionParams struct {
	ID                 string
	UserID             string
	TransactionDate    string
	Merchant           string
	AmountCents        int64
	DetailedCategoryID int64
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) error {
	_, err := q.db.ExecContext(ctx, createTransaction,
		arg.ID,
		arg.UserID,
		arg.TransactionDate,
		arg.Merchant,
		arg.AmountCents,
		arg.DetailedCategoryID,
	)
	return err
}

const getDetailedCategoryId = `-- name: GetDetailedCategoryId :one
SELECT id
FROM detailed_categories
WHERE name = ?1
`

func (q *Queries) GetDetailedCategoryId(ctx context.Context, name string) (int64, error) {
	row := q.db.QueryRowContext(ctx, getDetailedCategoryId, name)
	var id int64
	err := row.Scan(&id)
	return id, err
}
