package main
import "fmt"
func pilihEdit(A *TabTransaksi, i int, j int, barang *Tabjenis) {
	var pilih int
	fmt.Println("Pilih data yang ingin diubah")
	fmt.Println("============================")
	fmt.Println("= 1. Edit Data Sparepart   =")
	fmt.Println("= 2. Edit Data  Pelanggan  =")
	fmt.Println("============================")
	fmt.Scan(&pilih)
	if pilih == 1 {
		editSpare(&A[i].sparePart, j, barang)
	} else if pilih == 2 {
		editPel(A, i)
	}

}
func editSpare(A *TabSpare, j int, barang *Tabjenis) {
	var indeksA, indeksB int
	var kode int
	fmt.Println("spare part yang ingin diubah: ")
	fmt.Scan(&kode)
	// indeksA = cariIndeksA(*A, kode, j)
	indeksB = cariIndeksB(*A, kode, j)

	if indeksA == -1||indeksB==-1 {
		fmt.Println("Data tidak ditemukan")
	} else {
		tambahSpare((A), &indeksB,barang)
	}
}
func editPel(A *TabTransaksi, i int) {
	var indeks int
	var nama string
	fmt.Println("Nomor Pelanggan yang ingin diubah: ")
	fmt.Scan(&nama)
	indeks = cariPel(*A, i, nama)
	if indeks == -1 {
		fmt.Println("Data tidak ditemukan")
	} else {
		login(&A[indeks].dataPelanggan)
	}
}