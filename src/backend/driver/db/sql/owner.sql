-- name: CreateOwner :one
INSERT INTO owner (name, mail, login_id, password) VALUES ($1, $2, $3, $4)
RETURNING *;
