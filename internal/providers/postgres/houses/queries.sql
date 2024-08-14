-- name: createHouse :one
INSERT INTO house (address, year, developer)
VALUES ($1, $2, $3)
RETURNING id, address, year, developer, created_at, update_at;

-- name: getOrder :one
SELECT id, address, year, developer, created_at, update_at
FROM house
WHERE id = $1;