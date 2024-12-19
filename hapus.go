package main
import "fmt"
func hapusTpel(A *TabTransaksi, n *int) {
	var indeks int
	var nama string
	
	fmt.Println("Masukkan nama pelanggan yang ingin dihapus")
	fmt.Scan(&nama)
	indeks = cariPel(*A, *n, nama)
	if indeks != -1 {
		for i := indeks; i <= *n; i++ {
			A[i].dataPelanggan = A[i+1].dataPelanggan
		}
		*n--
		fmt.Println("Data pelanggan berhasil dihapus")

	} else {
		fmt.Println("Data tidak valid")
	}
}
func hapusspare(A *TabSpare,j *int) {
	var indeksB int
	var kode int
	fmt.Println("Nomor sparepart yang ingin dihapus:")
	fmt.Scan(&kode)
	indeksB = cariIndeksB(*A,kode,*j)
	if indeksB == -1 {
		fmt.Println("Posisi tidak valid")
	}else {
			for k := indeksB; k <= *j; k++ {
				A[k] = A[k+1]
			}
			*j--
			fmt.Println("Data berhasil dihapus")
	}
}
func hapusData(A *TabTransaksi, i *int,j *int) {
	var pilih int
	fmt.Println("Pilih data yang ingin dihapus")
	fmt.Println("============================")
	fmt.Println("= 1. Hapus Data Pelanggan    =")
	fmt.Println("= 2. Hapus Data Spare Part   =")
	fmt.Println("============================")
	fmt.Scan(&pilih)
	if pilih == 1 {
		hapusTpel(A,i)
	} else if pilih == 2 {
		hapusspare(&A[*i].sparePart,j)
	}
}