package main

import (
	"Rent-Book/controller"
	"Rent-Book/model"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectGorm() (*gorm.DB, error) {
	dsn := "root:@tcp(localhost:3306)/rent_book?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func clearBoard() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Buku{})
	db.AutoMigrate(&model.User{}) // add migrate user
	db.AutoMigrate(&model.LendBook{})

}

func main() {
	var isRun bool = true
	var inputMenu, input int
	var session uint
	var inputString string

	Conn, err := connectGorm()
	if err != nil {
		fmt.Println("Cant connect to Database", err.Error())
	}
	migrate(Conn)
	bukuMdl := model.BukuModel{Conn}
	bukuCtl := controller.BukuControll{bukuMdl}
	userMdl := model.UserModel{Conn}
	UserCtl := controller.UserControll{userMdl}
	lendMdl := model.LendBookModel{Conn}
	lendCtrl := controller.LendBookControl{lendMdl}

	for isRun {
		fmt.Println("APLIKASI RENT BOOK")
		fmt.Println("------------------")
		fmt.Println("1. Login User")
		fmt.Println("2. Update Profil")
		fmt.Println("3. Lihat Buku")
		fmt.Println("4. Buku Milikku")
		fmt.Println("5. Buku yang dipinjam")
		fmt.Println("6. Exit")
		fmt.Print("Masukkan Input : ")
		fmt.Scanln(&inputMenu)
		clearBoard()
		switch inputMenu {
		case 1:
			fmt.Println("1. Login User")
			fmt.Println("2. Register")
			fmt.Println("3. Kembali")
			fmt.Print("Masukan input : ")
			fmt.Scanln(&input)

			switch input {
			case 1: // add login
				var logIn model.User
				fmt.Print("Email : ")
				fmt.Scanln(&logIn.Email)
				fmt.Print("Password: ")
				fmt.Scanln(&logIn.Password)

				res, err := UserCtl.GetAll(logIn)
				if err != nil {
					fmt.Println("Username/Password Salah", err.Error())
				} else {
					session = res.Id_user
					clearBoard()
					fmt.Printf("Selamat Datang, %v!!! \n", res.Nama_user)
				}

			case 2: // add register
				var newUser model.User
				fmt.Println("Register Account User")
				fmt.Print("Masukan Nama : ")
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				newUser.Nama_user = scanner.Text()
				fmt.Print("Masukan Email : ")
				fmt.Scanln(&newUser.Email)
				fmt.Print("Password : ")
				fmt.Scanln(&newUser.Password)
				fmt.Print("Masukan Alamat : ")
				scanner = bufio.NewScanner(os.Stdin)
				scanner.Scan()
				newUser.Alamat = scanner.Text()
				fmt.Print("Masukan Foto Profil : ")
				fmt.Scanln(&newUser.Foto_profil)
				newUser.Status_boolean = true

				res, err := UserCtl.Add(newUser)
				if err != nil {
					fmt.Println("some error on register", err.Error())
				}
				fmt.Println("Berhasil Registrasi", res)
			case 3: //keluar
				break
			}
		case 2:
			if session == 0 {
				fmt.Println("Login Required")
				continue
			}
			var pilih bool = true
			var plh int

			for pilih {
				fmt.Println("Update Profil")
				fmt.Println("1. Update")
				fmt.Println("2. Non-aktifkan akun")
				fmt.Println("3. Exit")
				fmt.Println("Select menu: ")
				fmt.Scanln(&plh)
				switch plh {
				case 1:
					if session == 0 {
						fmt.Println("Anda harus login dulu")
						continue
					}
					res, err := UserCtl.LogIn(session)
					if err != nil {
						fmt.Println("Some error on get", err.Error())

					}
					if res != nil {
						for i := 0; i < len(res); i++ {
							fmt.Printf("%v \n", res[i])
						}
					}
					var updUser model.User
					var n, e, p, a, f string
					updUser.Id_user = uint(session)
					// fmt.Println("Masukan Nama Update:")
					// fmt.Scanln(&n)
					fmt.Print("Masukan Nama Update : ")
					scanner := bufio.NewScanner(os.Stdin)
					scanner.Scan()
					n = scanner.Text()
					fmt.Print("Masukan Email Update : ")
					scanner.Scan()
					e = scanner.Text()
					fmt.Print("Masukan Password : ")
					scanner.Scan()
					p = scanner.Text()
					fmt.Print("Masukan Alamat Update : ")
					scanner.Scan()
					a = scanner.Text()
					fmt.Print("Masukan Poto Profil : ")
					fmt.Scanln(&f)

					if n != "" {
						updUser.Nama_user = a
						UserCtl.UpdateNama(updUser)
					}
					if e != "" {
						updUser.Email = e
						UserCtl.UpdateEmail(updUser)
					}
					if p != "" {
						updUser.Password = p
						UserCtl.UpdatePassword(updUser)
					}
					if a != "" {
						updUser.Alamat = a
						UserCtl.UpdateAlamat(updUser)
					}
					if f != "" {
						updUser.Foto_profil = f
						UserCtl.UpdateFotoProfil(updUser)
					}
					fmt.Println(updUser)

				case 2: // Non Aktifkan Profil
					var stats model.User
					var choice string
					stats.Status_boolean = false
					stats.Id_user = session
					fmt.Print("Apakah anda yakin ingin menonaktifkan akun? (Y/N)")
					fmt.Scanln(&choice)
					if choice == "Y" {
						UserCtl.UpdateStatus(stats)
					}

				case 3:
					pilih = false
					clearBoard()
				}
			}

		case 3:
			// add list buku
			res, err := bukuCtl.GetAll(session)
			if err != nil {
				fmt.Println("Some error on get", err.Error())

			}

			fmt.Println("ID \t Code \t Nama Buku \t Pengarang \t Gambar \t Deskripsi")
			for i := 0; i < len(res); i++ {
				fmt.Printf("%v \t %v \t %v \t %v\n", res[i].Id_buku, res[i].Nama_buku, res[i].Pengarang, res[i].Deskripsi)
			}

			fmt.Println("Apakah Anda ingin meminjam buku? (Y/N)")
			fmt.Scanln(&inputString)
			if session == 0 {
				fmt.Println("Anda haru login untuk meminjam buku")
			} else if inputString == "Y" {
				fmt.Println("Masukkan ID Buku yang ingin dipinjam : ")
				fmt.Scanln(&input)
				var pinjamBuku model.LendBook
				var tempBuku model.Buku
				pinjamBuku.Id_buku = uint(input)
				pinjamBuku.Id_peminjam = session
				temp, _ := bukuCtl.GetName(pinjamBuku.Id_buku)
				pinjamBuku.Nama_buku = temp.Nama_buku
				inOneMonth := time.Now().AddDate(0, 1, 0)
				pinjamBuku.Batas_waktu = inOneMonth
				lendCtrl.Add(pinjamBuku)
				tempBuku.Id_buku = input
				tempBuku.Is_lend = true
				bukuCtl.Dipinjam(tempBuku)
			} else {
				continue
			}
		case 4:
			if session == 0 {
				fmt.Println("Anda harus login dulu")
				continue
			}

			var ulang bool = true
			var inputBuku, inputBook int
			for ulang {
				fmt.Println("Menu Buku Milikku")
				fmt.Println("1. Lihat Buku milikku")
				fmt.Println("2. Edit Buku milikku")
				fmt.Println("3. Tambah Buku Milikku")
				fmt.Println("4. Kembali")
				fmt.Print("Masukkan Input : ")
				fmt.Scanln(&input)
				switch input {

				case 1:
					res, err := bukuCtl.GetMyBook(session)
					if err != nil {
						fmt.Println("Some error on get", err.Error())

					}
					if res != nil {
						fmt.Println("ID \t Code \t Nama Buku \t Pengarang \t Gambar \t Deskripsi")
						for i := 0; i < len(res); i++ {
							fmt.Printf("%v \t %v \t %v \t %v \t %v \t %v\n", res[i].Id_buku, res[i].Code_buku, res[i].Nama_buku, res[i].Pengarang, res[i].Gambar_buku, res[i].Deskripsi)
						}
					}
					if res == nil {
						fmt.Println("Anda Tidak Punya Buku")
					}

				case 2:
					res, err := bukuCtl.GetMyBook(session)
					if err != nil {
						fmt.Println("Some error on get", err.Error())

					}
					if res != nil {
						fmt.Println("Buku Millikku")
						fmt.Println("------------ ")
						fmt.Println("1. Ubah Data Buku : ")
						fmt.Println("2. Hapus Buku : ")
						fmt.Print("Masukkan input : ")
						fmt.Scanln(&inputBook)
						switch inputBook {
						case 1:
							fmt.Println("ID \t Code \t Nama Buku \t Pengarang \t Gambar \t Deskripsi")
							for i := 0; i < len(res); i++ {
								fmt.Printf("%v \t %v \t %v \t %v \t %v \t %v\n", res[i].Id_buku, res[i].Code_buku, res[i].Nama_buku, res[i].Pengarang, res[i].Gambar_buku, res[i].Deskripsi)
							}
							fmt.Print("Masukkan ID Buku yang Ingin Diubah : ")
							fmt.Scanln(&inputBuku)
							var newBuku model.Buku
							var a, b, c, d, e string
							newBuku.Id_user = uint(session)
							newBuku.Id_buku = inputBuku
							fmt.Print("Masukan Kode Buku : ")
							fmt.Scanln(&a)
							fmt.Print("Masukan Nama Buku : ")
							scanner := bufio.NewScanner(os.Stdin)
							scanner.Scan()
							b = scanner.Text()
							fmt.Print("Masukan Pengarang : ")
							scanner.Scan()
							c = scanner.Text()
							fmt.Print("Masukan Gambar buku : ")
							fmt.Scanln(&d)
							fmt.Print("Masukan Deskripsi buku : ")
							scanner.Scan()
							e = scanner.Text()
							if a != "" {
								newBuku.Code_buku = a
								bukuCtl.UpdateCode(newBuku)
							}
							if b != "" {
								newBuku.Nama_buku = b
								bukuCtl.UpdateNama(newBuku)
							}
							if c != "" {
								newBuku.Pengarang = c
								bukuCtl.UpdatePengarang(newBuku)
							}
							if d != "" {
								newBuku.Gambar_buku = d
								bukuCtl.UpdateGambar(newBuku)
							}
							if e != "" {
								newBuku.Deskripsi = e
								bukuCtl.UpdateDeskripsi(newBuku)
							}
						case 2:
							fmt.Println("ID \t Code \t Nama Buku \t Pengarang \t Gambar \t Deskripsi")
							for i := 0; i < len(res); i++ {
								fmt.Printf("%v \t %v \t %v \t %v \t %v \t %v\n", res[i].Id_buku, res[i].Code_buku, res[i].Nama_buku, res[i].Pengarang, res[i].Gambar_buku, res[i].Deskripsi)
							}
							fmt.Print("Masukkan ID Buku yang Ingin Dihapus : ")
							fmt.Scanln(&inputBook)
							var newBuku model.Buku
							newBuku.Id_buku = inputBook
							_, err := bukuCtl.Delete(newBuku)
							if err != nil {
								fmt.Println("some error on delete", err.Error())
							} else {
								fmt.Println("Buku telah dihapus")
							}
						}

					} else {
						fmt.Println("Anda Tidak Punya Buku")
					}

				case 3:
					if session != 0 {
						var newBuku model.Buku
						newBuku.Id_user = uint(session)
						fmt.Print("Masukan Kode Buku : ")
						fmt.Scanln(&newBuku.Code_buku)
						fmt.Print("Masukan Nama Buku : ")
						scanner := bufio.NewScanner(os.Stdin)
						scanner.Scan()
						newBuku.Nama_buku = scanner.Text()
						fmt.Print("Masukan Pengarang : ")
						scanner.Scan()
						newBuku.Pengarang = scanner.Text()
						fmt.Print("Masukan Gambar buku : ")
						fmt.Scanln(&newBuku.Gambar_buku)
						fmt.Print("Masukan Deskripsi buku : ")
						scanner.Scan()
						newBuku.Deskripsi = scanner.Text()

						res, err := bukuCtl.Add(newBuku)
						if err != nil {
							fmt.Println("some error on register", err.Error())
						}
						fmt.Println("Berhasil Registrasi", res)
					} else {
						fmt.Println("Login dulu untuk menambah buku")
					}
				case 4:
					ulang = false
					clearBoard()
				}
			}
		case 5:
			if session == 0 {
				fmt.Println("Anda harus login dulu")
				continue
			}
			res, err := lendCtrl.GetAll(session)
			if err != nil {
				fmt.Println("Some error on get", err.Error())

			}
			if res != nil {
				for i := 0; i < len(res); i++ {
					fmt.Printf("%v \n", res[i])
				}
			}
			var ulang bool = true
			var pilih int
			for ulang {
				fmt.Println("Buku yang dipinjam")
				fmt.Println("1. melihat buku yang dipinjam")
				fmt.Println("2. kembalikan buku yang dipinjam")
				fmt.Println("3. kembali")
				fmt.Println("pilih menu: ")
				fmt.Scanln(&pilih)

				switch pilih {
				case 1:
					res, err := lendCtrl.GetAll(session)
					if err != nil {
						fmt.Println("Some error on get", err.Error())

					}
					if res != nil {
						for i := 0; i < len(res); i++ {
							fmt.Printf("%v \n", res[i])
						}
					}
				case 2:
					res, err := lendCtrl.GetAll(session)
					if err != nil {
						fmt.Println("Some error on get", err.Error())

					}
					if res != nil {
						for i := 0; i < len(res); i++ {
							fmt.Printf("%v \n", res[i])
						}
					}
					var ipt int
					var back model.LendBook
					fmt.Println("masukan ID buku yang ingin anda kembalikan")
					fmt.Scanln(&ipt)
					var bk model.Buku
					back.Id_buku = uint(ipt)
					back.Kembalikan = true
					_, err = lendCtrl.Return(back)
					if err != nil {
						fmt.Println("Some error on get", err.Error())
					} else {
						bk.Is_lend = false
						bk.Id_buku = ipt
						bukuCtl.Dikembalikan(bk)
						continue
					}

				case 3:
					ulang = false
					clearBoard()
				}
			}
		case 6:
			isRun = false
			clearBoard()
		}
	}

}
