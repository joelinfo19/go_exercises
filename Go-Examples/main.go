package main

func main(){
	//var card string ="Ace of spades"
	//card:= newCar()
	//slices
	//cards:= newDeck()
	////cards=append(cards,"I am new")
	//cards.saveToFile("my_cards")
	//cards:= newDeckFromFile("mydrds")
	cards:=newDeck()
	cards.shuffle()
	cards.print()
	//fmt.Println(cards.toString())
	//hand,remainingCards:=deal(cards,5)
	//cards.print()
	//hand.print()
	//remainingCards.print()
}
//func newCard() string{
//	return "Ace of spades"
//}