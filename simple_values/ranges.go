package main

import "fmt"

func main() {
	arr := []int{2, 3, 4}
	for i, val := range arr {
		fmt.Println(i,val)
	}
	// mp:=map[string]int
	// mp["rajat"]=69    invalid

	mp := make(map[string]int)

    // Adding values
    mp["rajat"] = 69
    mp["paras"] = 85
    mp["neha"] = 92
    mp["arjun"] = 74
    mp["kiran"] = 88

    // Print all key-value pairs
    for name, score := range mp {
        fmt.Printf("%s: %d\n", name, score)
	}


}