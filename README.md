# Project Rent Book

  Rent Book merupakan aplikasi peminjaman buku antar user yang berjalan dengan Command Line Interface (CLI). Dimana terdapat beberapa fitur menarik yang dapat kalian temukan didalamnya, yaitu cari buku, lihat semua buku, anda dapat meminjam buku dan juga dapat meminjamkan buku.
  
# Package yang digunakan dalam project ini
	- Package Main
	- Package Model
	- Package Controller
	
# Cara menginstall project

Lakukan clon dari project ini dengan cara
	
```
git clone https://github.com/khalidrianda/Rent-Book.git
```

Setelah clone selesai, jalankan perintah `cd Rent-Book` 

Lalu kalian dapat menjalankan aplikasi rent book dengan cara

```
go run main.go
```

How To Use :
1. registrasi akun terlebih dahulu di menu 1 "Login" , kemudian pilih menu 2 "Registrasi".
   - isikan nama, email, password, alamat, dan foto profil.
2. jika kamu memiliki buku bisa tambahkan bukumu pada menu 4 "Buku Miliku", kemudian pilih menu 3 "Tambah buku".
3. jika kamu ingin meminjam buku seseorang kamu bisa login terlebih dahulu, lalu pilih menu 3 "Lihat buku", 
   kemudian jawab "Y" untuk meminjam buku lalu pilih buku yang akan dipinjam berdasarkan nomor ID buku tsb.
4.  Setelah selesai meminjam buku kamu bisa mengembalikannya dengan masuk ke menu 5 "Buku yang di pinjam",
  kemudian pilih menu 2 "kembalikan buku".
5. Kamu bisa menonaktifkan akunmu pada menu 2 "Update Profil" dengan syarat:
  - kamu sedang tidak meminjam buku
  - bukumu tidak sedang dipinjam oleh orang lain
6. selamat menikmati layanan kami 
