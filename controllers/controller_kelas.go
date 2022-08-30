package controllers

import (
	"encoding/json"
	"log"
	"mux_crud/DB"
	"mux_crud/models"
	"net/http"

)

func DataKelas(w http.ResponseWriter, _ *http.Request) {
	var siswa models.Kelas
	var arr_siswa []models.Kelas
	var response models.Response

	db := DB.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT k.kelas, s.nama, j.jurusan FROM kelas k INNER JOIN siswa s JOIN jurusan j ON s.kelas_num = k.kelas_num AND s.prody_num = j.prody_num")
	if err != nil {
		log.Print(err.Error())
	}
	for rows.Next() {
		if err := rows.Scan(&siswa.Kelas, &siswa.Nama, &siswa.Jurusan); err != nil {
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
