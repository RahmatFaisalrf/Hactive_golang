package main

import (
	"fmt"
	"os"
)
type Profile struct {
	Code, Nama, Alamat, Pekerjaan, Alasan string
}

func main () {
	args := os.Args[1:]
	data := [] Profile {
		Profile{
			Code		: "001",
			Nama 		: "Rahmat Faisal",
			Alamat		: "jakarta",
			Pekerjaan	: "Pelajar",
			Alasan 		: "menambah skillset",
		},
		Profile{
			Code		: "002",
			Nama 		: "Sultan",
			Alamat		: "Bangka",
			Pekerjaan	: "Tukang Sate",
			Alasan 		: "Mau coba coba",	
		},
		Profile{
			Code		: "003",
			Nama 		: "Aya",
			Alamat		: "Tanjung Enim",
			Pekerjaan	: "Tukang Sate",
			Alasan 		: "Mau coba coba",
		},
	}
	for _, val := range data {
		if args[0] == val.Code{
			fmt.Println("Code", val.Code)
			fmt.Println("Nama", val.Nama)
			fmt.Println("Alamat", val.Alamat)
			fmt.Println("Pekerjaan", val.Pekerjaan)
			fmt.Println("Alasan", val.Alasan)
		}
	}
}