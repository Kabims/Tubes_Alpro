package main
import "fmt"
func Tampil(A TabTransaksi,i,j int, ) {
	var a,b int
	fmt.Println("Periode data yang ingin ditampilkan")
	fmt.Print("Dari tanggal:")
	fmt.Scan(&a)
	fmt.Print("Sampai tanggal:")
	fmt.Scan(&b)
	for k := 0; k <= i; k++ {
		if A[k].dataPelanggan.tanggal>=a&&A[k].dataPelanggan.tanggal<=b {
			
			fmt.Printf("Nama Pelanggan  :%s\n", A[k].dataPelanggan.nama)
			fmt.Printf("Alamat          :%s\n", A[k].dataPelanggan.alamat)
			fmt.Printf("Nomer Telepon   :%s\n", A[k].dataPelanggan.telepon)
			for l := 0; l <= j; l++ {
				fmt.Printf("Sparepart Dibeli:%s\n", A[k].sparePart[l].nama)
				fmt.Printf("Jumlah          :%d\n", A[k].sparePart[l].jumlah)
			}
			fmt.Printf("\n")
		}
	}
}
