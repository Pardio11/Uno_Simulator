package main

import "fmt"

func main() {
	d := newDeck()
	d.shuffle()
	d.showDeck()
	d, tableDeck := deal(d, 4)
	tableDeck.showDeck()
	d.saveToFile("deck.json")
	println("\nHand cards:")
	fmt.Println(d)
	tableDeck = newDeckFromFile("deck.json")
	println("\nTable cards:")
	tableDeck.showDeck()
	/* 	cards.shuffle()
	   	cards.showDeck()
	   	cards, tableDeck := deal(cards, 4)
	   	println("\nhand deck: ")
	   	cards.shuffle()
	   	cards.showDeck()
	   	println("\ntable deck: ")
	   	tableDeck.showDeck()
	   	fmt.Println(cards.toString())
	   	cards.saveToFile("Cartas") */
}