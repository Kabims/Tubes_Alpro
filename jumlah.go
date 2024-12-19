package main
func jumlah(A *TabSpare,x int, j *int){
	index := cariIndeksB(*A,x,*j-1)
	if index !=-1 {
		A[index].jumlah =A[index].jumlah +A[*j].jumlah 
		*j--
	}
}