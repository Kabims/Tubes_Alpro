package main
func totalHarga(A TabSpare, j int)int{
	var harga int
	for k := 0; k <= j; k++ {
		if A[k].kode == 1 {
			harga += A[k].jumlah * 20000
		} else if A[k].kode == 2 {
			harga += A[k].jumlah * 300000
		} else if A[k].kode == 3 {
			harga += A[k].jumlah * 250000
		} else if A[k].kode == 4 {
			harga += A[k].jumlah * 500000
		} else if A[k].kode == 5 {
			harga += A[k].jumlah * 100000
		} else if A[k].kode == 6 {
			harga += A[k].jumlah * 100000
		} else if A[k].kode == 7 {
			harga += A[k].jumlah * 50000
		} else if A[k].kode == 8 {
			harga += A[k].jumlah * 200000
		} else if A[k].kode == 9 {
			harga += A[k].jumlah * 100000
		} else if A[k].kode == 10 {
			harga += A[k].jumlah * 30000
		}  
	}
	return harga
}