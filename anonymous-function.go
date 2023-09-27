package main

// import "fmt"

// type Blacklist func(string) bool

// func registerUser(name string, blacklist Blacklist){
// 	if blacklist(name){
// 		fmt.Println("you are blocked" , name)
// 	}else {
// 		fmt.Println("welcome", name)
// 	}
// }

// func main() {
// 	// baris 17-19 merupakan anonymous function
// 	blacklist := func(name string) bool{
// 		return name == "anjing"
// 	}

// 	registerUser("anjing", func(name string)  bool{
// 		return name == "anjing"
// 	})
	
// 	registerUser("dani", blacklist)
// }