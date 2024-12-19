package main

import "fmt"

func urut(A *Tabjenis) {
	B := *A
	for k := 1; k < 10; k++ {
		min := k-1
		j := k
		for j < 10{
			if B[min].jumlah < B[j].jumlah {
				min = j
			}
			j++
		}
		ur := B[min]
		B[min] = B[k-1]
		B[k-1] = ur
	}
	fmt.Println("Data barang dibeli")
	for i := 0; i < 10; i++ {
		fmt.Println(B[i].nama," :",B[i].jumlah)
	}
}