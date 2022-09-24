package main

import (
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	name string
	address string
	job string
	reason string
}

func checkAbsen(absen int, slc []Student) {
	if absen <= len(slc) && absen > 0 {
		printBiodata(absen ,slc)
	} else {
		printError(len(slc))
	}
}

func printError(lengthValid int) {
	fmt.Printf("Maaf, absen yang valid hanya dari 1 hingga %d", lengthValid)
}

func printBiodata(absen int, slc []Student) {
	fmt.Printf("Absen ke-%d\n\nNama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan: %s", 
	absen, slc[absen - 1].name, slc[absen - 1].address, slc[absen - 1].job, slc[absen - 1].reason);
}

func main() {
	var student []Student = []Student{
		{
			name: "Agi Sahriza",
			address: "Banjarmasin",
			job: "Mahasiswa",
			reason: "Ingin menjadi backend developer dan kelas ini dapat memfasilitasi saya",
		},
		{
			name: "Daan Nur",
			address: "Batola",
			job: "Frontend Developer",
			reason: "Ingin menambah pengalaman di backend",
		},
		{
			name: "Fadri",
			address: "Batulicin",
			job: "IT",
			reason: "Saya cinta golang",
		},
		{
			name: "Gezalia Aylin",
			address: "Amuntai",
			job: "Mahasiswa",
			reason: "Karena golang cukup populer belakangan ini",
		},
		{
			name: "Hafizi",
			address: "Batola",
			job: "IT Support",
			reason: "Saya tertarik dengan backend",
		},
		{
			name: "Hania Zahratunnisa",
			address: "Amuntai",
			job: "Mahasiswa",
			reason: "Ikut ikut teman saja",
		},
		{
			name: "Hendra",
			address: "Banjarmasin",
			job: "Freelance",
			reason: "Banyak perusahaan besar mencari backend dengan bahasa golang sehingga saya tertarik mengikuti ini",
		},
		{
			name: "Muhammad Adit",
			address: "Batola",
			job: "Front End Developer",
			reason: "Ingin menambah pengalaman di backend",
		},
		{
			name: "Rahmat",
			address: "Banjarmasin",
			job: "Driver Ojek Online",
			reason: "Ingin mencoba hal baru",
		},
	}

	if len(os.Args) > 1 {
		absen, err := strconv.Atoi(os.Args[1])
		_ = err
		checkAbsen(absen, student)
	} else {
		fmt.Println("Anda tidak memasukan nomor absen!")
	}
}