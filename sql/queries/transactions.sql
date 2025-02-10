-- name: CreateTransaction :exec
INSERT INTO transactions (
        id,
        user_id,
        transaction_date,
        merchant,
        amount_cents,
        detailed_category_id
    )
VALUES (?1, ?2, ?3, ?4, ?5, ?6);
-- name: GetDetailedCategoryId :one
SELECT id
FROM detailed_categories
WHERE name = ?1;
-- name: UpdateTransactionByID :one
UPDATE transactions
SET transaction_date = ?1,
    merchant = ?2,
    amount_cents = ?3,
    detailed_category_id = ?4,
    updated_at = ?5
WHERE id = ?6
RETURNING id,
    transaction_date,
    merchant,
    amount_cents,
    detailed_category_id;
-- name: GetPrimaryCategories :many
SELECT *
FROM primary_categories;
-- name: GetDetailedCategories :many
SELECT *
FROM detailed_categories;
-- name: DeleteTransactionById :exec
DELETE FROM transactions
WHERE id = ?1
    AND user_id = ?2;
-- name: GetUserTransactionsFirstPage :many
SELECT id,
    user_id,
    transaction_date,
    merchant,
    amount_cents,
    detailed_category_id
FROM transactions
WHERE user_id = ?1
ORDER BY transaction_date DESC,
    id ASC
LIMIT ?2;
-- name: GetUserTransactionsPaginated :many
SELECT id,
    user_id,
    transaction_date,
    merchant,
    amount_cents,
    detailed_category_id
FROM transactions
WHERE user_id = ?1
    AND (
        transaction_date < ?2
        OR (
            transaction_date = ?3
            AND id < ?4
        )
    )
ORDER BY transaction_date DESC,
    id ASC
LIMIT ?5;