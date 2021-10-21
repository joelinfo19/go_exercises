package main

import "fmt"

func main(){
	//var card string ="Ace of spades"
	//card:= newCar()
	//slices
	cards:= []string{"Gola",newCard()}
	cards=append(cards,"I am new")
	for i,card:= range cards{
		fmt.Println(i,card)
	}
	fmt.Println(cards)
}
func newCard() string{
	return "Ace of spades"
}