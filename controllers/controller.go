package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mux_crud/DB"
	"mux_crud/models"
	"net/http"

	"github.com/gorilla/mux"

)

//select by id GET
func Byid(w http.ResponseWriter, r *http.Request) {
	db := DB.Connect()
	defer db.Close()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result, err := db.Query("SELECT s.id, s.nama, s.alamat, s.kelas_num, s.prody_num FROM siswa s WHERE id = ?", params["id"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var siswa models.Siswa1
	for result.Next() {
		err := result.Scan(&siswa.Id, &siswa.Nama, &siswa.Alamat, &siswa.Kelas_num, &siswa.Prody_num)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(siswa)
}

//select datasiswa POST
func Selectdata(w http.ResponseWriter, r *http.Request) {
	var siswa models.Siswa
	var arr_siswa []models.Siswa
	var response models.Response2
	var object models.Object
	db := DB.Connect()
	defer db.Close()

	var count int
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	nama := keyVal["nama"]
	err2 := db.QueryRow("SELECT COUNT(s.id) FROM siswa s INNER JOIN jurusan j JOIN kelas k ON s.prody_num = j.prody_num AND s.kelas_num = k.kelas_num WHERE s.nama LIKE ?", fmt.Sprintf("%%%s%%", nama)).Scan(&count)
	if err2 != nil {
		log.Fatal(err.Error())
	}
	rows, err := db.Query("SELECT s.id, s.nama, s.alamat, k.kelas, j.jurusan FROM siswa s INNER JOIN jurusan j JOIN kelas k ON s.prody_num = j.prody_num AND s.kelas_num = k.kelas_num WHERE s.nama LIKE ?", fmt.Sprintf("%%%s%%", nama))
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
	if count > 0 {
		response.Status = true
		response.Message = "Success"
		response.Data = arr_siswa
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		object.Status = false
		object.Message = "Tidak Ditemukan"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(object)
	}
}

//select datasiswa POST
func Selectkelas(w http.ResponseWriter, r *http.Request) {
	var siswa models.Kelas
	var arr_siswa []models.Kelas
	var response models.Response
	var object models.Object
	db := DB.Connect()
	defer db.Close()

	var count int
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	kelas := keyVal["kelas"]
	err2 := db.QueryRow("SELECT COUNT(k.kelas_num) FROM siswa s INNER JOIN jurusan j JOIN kelas k ON s.prody_num = j.prody_num AND s.kelas_num = k.kelas_num WHERE k.kelas LIKE ?", fmt.Sprintf("%%%s%%", kelas)).Scan(&count)
	if err2 != nil {
		log.Fatal(err.Error())
	}
	rows, err := db.Query("SELECT s.nama, k.kelas, j.jurusan FROM siswa s INNER JOIN jurusan j JOIN kelas k ON s.prody_num = j.prody_num AND s.kelas_num = k.kelas_num WHERE k.kelas LIKE ? ORDER BY k.kelas ", fmt.Sprintf("%%%s%%", kelas))
	if err != nil {
		log.Print(err.Error())
	}
	for rows.Next() {
		if err := rows.Scan(&siswa.Nama, &siswa.Kelas, &siswa.Jurusan); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_siswa = append(arr_siswa, siswa)
		}
	}
	if count > 0 {
		response.Status = true
		response.Message = "Success"
		response.Data = arr_siswa
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		fmt.Println("masuk")
	} else {
		object.Status = false
		object.Message = "Tidak Ditemukan"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(object)
	}
}

//select datasiswa(jurusan) POST
func Selectjurusan(w http.ResponseWriter, r *http.Request) {
	var jrsn models.S_Jurusan
	var arr_jurusan []models.S_Jurusan
	var response models.RS_jurusan
	var object models.Object
	db := DB.Connect()
	defer db.Close()

	var count int
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	jurusan := keyVal["jurusan"]
	err2 := db.QueryRow("SELECT COUNT(k.kelas_num) FROM siswa s INNER JOIN jurusan j JOIN kelas k ON s.prody_num = j.prody_num AND s.kelas_num = k.kelas_num WHERE j.jurusan LIKE ?", fmt.Sprintf("%%%s%%", jurusan)).Scan(&count)
	if err2 != nil {
		log.Fatal(err.Error())
	}
	rows, err := db.Query("SELECT s.nama, k.kelas, j.jurusan FROM siswa s INNER JOIN jurusan j JOIN kelas k ON s.prody_num = j.prody_num AND s.kelas_num = k.kelas_num WHERE j.jurusan LIKE ? ORDER BY k.kelas ", fmt.Sprintf("%%%s%%", jurusan))
	if err != nil {
		log.Print(err.Error())
	}
	for rows.Next() {
		if err := rows.Scan(&jrsn.Nama, &jrsn.Kelas, &jrsn.Jurusan); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_jurusan = append(arr_jurusan, jrsn)
		}
	}
	if count > 0 {
		response.Status = true
		response.Message = "Success"
		response.Data = arr_jurusan
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		fmt.Println("masuk")
	} else {
		object.Status = false
		object.Message = "Tidak Ditemukan"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(object)
	}
}

