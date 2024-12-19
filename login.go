package main
import "fmt"

func login(A *Pelanggan) {
	fmt.Println("Masukkan Data Pelanggan")
	fmt.Print("Tanggal Service(01-31): ")
	fmt.Scan(&A.tanggal)
	fmt.Print("Nama Pelanggan: ")
	fmt.Scan(&A.nama)
	fmt.Print("Alamat: ")
	fmt.Scan(&A.alamat)
	fmt.Print("Nomer Telepon: ")
	fmt.Scan(&A.telepon)
}

