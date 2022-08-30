package DB

import (
	"database/sql"
	"log"

)

//koneksi ke database
func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/profil_siswa")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
