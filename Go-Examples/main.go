package main

import "fmt"

func main(){
	//var card string ="Ace of spades"
	//card:= newCar()
	//slices
	cards:= newDeck()
	//cards=append(cards,"I am new")
	fmt.Println(cards)
	fmt.Println(cards.toString())
	//hand,remainingCards:=deal(cards,5)
	//cards.print()
	//hand.print()
	//remainingCards.print()
}
//func newCard() string{
//	return "Ace of spades"
//}