package main

import (
	"fmt"
)

func main() {
	var arrays = []int{1, 4, 5, 6, 8, 2, 3, 7}
	var maxRow = arrays[0]
	var rows, cols int
	//cari nilai tertinggi baris
	for _, array := range arrays { //sort array tanpa indeks
		if array > maxRow {
			maxRow = array
		}
	}
	barchart(maxRow, arrays)
	fmt.Print("\n\n\n")

	//insertion sort
	fmt.Print("insertion sort\n\n")
	//gambar diagram
	barchart(maxRow, arrays)
	fmt.Print("\n\n\n")
	//panjang elemen
	totalElemen := len(arrays) - 1

	for i := 0; i <= totalElemen; i++ {
		for j := i - 1; j >= 0; j-- {
			if arrays[j+1] < arrays[j] {
				temp := arrays[j]
				arrays[j] = arrays[j+1]
				arrays[j+1] = temp
				barchart(maxRow, arrays)
				fmt.Print("\n\n\n")
			} else {
				break
			}
		}
	}

	//reverse sorting
	fmt.Print("Reverse Sorting \n\n")
	//gambar diagram
	for rows = maxRow; rows > 0; rows-- { //mulai dari atas ke bawah
		for cols = totalElemen; cols >= 0; cols-- { // sort reverse array
			if arrays[cols] >= rows {
				fmt.Print("| ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
	//gambar legend reverse
	for k := totalElemen; k >= 0; k-- {
		fmt.Print(arrays[k], " ")
	}

	fmt.Println()
}

func barchart(maxRow int, arrays []int) {
	var rows, cols int
	//gambar diagramnya
	for rows = maxRow; rows > 0; rows-- { //mulai dari atas ke bawah
		for cols = 0; cols <= len(arrays)-1; cols++ { // sort array
			if arrays[cols] >= rows {
				fmt.Print("| ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
	//gambar legend
	for _, array := range arrays {
		fmt.Print(array, " ")
	}
}
