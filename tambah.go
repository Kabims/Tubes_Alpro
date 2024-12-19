package main
import "fmt"
func tambahSpare(A *TabSpare, j *int,barang *Tabjenis) {
	fmt.Println("=       Sparepart          =")
	fmt.Println("============================")
	fmt.Println("=1.Kampas Rem              =")
	fmt.Println("=2.Cakram Rem              =")
	fmt.Println("=3.Ban                     =")
	fmt.Println("=4.Aki                     =")
	fmt.Println("=5.Oli                     =")
	fmt.Println("=6.Lampu Depan             =")
	fmt.Println("=7.Lampu Belakang          =")
	fmt.Println("=8.Lampu Sen               =")
	fmt.Println("=9.Rantai                  =")
	fmt.Println("=10.Busi                   =")
	fmt.Println("============================")
	fmt.Print("Pilih Sparepart:")
	fmt.Scan(&A[*j].kode)
	fmt.Print("Jumlah yang diingikan:")
	fmt.Scan(&A[*j].jumlah)
	if A[*j].kode == 1 {
		A[*j].nama="Kampas rem"
		barang[0].jumlah +=A[*j].jumlah
	} else if A[*j].kode == 2 {
		A[*j].nama="Cakram Rem"
		barang[1].jumlah +=A[*j].jumlah
	} else if A[*j].kode == 3 {
		A[*j].nama="Ban"
		barang[2].jumlah+=A[*j].jumlah
	} else if A[*j].kode == 4 {
		A[*j].nama="Aki"
		barang[3].jumlah+=A[*j].jumlah
	} else if A[*j].kode == 5 {
		A[*j].nama="Oli"
		barang[4].jumlah+=A[*j].jumlah
	} else if A[*j].kode == 6 {
		A[*j].nama="Lampu Depan"
		barang[5].jumlah +=A[*j].jumlah
	} else if A[*j].kode == 7 {
		A[*j].nama="Lampu Belakang"
		barang[6].jumlah +=A[*j].jumlah
	} else if A[*j].kode == 8 {
		A[*j].nama="Lampu Sen"
		barang[7].jumlah +=A[*j].jumlah
	} else if A[*j].kode == 9 {
		A[*j].nama="Rantai"
		barang[8].jumlah +=A[*j].jumlah
	} else if A[*j].kode == 10 {
		A[*j].nama="Busi"
		barang[9].jumlah +=A[*j].jumlah
	} 
	
	jumlah(A,A[*j].kode,j)
	
		
}