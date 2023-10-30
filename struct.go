package main

// import "fmt"

// type Customer struct {
// 	Name, Address string 
// 	Age int
// }

// // struct function
// func (customer Customer) sayHi(name string){
// 	fmt.Println("hello", name, "my name is", customer.Name)
// }

// func main() {
// 	// cara 1
// 	var dani Customer 
// 	dani.Name = "Daniandra"
// 	dani.Address = "Indonesia"
// 	dani.Age = 30

// 	// pemanggilan struct function
// 	dani.sayHi("sembarang")

// 	fmt.Println(dani.Name)
// 	fmt.Println(dani.Address)
// 	fmt.Println(dani.Age)

// 	// cara 2
// 	joko := Customer{
// 		Name: "Dani",
// 		Address: "Antang",
// 		Age: 20,
// 	}
// 	fmt.Println(joko)

// 	// cara 3
// 	budi := Customer{"yudist" , "makassar", 10}
// 	fmt.Println(budi)
// }