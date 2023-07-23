package main

import "fmt"

func main(){

	palavra:= "palavra"
	for i:=0; i<len(palavra); i++ {
		fmt.Println(string(palavra[i]))
	}
}