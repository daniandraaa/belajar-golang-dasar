package main

// import "fmt"

// func main() {
// 	var months = [...]string{
// 		"januari",
// 		"februari",
// 		"maret",
// 		"april",
// 		"mei",
// 		"juni",
// 		"juli",
// 		"agustus",
// 		"september",
// 		"oktober",
// 		"november",
// 		"desember",
// 	}

// 	var slice1 = months[4 :7]
// 	fmt.Println(slice1)

// 	// length data slice
// 	fmt.Println(len(slice1))

// 	// capacity mulai dari yg nilai awal ke nilai akhir dikurang 1
// 	fmt.Println(cap(slice1))

// 	// months[5] = "diubah"
// 	// fmt.Println(slice1)

// 	// index 0 dibawah merupakan index awal dari variabel slice1
// 	// slice1[0] = "Mei Lagi"
// 	// fmt.Println(slice1)

// 	// slice append
// 	var slice2 = months[10:]
// 	fmt.Println(slice2)

// 	var slice3 = append(slice2, "dani")
// 	fmt.Println(slice3)
// 	slice3[1] = "bukan desember"
// 	fmt.Println(slice3)

// 	fmt.Println(slice2)
// 	fmt.Println(months)

// 	// pembuatan slice secara langsung

// 	newSlice := make([]string, 2, 5)

// 	newSlice[0] = "dani"
// 	newSlice[1] = "ilham"

// 	fmt.Println(newSlice)
// 	fmt.Println(len(newSlice))
// 	fmt.Println(cap(newSlice))

// 	// mengcopy slice yang ada di slice yang baru
// 	copySlice := make([]string, len(newSlice), cap(newSlice))
// 	copy(copySlice, newSlice)
// 	fmt.Println(newSlice)

// 	// perbedaan antara array dan slice
// 	iniArray := [5]int{1,2,3,4,5}
// 	iniSlice := []int{1,2,3,4,5}

// 	fmt.Println(iniArray)
// 	fmt.Println(iniSlice)

// }