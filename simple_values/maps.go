package main

import "fmt"

func main() {
	m := make(map[string]string)
	// var m  map[string] string
	m["name"] = "rajat"
	fmt.Println(m["name"])

	m2:=make(map[string]int)
	m2["rajat"]=69
	val,valex:=m2["rajat"]
	if valex{
		fmt.Println(val)
	}
}