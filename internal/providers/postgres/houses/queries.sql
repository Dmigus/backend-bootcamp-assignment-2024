-- name: createHouse :one
INSERT INTO house (address, year, developer, created_at)
VALUES ($1, $2, $3, clock_timestamp()::timestamp)
RETURNING id, address, year, developer, created_at, update_at;

-- name: getOrder :one
SELECT id, address, year, developer, created_at, update_at
FROM house
WHERE id = $1;