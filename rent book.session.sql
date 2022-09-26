create database rent-book;
use rent-book;

DROP TABLE user;
DROP TABLE buku;
DROP TABLE lend_book;

CREATE TABLE user(
    id_user int NOT NULL auto_increment primary key,
    nama_user varchar(50) NOT NULL,
    email varchar(50) NOT NULL,
    password varchar(50) NOT NULL,
    alamat varchar(50) NOT NULL,
    foto_profil varchar(50) NOT NULL,
    status boolean default false,
    create_at TIMESTAMP default CURRENT_TIMESTAMP(),
    updated_at TIMESTAMP default CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP()
);

CREATE TABLE buku(
    id_buku int NOT NULL auto_increment primary key,
    id_user int,
    code_buku varchar(50) NOT NULL,
    nama_buku varchar(50) NOT NULL,
    pengarang varchar(50) NOT NULL,
    gambar_buku varchar(50) NOT NULL,
    deskripsi varchar(255) NOT NULL,
    is_lend boolean default false,
    added_at TIMESTAMP default CURRENT_TIMESTAMP(),
    updated_at TIMESTAMP default CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP(),
    constraint FK_Id_User foreign key(id_user) references user(id_user)
);

CREATE TABLE lend_book(
    id_peminjaman int NOT NULL auto_increment primary key,
    id_peminjam int,
    id_buku int,
    nama_buku varchar(50) NOT NULL,
    batas_waktu datetime NOT NULL,
    added_at TIMESTAMP default CURRENT_TIMESTAMP(),
    return_at TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP(),
    constraint FK_User_Pinjam foreign key(id_peminjaman) references user(id_user),
    constraint FK_ID_Buku foreign key(id_buku) references buku(id_buku)
);