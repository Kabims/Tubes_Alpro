package main 
func dataSpare(barang *Tabjenis){
	barang[0].nama="Kampas rem"
	barang[1].nama="Cakram Rem"
	barang[2].nama="Ban"
	barang[3].nama="Aki"
	barang[4].nama="Oli"
	barang[5].nama="Lampu Depan"
	barang[6].nama="Lampu Belakang"
	barang[7].nama="Lampu Sen"
	barang[8].nama="Rantai"
	barang[9].nama="Busi"
}
func dummy(A *TabTransaksi,i *int,barang *Tabjenis){
	A[0].dataPelanggan.nama="Bimantara Ardi Winata"
	A[0].dataPelanggan.tanggal= 2
	A[0].dataPelanggan.alamat ="Jalan Raya Bojongsoang no 32"
	A[0].dataPelanggan.telepon= "081234567890"
		A[0].sparePart[0].kode = 1
		A[0].sparePart[0].jumlah = 2
		A[0].sparePart[0].nama = "Kampas rem"
		barang[0].jumlah+=2
		barang[0].nama = "Kampas rem"
		
	*i++
	A[1].dataPelanggan.nama="Damar Muharram"
	A[1].dataPelanggan.tanggal= 27
	A[1].dataPelanggan.alamat ="Jalan Raya Bojongsoang no 31"
	A[1].dataPelanggan.telepon= "081234567891"
		A[1].sparePart[0].kode = 3
		A[1].sparePart[0].jumlah = 2
		A[1].sparePart[0].nama = "Ban"
		barang[2].jumlah+=2
		barang[2].nama = "Ban"
	
	*i++
}