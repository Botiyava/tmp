/*
 *https://www.codewars.com/kata/53da3dbb4a5168369a0000fe/
 *Create a function (or write a script in Shell) that takes an integer as an
 *argument and returns "Even" for even numbers or "Odd" for odd numbers.
 */
package main

import "fmt"

func main(){
	number := 5
	fmt.Println([]string{"Even", "Odd"}[number & 1])
}
