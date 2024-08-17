-- name: createHouse :one
INSERT INTO house (address, year, developer, created_at)
VALUES ($1, $2, $3, clock_timestamp()::timestamp)
RETURNING id, address, year, developer, created_at, update_at;

-- name: getOrder :one
SELECT id, address, year, developer, created_at, update_at
FROM house
WHERE id = $1;

-- name: createFlat :one
INSERT INTO flat (house_id, price, rooms, status)
VALUES ($1, $2, $3, 'created')
RETURNING id, house_id, price, rooms, status;

-- name: updateFlat :one
UPDATE flat
SET price = $2, rooms= $3, status = $4
WHERE id = $1
RETURNING id, house_id, price, rooms, status;

-- name: getFlats :many
SELECT id, house_id, price, rooms, status
FROM flat
WHERE house_id = $1;

-- name: getApprovedFlats :many
SELECT id, house_id, price, rooms, status
FROM flat
WHERE house_id = $1 AND status = 'approved';