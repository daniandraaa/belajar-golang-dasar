package main

// import "fmt"

// func endApp(){
// 	// message untuk menangkap pesan yang ada di panic
// 	message := recover()
// 	if message != nil {
// 		fmt.Println("erorr dengan message: ", message)
// 	}
// 	fmt.Println("aplikasi selesai" )
// }

// // dengan panic maka akan muncul di terminal beruipa output panic
// func runApp(error bool){

// 	defer endApp()

// 	if error{
// 		panic("aplikasi error")
// 	}
// 	fmt.Println("aplikasi berjalan")
// }

// func main(){
// 	runApp(false)
// }