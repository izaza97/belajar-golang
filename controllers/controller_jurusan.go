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

//POST datajurusan
func DataJurusan(w http.ResponseWriter, _ *http.Request) {
	var pelajaran models.Pelajaran
	var arr_mapel []models.Pelajaran
	var response models.Response3

	db := DB.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT j.jurusan, p.mapel1, p.mapel2, p.mapel3 FROM jurusan j INNER JOIN pelajaran p ON j.prody_num = p.mapel_num")
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		if err := rows.Scan(&pelajaran.Jurusan, &pelajaran.Mapel1, &pelajaran.Mapel2, &pelajaran.Mapel3); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_mapel = append(arr_mapel, pelajaran)
		}
	}
	response.Status = true
	response.Message = "Success"
	response.Data = arr_mapel

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

//insert jurusan
func Insertjurusan(w http.ResponseWriter, r *http.Request) {
	db := DB.Connect()
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO jurusan (prody_num, jurusan) values (?,?)")
	if err != nil {
		log.Print(err)
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	prody_num := keyVal["prody_num"]
	jurusan := keyVal["jurusan"]

	_, err = stmt.Exec(prody_num, jurusan)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New post was created")

}

//update jurusan
func Updatejurusan(w http.ResponseWriter, r *http.Request) {
	db := DB.Connect()
	defer db.Close()
	stmt, err := db.Prepare("UPDATE jurusan SET jurusan = ? WHERE prody_num = ?")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	prody_num := keyVal["prody_num"]
	jurusan := keyVal["jurusan"]

	_, err = stmt.Exec(prody_num, jurusan)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Post was updated")

}

//delete jurusan
func Deletejurusan(w http.ResponseWriter, r *http.Request) {
	db := DB.Connect()
	defer db.Close()
	stmt, err := db.Prepare("DELETE FROM jurusan WHERE prody_num = ?")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	prody_num := keyVal["prody_num"]

	_, err = stmt.Exec(prody_num)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Post was deleted")
}
