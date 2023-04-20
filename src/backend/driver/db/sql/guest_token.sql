-- name: CreateGuestToken :one
INSERT INTO guest_token (mail, token, expiration_date) VALUES ($1, $2, $3)
RETURNING *;

-- name: GetGuestTokenByToken :one
SELECT * FROM guest_token WHERE token = $1;

-- name: GetGuestTokenByMailAndToken :one
SELECT * FROM guest_token WHERE mail = $1 AND token = $2;

-- name: DeleteGuestToken :exec
DELETE FROM guest_token WHERE id = $1;
