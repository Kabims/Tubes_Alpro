package main
import "fmt"
func menu(A *TabTransaksi, i *int, j *int, barang *Tabjenis) {
	var pilih int = 0
	*j = 0
	for pilih != 4 {
		fmt.Println("            MENU           =")
		fmt.Println("============================")
		fmt.Println("= 1. Tambah Spare Part     =")
		fmt.Println("= 2. Edit Data             =")
		fmt.Println("= 3. Hapus Data            =")
		fmt.Println("* 4. Kembali               =")
		fmt.Println("============================")
		fmt.Print("Pilih Menu:")
		fmt.Scan(&pilih)
		if pilih == 1 {
			tambahSpare(&A[*i].sparePart, j,barang)
			*j++
		} else if pilih == 2 {
			pilihEdit(A, *i, *j,barang)
		} else if pilih == 3 {
			hapusData(A, i, j)
		}else if pilih==4{
			A[*i].harga= totalHarga(A[*i].sparePart,*j)
			fmt.Printf("Total bayar: Rp.%d,00\n",A[*i].harga)
		}
	}
}