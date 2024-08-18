-- name: register :exec
INSERT INTO "user" (id, email, salt, password_hash, role)
VALUES ($1, $2, $3, $4, $5);

-- name: getAuthData :one
SELECT salt, password_hash, role
FROM "user"
WHERE id = $1;
