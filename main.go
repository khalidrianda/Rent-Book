package main

import (
	"Rent-Book/controller"
	"Rent-Book/model"
	"fmt"
	"os"
	"os/exec"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectGorm() (*gorm.DB, error) {
	dsn := "root:@tcp(localhost:3306)/nasi_kotak?charset=utf8mb4&parseTime=True&loc=Local"
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
}

func main() {
	var isRun bool = true
	var inputMenu, input int

	Conn, err := connectGorm()
	if err != nil {
		fmt.Println("Cant connect to Database", err.Error())
	}
	migrate(Conn)
	bukuMdl := model.BukuModel{Conn}
	bukuCtl := controller.BukuControll{bukuMdl}

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
			case 1:
				fmt.Println("Ya")
			case 2:

			case 3:
				break
			}
		case 2:

		case 3:

		case 4:

		case 5:

		case 6:
			isRun = false
			clearBoard()
		}
	}

}
