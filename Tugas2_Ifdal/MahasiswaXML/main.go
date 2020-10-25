package main

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

type mahasiswa struct {
	NoBp    int    `json:"NoBp"`
	Nama    string `json:"Nama"`
	Jurusan string `json:"Jurusan"`
	Alamat  struct {
		Jalan     string `json:"Jalan"`
		Kelurahan string `json:"Kelurahan"`
		Kecamatan string `json:"Kecamatan"`
		Kabupaten string `json:"Kabupaten"`
		Provinsi  string `json:"Provinsi"`
	} `json:"Alamat"`
	Nilai []nilai `json:"Nilai"`
}

type nilai struct {
	NoBp     int    `json:"NoBp"`
	Ip       string `json:"Ip"`
	Semester string `json:"Semester"`
}

func getMahasiswa(w http.ResponseWriter, r *http.Request) {

	var mhs []mahasiswa
	params := mux.Vars(r)
	sql := `SELECT
				nobp,
				IFNULL(nama,'') nama,
				IFNULL(nobp,'') nobp,
				IFNULL(jurusan,'') jurusan,
				IFNULL(jalan,'') jalan,
				IFNULL(kelurahan,'') kelurahan,
				IFNULL(kecamatan,'') kecamatan,
				IFNULL(kabupaten,'') kabupaten,
				IFNULL(provinsi,'') provinsi				
			FROM mahasiswa WHERE nobp IN (?)`

	result, err := db.Query(sql, params["NoBp"])
	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var mh mahasiswa
		err := result.Scan(&mh.NoBp, &mh.Nama, &mh.Jurusan, &mh.Alamat.Jalan, &mh.Alamat.Kelurahan, &mh.Alamat.Kecamatan, &mh.Alamat.Kabupaten, &mh.Alamat.Provinsi)

		if err != nil {
			panic(err.Error())
		}

		sqlNilai := `Select nobp, mahasiswa.nobp, mahasiswa.nama, ip, semester from mahasiswa INNER JOIN nilai ON (mahasiswa.nobp = nilai.nobp) WHERE nobp = ?`

		resultNilai, errNilai := db.Query(sqlNilai, mh.NoBp)

		defer resultNilai.Close()

		if errNilai != nil {
			panic(err.Error())
		}

		for resultNilai.Next() {
			var value nilai
			err := resultNilai.Scan(&value.NoBp, &value.Ip, &value.Semester)
			if err != nil {
				panic(err.Error())
			}
			mh.Nilai = append(mh.Nilai, value)
		}
		mhs = append(mhs, mh)
	}

	w.Write([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n"))
	xml.NewEncoder(w).Encode(mhs)

}

func main() {

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/mahasiswa")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	r := mux.NewRouter()

	r.HandleFunc("/mahasiswa/{NoBp}", getMahasiswa).Methods("GET")

	fmt.Println("Server on :8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
