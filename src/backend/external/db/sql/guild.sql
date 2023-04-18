-- name: CreateGuildWithRegistering :one
INSERT INTO guild (name, status) VALUES ($1, 1)
RETURNING *;
