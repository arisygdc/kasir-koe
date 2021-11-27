// Code generated by sqlc. DO NOT EDIT.

package database

import (
	"database/sql"
)

type DetailPesanan struct {
	PesananID int32 `json:"pesanan_id"`
	MenuID    int32 `json:"menu_id"`
	Jumlah    int32 `json:"jumlah"`
}

type Kategori struct {
	ID       int32  `json:"id"`
	Kategori string `json:"kategori"`
}

type Meja struct {
	Nomor      int32        `json:"nomor"`
	DibuatPada sql.NullTime `json:"dibuat_pada"`
}

type Menu struct {
	ID         int32 `json:"id"`
	KategoriID int32 `json:"kategori_id"`
}

type Pesanan struct {
	ID          int32        `json:"id"`
	Kode        string       `json:"kode"`
	DipesanPada sql.NullTime `json:"dipesan_pada"`
}
