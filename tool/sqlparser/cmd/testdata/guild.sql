-- name: CreateGuildWithRegistering :one
INSERT INTO guild (name, status) VALUES ($1, 1)
RETURNING *;

-- name: UpdateGuildWithRegistered :one
UPDATE guild SET status = 2 WHERE id = $1
RETURNING *;

-- name: GetGuildByID :one
SELECT * FROM guild g INNER JOIN guild_owner_relation r ON r.guild_id = g.id WHERE g.id = $1;
