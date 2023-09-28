-- name: CreateTransfer :one
INSERT INTO transfers (
    from_account_id,
    to_account_id,
    amount
) VALUES (
             $1,$2,$3
         ) returning *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1
limit 1;

-- name: ListTransfers :many
SELECT * FROM transfers
ORDER BY id
limit $1
offset $2;

-- name: UpdateTransfer :one
UPDATE transfers
SET amount = $2
where id = $1
returning *;

-- name: DeleteTransfer :exec
DELETE FROM transfers WHERE id = $1;

