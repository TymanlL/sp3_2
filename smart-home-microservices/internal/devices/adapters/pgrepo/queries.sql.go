// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package pgrepo

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createDevice = `-- name: CreateDevice :exec
insert into devices (id, type, name, online, on_off, user_id, home_id)
values ($1, $2, $3, $4, $5, $6, $7)
`

type CreateDeviceParams struct {
	ID     string
	Type   string
	Name   string
	Online bool
	OnOff  bool
	UserID string
	HomeID pgtype.Text
}

func (q *Queries) CreateDevice(ctx context.Context, arg CreateDeviceParams) error {
	_, err := q.db.Exec(ctx, createDevice,
		arg.ID,
		arg.Type,
		arg.Name,
		arg.Online,
		arg.OnOff,
		arg.UserID,
		arg.HomeID,
	)
	return err
}

const deleteDeviceByID = `-- name: DeleteDeviceByID :exec
delete
from devices
where id = $1
  and user_id = $2
`

type DeleteDeviceByIDParams struct {
	ID     string
	UserID string
}

func (q *Queries) DeleteDeviceByID(ctx context.Context, arg DeleteDeviceByIDParams) error {
	_, err := q.db.Exec(ctx, deleteDeviceByID, arg.ID, arg.UserID)
	return err
}

const deleteOutboxMessagesDeviceCreated = `-- name: DeleteOutboxMessagesDeviceCreated :exec
delete
from devices_created_outbox
`

func (q *Queries) DeleteOutboxMessagesDeviceCreated(ctx context.Context) error {
	_, err := q.db.Exec(ctx, deleteOutboxMessagesDeviceCreated)
	return err
}

const deleteOutboxMessagesDeviceDeleted = `-- name: DeleteOutboxMessagesDeviceDeleted :exec
delete
from devices_deleted_outbox
`

func (q *Queries) DeleteOutboxMessagesDeviceDeleted(ctx context.Context) error {
	_, err := q.db.Exec(ctx, deleteOutboxMessagesDeviceDeleted)
	return err
}

const deleteOutboxMessagesDeviceUpdated = `-- name: DeleteOutboxMessagesDeviceUpdated :exec
delete
from devices_updated_outbox
`

func (q *Queries) DeleteOutboxMessagesDeviceUpdated(ctx context.Context) error {
	_, err := q.db.Exec(ctx, deleteOutboxMessagesDeviceUpdated)
	return err
}

const getDeviceByID = `-- name: GetDeviceByID :one
select id, type, name, online, on_off, user_id, home_id, created_at
from devices
where id = $1
  and user_id = $2
`

type GetDeviceByIDParams struct {
	ID     string
	UserID string
}

func (q *Queries) GetDeviceByID(ctx context.Context, arg GetDeviceByIDParams) (Device, error) {
	row := q.db.QueryRow(ctx, getDeviceByID, arg.ID, arg.UserID)
	var i Device
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Name,
		&i.Online,
		&i.OnOff,
		&i.UserID,
		&i.HomeID,
		&i.CreatedAt,
	)
	return i, err
}

const getDeviceByIDForUpdate = `-- name: GetDeviceByIDForUpdate :one
select id, type, name, online, on_off, user_id, home_id, created_at
from devices
where id = $1
  and user_id = $2
    for update
`

type GetDeviceByIDForUpdateParams struct {
	ID     string
	UserID string
}

func (q *Queries) GetDeviceByIDForUpdate(ctx context.Context, arg GetDeviceByIDForUpdateParams) (Device, error) {
	row := q.db.QueryRow(ctx, getDeviceByIDForUpdate, arg.ID, arg.UserID)
	var i Device
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Name,
		&i.Online,
		&i.OnOff,
		&i.UserID,
		&i.HomeID,
		&i.CreatedAt,
	)
	return i, err
}

const getOutboxMessagesDeviceCreated = `-- name: GetOutboxMessagesDeviceCreated :many
select device_id
from devices_created_outbox
`

