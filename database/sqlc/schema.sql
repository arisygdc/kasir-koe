CREATE TABLE meja (
    nomor INT PRIMARY KEY,
    dibuat_pada timestamp DEFAULT NOW()
);

CREATE TABLE kategori (
    id SERIAL PRIMARY KEY,
    kategori VARCHAR(20) NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE menu (
    id SERIAL PRIMARY KEY,
    kategori_id INT NOT NULL,
    menu varchar(20) NOT NULL
);

CREATE TABLE pesanan (
    id SERIAL PRIMARY KEY,
    kode VARCHAR(5) NOT NULL,
    dipesan_pada timestamp DEFAULT NOW()
);

CREATE TABLE detail_pesanan (
    pesanan_id INT PRIMARY KEY,
    menu_id INT NOT NULL,
    jumlah INT NOT NULL
);