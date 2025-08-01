// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: trans_bus_query.sql

package database

import (
	"context"
	"database/sql"
)

const addBus = `-- name: AddBus :execresult
INSERT INTO buses (
    bus_license_plate, bus_company, bus_price, bus_capacity, bus_created_at, bus_updated_at
) VALUES (
    ?, ?, ?, ?, ?, ?
)
`

type AddBusParams struct {
	BusLicensePlate string
	BusCompany      string
	BusPrice        string
	BusCapacity     int32
	BusCreatedAt    sql.NullTime
	BusUpdatedAt    sql.NullTime
}

func (q *Queries) AddBus(ctx context.Context, arg AddBusParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, addBus,
		arg.BusLicensePlate,
		arg.BusCompany,
		arg.BusPrice,
		arg.BusCapacity,
		arg.BusCreatedAt,
		arg.BusUpdatedAt,
	)
}

const deleteBus = `-- name: DeleteBus :exec
UPDATE buses
SET bus_deleted_at = NOW()
WHERE bus_id = ?
`

func (q *Queries) DeleteBus(ctx context.Context, busID int32) error {
	_, err := q.db.ExecContext(ctx, deleteBus, busID)
	return err
}

const deleteForceBus = `-- name: DeleteForceBus :exec
DELETE FROM buses WHERE bus_id = ?
`

func (q *Queries) DeleteForceBus(ctx context.Context, busID int32) error {
	_, err := q.db.ExecContext(ctx, deleteForceBus, busID)
	return err
}

const getBusById = `-- name: GetBusById :one
SELECT bus_id, bus_license_plate, bus_company, bus_price, bus_capacity, bus_created_at, bus_updated_at, bus_deleted_at
FROM ` + "`" + `buses` + "`" + `
WHERE bus_id = ?
`

func (q *Queries) GetBusById(ctx context.Context, busID int32) (Bus, error) {
	row := q.db.QueryRowContext(ctx, getBusById, busID)
	var i Bus
	err := row.Scan(
		&i.BusID,
		&i.BusLicensePlate,
		&i.BusCompany,
		&i.BusPrice,
		&i.BusCapacity,
		&i.BusCreatedAt,
		&i.BusUpdatedAt,
		&i.BusDeletedAt,
	)
	return i, err
}

const updateBus = `-- name: UpdateBus :execrows
UPDATE buses
SET
    bus_license_plate = ?,
    bus_company = ?,
    bus_price = ?,
    bus_capacity = ?,
    bus_updated_at = NOW()
WHERE
    bus_id = ? AND bus_updated_at = ? AND bus_deleted_at IS NULL
`

type UpdateBusParams struct {
	BusLicensePlate string
	BusCompany      string
	BusPrice        string
	BusCapacity     int32
	BusID           int32
	BusUpdatedAt    sql.NullTime
}

func (q *Queries) UpdateBus(ctx context.Context, arg UpdateBusParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, updateBus,
		arg.BusLicensePlate,
		arg.BusCompany,
		arg.BusPrice,
		arg.BusCapacity,
		arg.BusID,
		arg.BusUpdatedAt,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
