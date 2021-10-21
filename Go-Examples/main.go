package main

import "fmt"

func main(){
	//var card string ="Ace of spades"
	card:= newCar()
	fmt.Println(card)
}
func newCar() string{
	return "Ace of spades"
}