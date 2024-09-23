-- Membuat tabel book
CREATE TABLE book (
    id SERIAL PRIMARY KEY,                -- Kolom ID unik dan auto-increment
    title VARCHAR(255) NOT NULL,          -- Judul buku
    author VARCHAR(255),                  -- Nama penulis
    published_date DATE,                  -- Tanggal publikasi
    category_id INT REFERENCES category_book(id), -- Foreign key ke tabel category_book
    price NUMERIC(10, 2),                 -- Harga buku (dengan 2 desimal)
    stock INT DEFAULT 0                   -- Jumlah stok buku
);