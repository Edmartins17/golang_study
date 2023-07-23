package main

import "fmt"

func main(){

	palavra:= "palavra"
	for i:=0; i<len(palavra); i++ {
		if string(palavra[i])=="r"{
			break
		}
		fmt.Println(string(palavra[i]))
	}
}