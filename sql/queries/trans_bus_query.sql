-- name: GetBusById :one
SELECT *
FROM `buses`
WHERE bus_id = ?;

-- name: DeleteBus :exec
UPDATE buses
SET bus_deleted_at = NOW()
WHERE bus_id = ?;

-- name: AddBus :execresult
INSERT INTO buses (
    bus_license_plate, bus_company, bus_price, bus_capacity, bus_created_at, bus_updated_at
) VALUES (
    ?, ?, ?, ?, ?, ?
);

-- name: UpdateBus :execrows
UPDATE buses
SET
    bus_license_plate = ?,
    bus_company = ?,
    bus_price = ?,
    bus_capacity = ?,
    bus_updated_at = NOW()
WHERE
    bus_id = ? AND bus_updated_at = ? AND bus_deleted_at IS NULL;

-- name: DeleteForceBus :exec
DELETE FROM buses WHERE bus_id = ?;