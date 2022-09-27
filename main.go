package main

import (
	"Rent-Book/controller"
	"Rent-Book/model"
	"bufio"
	"fmt"
	"os"
	"os/exec"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectGorm() (*gorm.DB, error) {
	dsn := "root:@tcp(localhost:3306)/rent-book?charset=utf8mb4&parseTime=True&loc=Local"
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
	// var inputString string

	Conn, err := connectGorm()
	if err != nil {
		fmt.Println("Cant connect to Database", err.Error())
	}
	migrate(Conn)
	bukuMdl := model.BukuModel{Conn}
	bukuCtl := controller.BukuControll{bukuMdl}
	userMdl := model.UserModel{Conn}
	UserCtl := controller.UserControll{userMdl}
	// lendMdl := model.UserModel{Conn}
	// lendCtrl := controller.UserControll{lendMdl}

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
				}

				// session := res[logIn.Id_user].Id_user
				// fmt.Println(session)

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

				res, err := UserCtl.Add(newUser)
				if err != nil {
					fmt.Println("some error on register", err.Error())
				}
				fmt.Println("Berhasil Registrasi", res)
			case 3:
				break
			}
		case 2:

		case 3:
			// add list buku
			res, err := bukuCtl.GetAll()
			if err != nil {
				fmt.Println("Some error on get", err.Error())

			}
			for i := 0; i < len(res); i++ {
				fmt.Printf("%v \n", res[i])
			}

		case 4:
			if session == 0 {
				fmt.Println("Anda harus login dulu")
				continue
			}

			var ulang bool = true
			var inputBuku int
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
						for i := 0; i < len(res); i++ {
							fmt.Printf("%v \n", res[i])
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
						for i := 0; i < len(res); i++ {
							fmt.Printf("%v \n", res[i])
						}
						fmt.Print("Masukkan code buku yang ingin anda ubah : ")
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
						fmt.Println(newBuku)
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

		case 6:
			isRun = false
			clearBoard()
		}
	}

}
