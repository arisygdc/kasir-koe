-- name: CreateMeja :exec
INSERT INTO meja (nomor, dibuat_pada) VALUES ($1, DEFAULT);

-- name: GetMejaAll :many
SELECT nomor FROM meja;

-- name: CreateKategori :exec
INSERT INTO kategori (id, kategori) VALUES (DEFAULT, $1);

-- name: CreateMenu :exec
INSERT INTO menu (id, kategori_id, menu, foto, deskripsi, tersedia, harga) VALUES (DEFAULT, $1, $2, $3, $4, $5, $6);

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
INSERT INTO pembayaran (id, pesanan_id, bayar, dibayar_pada) VALUES (DEFAULT, $1, $2, $3);

-- name: GetPesananHistory :many
SELECT kode, meja_nomor, dipesan_pada, bayar, dibayar_pada FROM pesanan 
LEFT JOIN pembayaran ON pesanan.id = pembayaran.pesanan_id 
where pesanan.dipesan_pada::date = date @date;

-- name: GetPesananByKD :one
SELECT p.id, p.meja_nomor, p.dipesan_pada, sum(dp.jumlah) as item, sum(dp.harga*dp.jumlah) as total FROM pesanan p 
LEFT JOIN detail_pesanan dp ON p.id = dp.pesanan_id 
where p.kode = $1 GROUP BY p.id;

-- name: GetPembayaranID :one
SELECT id FROM pembayaran WHERE pesanan_id = $1;