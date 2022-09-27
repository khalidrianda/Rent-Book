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
	// var session model.User

	Conn, err := connectGorm()
	if err != nil {
		fmt.Println("Cant connect to Database", err.Error())
	}
	migrate(Conn)
	bukuMdl := model.BukuModel{Conn}
	bukuCtl := controller.BukuControll{bukuMdl}
	userMdl := model.UserModel{Conn}
	UserCtl := controller.UserControll{userMdl}

	for isRun {
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

				res, err := UserCtl.GetAll()
				if err != nil {
					fmt.Println("Username/Password Salah", err)
				}

				session := res[logIn.Id_user].Id_user
				fmt.Println(session)

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
			// update
			var N UserModel
			fmt.Println("pilih id user")
			fmt.Scanln(&N.id_user)
			user, err := UserCtl.GetAll(N)
			fmt.Println("Ganti Nama user")
			fmt.Scanln(&user.Nama_user)
			fmt.Println("Ganti Email user")
			fmt.Scanln(&user.Email)
			fmt.Println("Ganti Password user")
			fmt.Scanln(&user.Password)
			fmt.Println("Ganti Alamat user")
			fmt.Scanln(&user.Alamat)
			err = userMdl.Update(&user)
			if err != nil {
				fmt.Println("Update failed")
			} else {

				fmt.Println(user.Insert(model.UserModel))
			}

		case 3:
			// add list buku
			res, err := bukuCtl.GetAll()
			if err != nil {
				fmt.Println("Some error on get", err.Error())

			}
			fmt.Println(res)

		case 4:
			fmt.Println("Menu Buku Milikku")
			fmt.Println("1. Lihat Buku milikku")
			fmt.Println("2. Tambah Buku Milikku")
			fmt.Println("3. Kembali")
			fmt.Print("Masukkan Input : ")
			fmt.Scanln(&input)
			switch input {
			case 1:

			case 2:
				var newBuku model.Buku
				fmt.Print("Masukan Nama : ")
				scanner := bufio.NewScanner(os.Stdin)
				scanner.Scan()
				newBuku.Nama_buku = scanner.Text()
				// fmt.Print("Masukan Email : ")
				// fmt.Scanln(&newUser.Email)
				// fmt.Print("Password : ")
				// fmt.Scanln(&newUser.Password)
			case 3:
				break
			}
			var newUser model.User // newuser model
			fmt.Print("Masukan Nama : ")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			newUser.Nama_user = scanner.Text()
			fmt.Print("Masukan Email : ")
			fmt.Scanln(&newUser.Email)
			fmt.Print("Password : ")
			fmt.Scanln(&newUser.Password)
		case 5:

		case 6:
			isRun = false
			clearBoard()
		}
	}

}
