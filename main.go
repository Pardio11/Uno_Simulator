package main

import "fmt"

func main() {
	cards :=  newDeck()
	cards.shuffle()
	cards.showDeck()
	cards, tableDeck := deal(cards, 4)
	println("\nhand deck: ")
	cards.shuffle()
	cards.showDeck()
	println("\ntable deck: ")
	tableDeck.showDeck()
	fmt.Println(cards.toString())
	cards.saveToFile("Cartas")
}