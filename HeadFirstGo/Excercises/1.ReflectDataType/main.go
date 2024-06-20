package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(42))    //print type: int
	fmt.Println(reflect.TypeOf(3.14))  //print type: float64
	fmt.Println(reflect.TypeOf(true))  //print type: bool
	fmt.Println(reflect.TypeOf("Sam")) //print type: string
}
