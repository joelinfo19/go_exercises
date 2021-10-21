package main

import "fmt"

func main(){
	//var card string ="Ace of spades"
	//card:= newCar()
	//slices
	cards:= []string{"Gola",newCard()}
	cards=append(cards,"I am new")
	fmt.Println(cards)
}
func newCard() string{
	return "Ace of spades"
}