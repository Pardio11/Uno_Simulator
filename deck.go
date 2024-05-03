package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardColors := []string{"Red", "Blue", "Green", "Yellow"}
	cardNumbers := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight","Nine","Zero"}

	for _, color := range cardColors{
		for _, number := range cardNumbers{
			cards = append(cards, number+" "+color)
		}
	}
	return cards
}

func (d deck) showDeck() {
	for _, card := range d {
		fmt.Printf("%v\n",card)
	}
}

func deal(d deck,  handSize int) (deck, deck) {
	return d[:handSize],d[handSize:]
}

func (d deck) toString() string  {
	return strings.Join([]string(d), ",")
}
func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck{
	bs, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	s := strings.Split(string(bs),",") 
	return deck(s)
}

func (d deck) shuffle(){
	for i := range d{
		newPos := rand.Intn(len(d)-1)
		d[i], d[newPos] = d[newPos], d[i] 
	}
}