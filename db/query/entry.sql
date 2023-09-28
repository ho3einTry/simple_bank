-- name: CreatEntry :one
INSERT INTO entries (
    account_id,
    amount
) VALUES (
             $1,$2
         ) returning *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1
limit 1;


-- name: ListEntryByAccountId :many
SELECT * FROM entries
where account_id = $1
ORDER BY id
limit $2
offset $3;

-- name: ListEntry :many
SELECT * FROM entries
ORDER BY id
limit $1
offset $2;

-- name: UpdateEntry :one
UPDATE entries
SET amount = $2
where id = $1
returning *;

-- name: DeleteEntry :exec
DELETE FROM entries WHERE id = $1;

