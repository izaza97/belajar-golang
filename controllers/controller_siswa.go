package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mux_crud/DB"
	"mux_crud/models"
	"net/http"

)

//POST datasiswa
func DataSiswa(w http.ResponseWriter, r *http.Request) {
	var siswa models.Siswa
	var arr_siswa []models.Siswa
	var response models.Response2

	db := DB.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT s.id, s.nama, s.alamat, k.kelas, j.jurusan FROM siswa s INNER JOIN jurusan j JOIN kelas k ON s.prody_num = j.prody_num AND s.kelas_num = k.kelas_num ORDER BY s.id")
	if err != nil {
		log.Print(err.Error())
	}
	for rows.Next() {
		if err := rows.Scan(&siswa.Id, &siswa.Nama, &siswa.Alamat, &siswa.Kelas, &siswa.Jurusan); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_siswa = append(arr_siswa, siswa)
		}
	}
	response.Status = true
	response.Message = "Success"
	response.Data = arr_siswa

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//insert siswa
func Insertsiswa(w http.ResponseWriter, r *http.Request) {
	db := DB.Connect()
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO siswa (nama, alamat, kelas_num, prody_num) values (?,?,?,?)")
	if err != nil {
		log.Print(err)
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	nama := keyVal["nama"]
	alamat := keyVal["alamat"]
	kelas_num := keyVal["kelas_num"]
	prody_num := keyVal["prody_num"]

	_, err = stmt.Exec(nama, alamat, kelas_num, prody_num)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New post was created")
}

//update siswa
func Updatesiswa(w http.ResponseWriter, r *http.Request) {
	db := DB.Connect()
	defer db.Close()
	stmt, err := db.Prepare("UPDATE siswa SET nama = ?, alamat = ?, kelas_num = ?, prody_num = ? WHERE id = ?")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	nama := keyVal["nama"]
	alamat := keyVal["alamat"]
	kelas_num := keyVal["kelas_num"]
	prody_num := keyVal["prody_num"]
	id := keyVal["id"]

	_, err = stmt.Exec(nama, alamat, kelas_num, prody_num, id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Post was updated")
}

//delete siswa
func Deletesiswa(w http.ResponseWriter, r *http.Request) {
	db := DB.Connect()
	defer db.Close()
	stmt, err := db.Prepare("DELETE FROM siswa WHERE id = ?")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	id := keyVal["id"]

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Post was deleted")
}
