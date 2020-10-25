package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

func main() {
	getRequest, err := http.Get("http://localhost:8080/mahasiswa")

	if err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
	}

	defer getRequest.Body.Close()

	rawData, err := ioutil.ReadAll(getRequest.Body)

	if err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
	}
	mhs := []mahasiswa{}
	jsonErr := json.Unmarshal(rawData, &mhs)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}
	for i := 0; i < len(mhs); i++ {
		fmt.Println("NoBP : ", mhs[i].NoBp)
		fmt.Println("Nama : ", mhs[i].Nama)
		fmt.Println("Jurusan : ", mhs[i].Jurusan)
		fmt.Println("Jalan : ", mhs[i].Alamat.Jalan)
		fmt.Println("Kelurahan : ", mhs[i].Alamat.Kelurahan)
		fmt.Println("Kecamatan : ", mhs[i].Alamat.Kecamatan)
		fmt.Println("Kabupaten : ", mhs[i].Alamat.Kabupaten)
		fmt.Println("Provinsi : ", mhs[i].Alamat.Provinsi)

		for _, nilai := range mhs[i].Nilai {
			fmt.Println("No BP", nilai.NoBp)
			fmt.Println("Ip", nilai.Ip)
			fmt.Println("Semester", nilai.Semester)
		}
		fmt.Printf("\n")
	}
}
