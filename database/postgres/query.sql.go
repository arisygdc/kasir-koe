// Code generated by sqlc. DO NOT EDIT.
// source: query.sql

package postgres

import (
	"context"
)

const createKategori = `-- name: CreateKategori :exec
INSERT INTO kategori (id, kategori) VALUES (DEFAULT, $1)
`

func (q *Queries) CreateKategori(ctx context.Context, kategori string) error {
	_, err := q.db.ExecContext(ctx, createKategori, kategori)
	return err
}

const createMeja = `-- name: CreateMeja :exec
INSERT INTO meja (nomor, dibuat_pada) VALUES ($1, DEFAULT)
`

func (q *Queries) CreateMeja(ctx context.Context, nomor int32) error {
	_, err := q.db.ExecContext(ctx, createMeja, nomor)
	return err
}

const createMenu = `-- name: CreateMenu :exec
INSERT INTO menu (id, kategori_id, menu, harga) VALUES (DEFAULT, $1, $2, $3)
`

type CreateMenuParams struct {
	KategoriID int32  `json:"kategori_id"`
	Menu       string `json:"menu"`
	Harga      int32  `json:"harga"`
}

func (q *Queries) CreateMenu(ctx context.Context, arg CreateMenuParams) error {
	_, err := q.db.ExecContext(ctx, createMenu, arg.KategoriID, arg.Menu, arg.Harga)
	return err
}
