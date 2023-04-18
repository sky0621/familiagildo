-- name: CreateGuildOwnerRelation :one
INSERT INTO guild_owner_relation (guild_id, owner_id) VALUES ($1, $2)
RETURNING *;
