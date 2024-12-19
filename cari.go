package main 
func cariPel(A TabTransaksi, n int, x string) int {
	var found int
	found = -1
	for i := 0; found == -1 && i <= n; i++ {
		if A[i].dataPelanggan.nama == x {
			found = i
		}
	}
	return found

}
func cariIndeksB(A TabSpare, x int, j int) int {
	var found int
	found = -1
	for K:= 0; K<=j && found ==-1;K++{
		if A[K].kode==x {
			found = K
		}
	}
	return found
}