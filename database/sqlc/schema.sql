CREATE TABLE meja (
    nomor INT PRIMARY KEY,
    dibuat_pada timestamp DEFAULT NOW()
);

CREATE TABLE kategori (
    id SERIAL PRIMARY KEY,
    kategori VARCHAR(20) NOT NULL
);

CREATE TABLE menu (
    id SERIAL PRIMARY KEY,
    kategori_id INT NOT NULL,
    menu varchar(20) NOT NULL,
    harga INT NOT NULL
);

CREATE TABLE pesanan (
    id SERIAL PRIMARY KEY,
    kode VARCHAR(5) NOT NULL,
    meja_nomor INT NOT NULL,
    dipesan_pada timestamp DEFAULT NOW()
);

CREATE TABLE detail_pesanan (
    pesanan_id INT NOT NULL,
    menu_id INT NOT NULL,
    harga INT NOT NULL,
    jumlah INT NOT NULL
);

CREATE TABLE pembayaran (
    id SERIAL PRIMARY KEY,
    pesanan_id INT NOT NULL,
    bayar INT NOT NULL,
    dibayar_pada timestamp DEFAULT NOW()
);