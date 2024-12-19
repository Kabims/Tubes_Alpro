package main

import "fmt"
type Transaksi struct {
	dataPelanggan Pelanggan
	sparePart     TabSpare
	harga  int
}
type Pelanggan struct {
	tanggal int
	nama    string
	alamat  string
	telepon string
}
type sparePart struct {
	kode   int
	nama   string
	jumlah int
}

type jenisSpare struct{
	nama string
	jumlah int
}

type TabTransaksi [100]Transaksi
type TabSpare [100]sparePart
type Tabjenis [11] jenisSpare

func main() {
	var A TabTransaksi
	var pilih, j, i int = 0, 0, 0
	var barang Tabjenis
	dataSpare(&barang)
	dummy(&A,&i,&barang)
	for pilih != 4 {
		fmt.Println("            Home")
		fmt.Println("============================")
		fmt.Println("= 1. Login                 =")
		fmt.Println("= 2. Tampilkan Data        =")
		fmt.Println("= 3. Data Barang           =")
		fmt.Println("= 4. Exit                  =")
		fmt.Println("============================")
		fmt.Scan(&pilih)
		if pilih == 1 {
			login(&A[i].dataPelanggan)
			menu(&A, &i, &j, &barang)
			i++
		} else if pilih == 2 {

			Tampil(A, i, j)
		} else if pilih == 3{
			urut(&barang)
		}
	}
}