func (q *Queries) GetOutboxMessagesDeviceCreated(ctx context.Context) ([]pgtype.Text, error) {
	rows, err := q.db.Query(ctx, getOutboxMessagesDeviceCreated)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []pgtype.Text
	for rows.Next() {
		var device_id pgtype.Text
		if err := rows.Scan(&device_id); err != nil {
			return nil, err
		}
		items = append(items, device_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOutboxMessagesDeviceDeleted = `-- name: GetOutboxMessagesDeviceDeleted :many
select device_id
from devices_deleted_outbox
`

func (q *Queries) GetOutboxMessagesDeviceDeleted(ctx context.Context) ([]pgtype.Text, error) {
	rows, err := q.db.Query(ctx, getOutboxMessagesDeviceDeleted)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []pgtype.Text
	for rows.Next() {
		var device_id pgtype.Text
		if err := rows.Scan(&device_id); err != nil {
			return nil, err
		}
		items = append(items, device_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOutboxMessagesDeviceUpdated = `-- name: GetOutboxMessagesDeviceUpdated :many
select device_id
from devices_updated_outbox
`

func (q *Queries) GetOutboxMessagesDeviceUpdated(ctx context.Context) ([]pgtype.Text, error) {
	rows, err := q.db.Query(ctx, getOutboxMessagesDeviceUpdated)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []pgtype.Text
	for rows.Next() {
		var device_id pgtype.Text
		if err := rows.Scan(&device_id); err != nil {
			return nil, err
		}
		items = append(items, device_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listHomeDevices = `-- name: ListHomeDevices :many
select id, type, name, online, on_off, user_id, home_id, created_at
from devices
where home_id = $1
  and user_id = $2
`

type ListHomeDevicesParams struct {
	HomeID pgtype.Text
	UserID string
}

func (q *Queries) ListHomeDevices(ctx context.Context, arg ListHomeDevicesParams) ([]Device, error) {
	rows, err := q.db.Query(ctx, listHomeDevices, arg.HomeID, arg.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Device
	for rows.Next() {
		var i Device
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.Name,
			&i.Online,
			&i.OnOff,
			&i.UserID,
			&i.HomeID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUserHomes = `-- name: ListUserHomes :many
select id, name, user_id, created_at
from homes
where user_id = $1
`

func (q *Queries) ListUserHomes(ctx context.Context, userID string) ([]Home, error) {
	rows, err := q.db.Query(ctx, listUserHomes, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Home
	for rows.Next() {
		var i Home
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.UserID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const saveDeviceCreatedToOutbox = `-- name: SaveDeviceCreatedToOutbox :exec
insert into devices_created_outbox (device_id)
values ($1)
`

func (q *Queries) SaveDeviceCreatedToOutbox(ctx context.Context, deviceID pgtype.Text) error {
	_, err := q.db.Exec(ctx, saveDeviceCreatedToOutbox, deviceID)
	return err
}

const saveDeviceDeletedToOutbox = `-- name: SaveDeviceDeletedToOutbox :exec
insert into devices_deleted_outbox (device_id)
values ($1)
`

func (q *Queries) SaveDeviceDeletedToOutbox(ctx context.Context, deviceID pgtype.Text) error {
	_, err := q.db.Exec(ctx, saveDeviceDeletedToOutbox, deviceID)
	return err
}

const saveDeviceUpdatedToOutbox = `-- name: SaveDeviceUpdatedToOutbox :exec
insert into devices_updated_outbox (device_id)
values ($1)
`

func (q *Queries) SaveDeviceUpdatedToOutbox(ctx context.Context, deviceID pgtype.Text) error {
	_, err := q.db.Exec(ctx, saveDeviceUpdatedToOutbox, deviceID)
	return err
}

const updateDevice = `-- name: UpdateDevice :exec
update devices
set name   = $1,
    online = $2,
    on_off = $3
where id = $4
  and user_id = $5
`

type UpdateDeviceParams struct {
	Name   string
	Online bool
	OnOff  bool
	ID     string
	UserID string
}

func (q *Queries) UpdateDevice(ctx context.Context, arg UpdateDeviceParams) error {
	_, err := q.db.Exec(ctx, updateDevice,
		arg.Name,
		arg.Online,
		arg.OnOff,
		arg.ID,
		arg.UserID,
	)
	return err
}

const checkHomeBelongsToUser = `-- name: checkHomeBelongsToUser :one
select exists(select id, name, user_id, created_at from homes where user_id = $1 and id = $2)
`

type checkHomeBelongsToUserParams struct {
	UserID string
	ID     string
}

func (q *Queries) checkHomeBelongsToUser(ctx context.Context, arg checkHomeBelongsToUserParams) (bool, error) {
	row := q.db.QueryRow(ctx, checkHomeBelongsToUser, arg.UserID, arg.ID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}