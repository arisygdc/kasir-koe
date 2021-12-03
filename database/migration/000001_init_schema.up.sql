CREATE TABLE meja (
    nomor INT,
    dibuat_pada timestamp DEFAULT NOW(),
    PRIMARY KEY(nomor)
);

CREATE TABLE kategori (
    id SERIAL,
    kategori VARCHAR(20) NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE menu (
    id SERIAL,
    kategori_id INT NOT NULL,
    menu varchar(20) NOT NULL,
    deskripsi varchar(150) NOT NULL,
    foto varchar(50) NOT NULL,
    tersedia boolean NOT NULL DEFAULT true,
    harga INT NOT NULL,
    PRIMARY KEY(id)
);

CREATE TABLE pesanan (
    id SERIAL,
    kode VARCHAR(10) NOT NULL UNIQUE,
    meja_nomor INT NOT NULL,
    dipesan_pada timestamp DEFAULT NOW(),
    PRIMARY KEY(id),

    CONSTRAINT meja
        FOREIGN KEY (meja_nomor)
        REFERENCES meja(nomor)
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);

CREATE TABLE detail_pesanan (
    pesanan_id INT NOT NULL,
    menu_id INT NOT NULL,
    harga INT NOT NULL,
    jumlah INT NOT NULL,

    CONSTRAINT pesanan
        FOREIGN KEY (pesanan_id)
        REFERENCES pesanan(id),

    CONSTRAINT menu
        FOREIGN KEY (menu_id)
        REFERENCES menu(id)
);

CREATE TABLE pembayaran (
    id SERIAL,
    pesanan_id INT NOT NULL UNIQUE,
    bayar INT NOT NULL,
    dibayar_pada timestamp DEFAULT NOW(),
    PRIMARY KEY(id),

    CONSTRAINT pesanan
        FOREIGN KEY (pesanan_id)
        REFERENCES pesanan(id)
);