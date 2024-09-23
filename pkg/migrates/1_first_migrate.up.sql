-- Membuat tabel category_book
CREATE TABLE category_book (
    id SERIAL PRIMARY KEY,       -- Kolom ID unik dan auto-increment
    name VARCHAR(100) NOT NULL   -- Nama kategori buku
);
