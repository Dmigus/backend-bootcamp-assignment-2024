-- name: createFlat :one
INSERT INTO flat (house_id, price, rooms, status)
VALUES ($1, $2, $3, 'Created')
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
WHERE house_id = $1 AND status = 'Approved';

-- name: getFlatForUpdate :one
SELECT id, house_id, price, rooms, status
FROM flat
WHERE id = $1 FOR UPDATE;

-- name: updateStatus :one
UPDATE flat
SET status = $2
WHERE id = $1
RETURNING id, house_id, price, rooms, status;