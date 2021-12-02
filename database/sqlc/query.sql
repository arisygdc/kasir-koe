-- name: CreateMeja :exec
INSERT INTO meja (nomor, dibuat_pada) VALUES ($1, DEFAULT);

-- name: GetMejaAll :many
SELECT nomor FROM meja;

-- name: CreateKategori :exec
INSERT INTO kategori (id, kategori) VALUES (DEFAULT, $1);

-- name: CreateMenu :exec
INSERT INTO menu (id, kategori_id, menu, harga) VALUES (DEFAULT, $1, $2, $3);

-- name: GetMenuAll :many
SELECT kategori, menu, harga FROM menu RIGHT JOIN kategori ON menu.kategori_id = kategori.id;

-- name: CreatePesanan :exec
INSERT INTO pesanan (id, kode, meja_nomor, dipesan_pada) VALUES (DEFAULT, $1, $2, DEFAULT);

-- name: GetPesananID :one
SELECT id FROM pesanan WHERE kode = $1;

-- name: GetHarga :one
SELECT harga FROM menu WHERE id = $1;

-- name: CreateDetailPesanan :exec
INSERT INTO detail_pesanan (pesanan_id, menu_id, harga, jumlah) VALUES ($1, $2, $3, $4);

-- name: CreatePembayaran :exec
INSERT INTO pembayaran (id, pesanan_id, bayar, dibayar_pada) VALUES (DEFAULT, $1, $2, DEFAULT);