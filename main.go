package main

import (
	"fmt"
	"log"
	"mux_crud/controllers"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/getsiswa/{id}", controllers.Byid).Methods("POST")
	router.HandleFunc("/siswa", controllers.Selectdata).Methods("POST")
	router.HandleFunc("/kelas", controllers.Selectkelas).Methods("POST")
	router.HandleFunc("/addsiswa", controllers.Insertsiswa).Methods("POST")
	router.HandleFunc("/addjurusan", controllers.Insertjurusan).Methods("POST")
	router.HandleFunc("/addmapel", controllers.Insertmapel).Methods("POST")
	router.HandleFunc("/updatesiswa", controllers.Updatesiswa).Methods("POST")
	router.HandleFunc("/updatejurusan", controllers.Updatejurusan).Methods("POST")
	router.HandleFunc("/updatemapel", controllers.Updatemapel).Methods("POST")
	router.HandleFunc("/deletesiswa", controllers.Deletesiswa).Methods("POST")
	router.HandleFunc("/deletejurusan", controllers.Deletejurusan).Methods("POST")
	router.HandleFunc("/deletemapel", controllers.Deletemapel).Methods("POST")
	router.HandleFunc("/datasiswa", controllers.DataSiswa).Methods("GET")
	router.HandleFunc("/datakelas", controllers.DataKelas).Methods("GET")
	router.HandleFunc("/datajurusan", controllers.DataJurusan).Methods("GET")
	http.Handle("/", router)
	fmt.Println("Connected to port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
