-- name: CreateGuestToken :one
INSERT INTO guest_token (guild_id, mail, token, expiration_date, accepted_number) VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetGuestTokenByToken :one
SELECT * FROM guest_token WHERE token = $1;

-- name: GetGuestTokenByMailAndToken :one
SELECT * FROM guest_token WHERE mail = $1 AND token = $2;

-- name: DeleteGuestToken :exec
DELETE FROM guest_token WHERE id = $1;
