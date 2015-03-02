package main

import (
	"fmt"
)

type Person struct{
	name string
	i int32	
}

type Emp struct {
	Person
	i int8
}

func main() {

	app := Emp{Person{"Me",20},10} 
	//fmt.Println("envr  %+v : ",env)
	fmt.Println("appl  %+v : ", app)

}
