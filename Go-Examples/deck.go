package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string
// "deck" TYPE is how a "class"
// deck[includenumber:notincludenumber]
func newDeck() deck{
	cards:=deck{}
	cardSuits:=[]string{"Spades","Diamonds","Hearts","Clubs"}
	cardValue:=[]string{"one","two","three"}
	for _,suit :=range cardSuits{
		for _,value:= range cardValue{
			cards=append(cards,value + " of " + suit)
		}
	}
	return cards

}
func (d deck) print(){
	for i,card:=range d{
		fmt.Println(i,card)
	}
}
func deal(d deck,handSize int) (deck,deck){
	return d[:handSize],d[handSize:]
}