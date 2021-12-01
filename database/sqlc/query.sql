-- name: CreateMeja :exec
INSERT INTO meja (nomor, dibuat_pada) VALUES ($1, DEFAULT);

-- name: CreateKategori :exec
INSERT INTO kategori (id, kategori) VALUES (DEFAULT, $1);

-- name: CreateMenu :exec
INSERT INTO menu (id, kategori_id, menu, harga) VALUES (DEFAULT, $1, $2, $3);

-- name: CreatePesanan :exec
INSERT INTO pesanan (id, kode, meja_nomor, dipesan_pada) VALUES (DEFAULT, $1, $2, DEFAULT);

-- name: GetPesananID :one
SELECT id FROM pesanan WHERE kode = $1;

-- name: CreateDetailPesanan :exec
INSERT INTO detail_pesanan (pesanan_id, menu_id, jumlah) VALUES ($1, $2, $3);