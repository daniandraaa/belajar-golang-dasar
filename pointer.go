package main 

import "fmt"

type Address struct {
	City, Province, Country string
}

func main(){
	address1 := Address{"Subang", "jawa barat", "Indonesia"}
	address2 := &address1 

	address2.City = "Bandung"

	address2 = &Address{"Malang", "jawa timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
}