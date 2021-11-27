CREATE TABLE meja (
    nomor INT NOT NULL,
    dibuat_pada timestamp NULL DEFAULT NOW(),
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
    PRIMARY KEY(id),

    CONSTRAINT meja
        FOREIGN KEY (kategori_id)
        REFERENCES kategori(id)
);

CREATE TABLE pesanan (
    id SERIAL,
    kode VARCHAR(5) NOT NULL,
    dipesan_pada timestamp NULL DEFAULT NOW(),
    PRIMARY KEY(id)
);

CREATE TABLE detail_pesanan (
    pesanan_id INT NOT NULL,
    menu_id INT NOT NULL,
    jumlah INT NOT NULL,

    CONSTRAINT pesanan
        FOREIGN KEY (pesanan_id)
        REFERENCES pesanan(id),

    CONSTRAINT menu
        FOREIGN KEY (menu_id)
        REFERENCES menu(id)
);