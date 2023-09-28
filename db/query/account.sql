-- name: CreateAccount :one
INSERT INTO accounts (
owner,
balance,
currency
) VALUES (
$1,$2,$3
) returning *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1
limit 1;

-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE id = $1
limit 1
for update;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
limit $1
offset $2;

-- name: UpdateAccount :one
UPDATE accounts
SET balance = $2
where id = $1
returning *;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;

