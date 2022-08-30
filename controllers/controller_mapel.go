package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mux_crud/DB"
	"net/http"

)

//insert mapel
func Insertmapel(w http.ResponseWriter, r *http.Request) {
	db := DB.Connect()
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO pelajaran (mapel1, mapel2, mapel3, mapel_num) values (?,?,?,?)")
	if err != nil {
		log.Print(err)
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	mapel1 := keyVal["mapel1"]
	mapel2 := keyVal["mapel2"]
	mapel3 := keyVal["mapel3"]
	mapel_num := keyVal["mapel_num"]

	_, err = stmt.Exec(mapel1, mapel2, mapel3, mapel_num)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "New post was created")
}

//insert mapel
func Updatemapel(w http.ResponseWriter, r *http.Request) {
	db := DB.Connect()
	defer db.Close()
	stmt, err := db.Prepare("UPDATE pelajaran SET mapel1 = ?, mapel2 = ?, mapel3 = ? WHERE mapel_num = ?")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	mapel1 := keyVal["mapel1"]
	mapel2 := keyVal["mapel2"]
	mapel3 := keyVal["mapel3"]
	mapel_num := keyVal["mapel_num"]

	_, err = stmt.Exec(mapel1, mapel2, mapel3, mapel_num)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Post was updated")
}

//insert mapel
func Deletemapel(w http.ResponseWriter, r *http.Request) {
	db := DB.Connect()
	defer db.Close()
	stmt, err := db.Prepare("DELETE FROM pelajaran WHERE mapel_num= ?")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	mapel_num := keyVal["mapel_num"]

	_, err = stmt.Exec(mapel_num)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Post was deleted")
}
