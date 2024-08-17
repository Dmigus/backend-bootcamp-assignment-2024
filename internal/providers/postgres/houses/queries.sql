-- name: createHouse :one
INSERT INTO house (address, year, developer, created_at)
VALUES ($1, $2, $3, clock_timestamp()::timestamp)
RETURNING id, address, year, developer, created_at, update_at;

-- name: houseUpdated :exec
UPDATE house
SET update_at = clock_timestamp()::timestamp
WHERE id = $1;

-- name: checkHouseExistence :exec
SELECT 1
FROM house
WHERE id = $1;

