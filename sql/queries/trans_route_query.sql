-- name: GetRouteById :one
SELECT *
FROM `routes`
WHERE route_id = ?;

-- name: DeleteRoute :exec
UPDATE routes
SET route_deleted_at = NOW()
WHERE route_id = ?;

-- name: AddRoute :execresult
INSERT INTO routes (
    route_start_location, route_end_location, route_estimated_duration, route_created_at, route_updated_at
) VALUES (
    ?, ?, ?, ?, ?
);

-- name: UpdateRoute :execrows
UPDATE routes
SET route_updated_at = NOW(), route_start_location = ?, route_end_location = ?, route_estimated_duration = ?
WHERE route_id = ? AND route_updated_at = ? AND route_deleted_at IS NULL;

-- name: DeleteForceRoute :exec
DELETE FROM routes WHERE route_id = ?;