-- name: CreateGuestToken :one
INSERT INTO guest_token (mail, token, expiration_date) VALUES ($1, $2, $3)
RETURNING *;
