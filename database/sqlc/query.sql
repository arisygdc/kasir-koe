-- name: CreateMeja :exec
INSERT INTO meja (nomor, dibuat_pada) VALUES ($1, DEFAULT);

-- name: CreateKategori :exec
INSERT INTO kategori (id, kategori) VALUES (DEFAULT, $1);