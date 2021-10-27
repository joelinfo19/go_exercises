package main

import "fmt"

/* pointer exercise
type contactInfo struct{
	email string
	zipCode  int

}
type person struct {
	firstName string
	lastName string
	contactInfo


}

func main(){
	//PROJECT CARDS
	//var card string ="Ace of spades"
	//card:= newCar()
	//slices
	//cards:= newDeck()
	////cards=append(cards,"I am new")
	//cards.saveToFile("my_cards")
	//cards:= newDeckFromFile("mydrds")
	//cards:=newDeck()
	//cards.shuffle()
	//cards.print()
	//fmt.Println(cards.toString())
	//hand,remainingCards:=deal(cards,5)
	//cards.print()
	//hand.print()
	//remainingCards.print()
	//PROJECT STRUCT


	jim:=person{
		firstName: "Jim",
		lastName: "Carrey",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 9400,
		},

	}
	//jimPointer:= &jim
	//jimPointer.updateName("jimyto")
	//is the same when is a pointer
	jim.updateName("jimyto")
	jim.print()

}
func (p person) print(){
	fmt.Printf("%+v",p)


}
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName=newFirstName

}
           pointer exercise*/
//func newCard() string{
//	return "Ace of spades"
//}

/* map exercises
func main(){

	//first mode for map
	colors:= map[string]string{
		"red":"red",
		"green":"green",
	}
	//fmt.Println(colors)
	//second mode
	//colors:= make(map[string]string)
	//colors["hola"]="hola"
	//delete(colors,"hola")
	//fmt.Println(colors)
	printMap(colors)
}
func printMap(c map[string]string) {
	for color,hex:=range c{
		fmt.Println("Hex code for",color,"is",hex)
	}

}
*/
/* INTERFACE GO*/
// It help us to reduce code
type bot interface {
	getGreeting() string
}

type englishBot struct{
}
type spanishBot struct {

}
func main(){
	eb:= englishBot{}
	printGreeting(eb)
}
func printGreeting(b bot)  {
	fmt.Println(b.getGreeting())
}
func (englishBot) getGreeting() string{
	return "hi there!"
}