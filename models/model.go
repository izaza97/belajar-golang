package models

// type POST/GET datasiswa
type Siswa struct {
	Id      string `form:"id" json:"id"`
	Nama    string `form:"nama" json:"nama"`
	Alamat  string `form:"alamat" json:"alamat"`
	Kelas   int    `form:"kelas" json:"kelas"`
	Jurusan string `form:"jurusan" json:"jurusan"`
}

type Siswa1 struct {
	Id        string `form:"id" json:"id"`
	Nama      string `form:"nama" json:"nama"`
	Alamat    string `form:"alamat" json:"alamat"`
	Kelas_num int    `form:"kelas_num" json:"kelas_num"`
	Prody_num string `form:"jurusan_num" json:"jurusan_num"`
}

type Kelas struct {
	Kelas   int    `form:"kelas" json:"kelas"`
	Nama    string `form:"nama" json:"nama"`
	Jurusan string `form:"jurusan" json:"jurusan"`
}

//type POST datajurusan
type Pelajaran struct {
	Jurusan string `form:"jurusan" json:"List Pelajaran"`
	Mapel1  string `form:"mapel1" json:"1"`
	Mapel2  string `form:"mapel2" json:"2"`
	Mapel3  string `form:"mapel3" json:"3"`
}

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    []Kelas
}

//response datasiswa
type Response2 struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    []Siswa
}

//response datajurusan
type Response3 struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    []Pelajaran
}

//response UID
type Response4 struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type Object struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
